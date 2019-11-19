package openFoodFacts

import (
	"encoding/csv"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
	"context"

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
	str, _ := formating.ProductToString(names, data, bsonData)
	// Logger.Printf("%q -> %q -> %s", names, data, str)
	// TODO Set token cookie
	// r.Header.Set("Cookie", "token="+config.token)
	http.Post(config.APIURL+"/api/consumable/add",
		"application/json", strings.NewReader(str))
	time.Sleep(config.Timing * time.Millisecond)
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
        if config.Ctx != nil {
            // TODO Find proper filter to find the proper item
            // filter := bson.D{{"_id": 0}}
            filter := bson.D{{}}
            err = config.Collection.FindOne(config.Ctx, filter).Decode(&result)
            if err != nil {
                Logger.Println(err)
                result = nil
            }
        }
		sendData(names, rec, result, i)
		i++
		bar.Increment()
	}
	bar.Finish()
}

func Login(c *conf.Configuration) {
	config = c
	config.SetCtx(context.Background())
    client, err := mongo.Connect(config.Ctx,
        options.Client().ApplyURI(config.DBURL))
    if err != nil {
        Logger.Fatal(err)
    }

    config.SetClient(client)
    if err = config.Client.Ping(config.Ctx, readpref.Primary()); err != nil {
        Logger.Fatal(err)
    }
    config.SetCollection(config.Client.Database("off").Collection("products"))
    Logger.Println("Database connected")
}

func Run(c *conf.Configuration) {
    config = c
	sendAllData()
}
