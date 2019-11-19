package openFoodFacts

import (
	"encoding/csv"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
	"context"
	"crypto/tls"

	conf "github.com/eiko-team/eiko-import/config"
	"github.com/eiko-team/eiko-import/formating"
	"github.com/eiko-team/eiko/misc/log"

	"github.com/cheggaaa/pb/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	// Logger used to log output
	Logger = log.New(os.Stdout, "openFoodFacts",
		log.Ldate|log.Ltime|log.Lshortfile)

	config *conf.Configuration
)

type FieldsReader struct {
	*csv.Reader
}

func sendData(names []string, data []string, bsonData bson.M, i int64) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	str, _ := formating.ProductToString(names, data, bsonData)
	// Logger.Printf("%q -> %q -> %s", names, data, str)
	// TODO Set token cookie
	req, _ := http.NewRequest("POST", config.APIURL+"/api/consumable/add",
		strings.NewReader(str))
	req.Header.Set("Cookie", "token="+config.Token)
	got, err := config.HClient.Do(req)
	if got != nil || err != nil {
		body, _ := ioutil.ReadAll(got.Body)
		Logger.Println(string(body), err)
	}
	time.Sleep(config.Timing * time.Millisecond)
}

func findName(names []string, data []string) string {
	for i, val := range names {
		if val == "code" {
			return data[i]
		}
	}
	return ""
}

func sendAllData() {
	file, err := os.Open(config.OffFile)
	if err != nil {
		Logger.Fatal(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	r.Comma = '\t'

	// TODO Find proper file length
	count := 1047595
	bar := pb.StartNew(count)

	names, err := r.Read()
	if err != nil {
		Logger.Fatal(err)
	}
	var i int64
	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			Logger.Fatal(err)
		}

		var result bson.M
		if code := findName(names, rec); config.Ctx != nil &&
			config.Collection != nil && code != "" {
			filter := bson.M{"code": code}
			err = config.Collection.FindOne(config.Ctx, filter).Decode(&result)
			if err != nil {
				Logger.Println(err)
				result = nil
			}
		}
		sendData(names, rec, result, i)
		i++
		bar.Increment()
		if i > 100 {
			break
		}
	}
	bar.Finish()
}

func Login(c *conf.Configuration) {
	config = c
	config.SetCtx(context.Background())
	client, err := mongo.Connect(config.Ctx,
		options.Client().ApplyURI(config.DBURL))
	if err != nil {
		Logger.Println("Could not connect to database -", err)
		return
	}

	config.SetClient(client)
	if err = config.Client.Ping(config.Ctx, readpref.Primary()); err != nil {
		Logger.Println("Could not connect to database -", err)
		return
	}
	config.SetCollection(config.Client.Database("off").Collection("products"))
	Logger.Println("Database connected")
}

func Run() {
	sendAllData()
}
