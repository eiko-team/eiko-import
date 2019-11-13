package openFoodFacts

import (
	"context"
	"net/http"
	"os"
	"strings"
	"time"

	conf "github.com/eiko-team/eiko-import/config"
	"github.com/eiko-team/eiko-import/formating"
	"github.com/eiko-team/eiko/misc/log"

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

func Login(c *conf.Configuration) {
	config = c
	var err error
	config.SetCtx(context.Background())
	client, err := mongo.Connect(config.GetCtx(),
		options.Client().ApplyURI(config.GetDBURL()))
	if err != nil {
		Logger.Fatal(err)
	}
	config.SetClient(client)
	err = config.GetClient().Ping(config.GetCtx(), readpref.Primary())
	if err != nil {
		Logger.Fatal(err)
	}
	config.SetCollection(config.
		GetClient().
		Database("off").
		Collection("products"))
	Logger.Println("connected")
}

func sendData(data bson.M, i int64) {
	str, _ := formating.BsonToString(data)
	Logger.Println(str)
	// TODO: set token cookie
	// r.Header.Set("Cookie", "token="+config.token)
	http.Post(config.GetAPIURL()+"/api/consumable/add",
		"application/json", strings.NewReader(str))
	time.Sleep(config.GetTiming() * time.Millisecond)
}

func sendAllData() {
	cur, err := config.Collection.Find(config.GetCtx(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(config.GetCtx())

	var nb int64
	for cur.Next(config.GetCtx()) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		go sendData(result, nb)
		nb += 1
		if nb > 500 {
			break
		}
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
}

func Run() {
	sendAllData()
	// TODO: fetch the 14 days update
	// https://static.openfoodfacts.org/data/delta/index.txt
	/*
		go func() {
			for { // infinite
				time.Sleep(14 * 60 * time.Hour)
				fetchUpdate()
			}
		}()
	*/
}
