package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type configuration struct {
	APIEmail   string        `json:"api_email"`
	APIPass    string        `json:"api_pass"`
	APIHost    string        `json:"api_host"`
	APIPort    string        `json:"api_port"`
	DBHost     string        `json:"db_host"`
	DBPort     string        `json:"db_port"`
	Timing     time.Duration `json:"timing"`
	Token      string
	APIURL     string
	DBURL      string
	Ctx        context.Context
	Client     *mongo.Client
	Collection *mongo.Collection
}

var (
	// Logger used to log output
	Logger = log.New(os.Stdout, "eiko-OFF: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	config configuration
)

func Init() configuration {
	c := configuration{}
	configFile := os.Getenv("CONFIG")
	if configFile == "" {
		configFile = "config.json"
	}
	Logger.Printf("Setting configuration to %s", configFile)
	file, err := os.Open(configFile)
	if err != nil {
		Logger.Fatal(err)
	}
	err = json.NewDecoder(file).Decode(&c)
	if err != nil {
		Logger.Fatal(err)
	}

	c.APIURL = c.APIHost
	if c.APIPort != "" {
		c.APIURL += ":" + c.APIPort
	}
	c.DBURL = "mongodb://" + c.DBHost
	if c.DBPort != "" {
		c.DBURL += ":" + c.DBPort
	}
	return c
}

func loginAPI() {
	body := fmt.Sprintf(`{"user_email":"%s","user_password":"%s"}`,
		config.APIEmail, config.APIPass)
	got, err := http.Post(config.APIURL+"/api/login",
		"application/json", strings.NewReader(body))
	if err != nil {
		Logger.Fatal(err)
	}
	t := struct {
		Token string `json:"token"`
	}{}
	err = json.NewDecoder(got.Body).Decode(&t)
	if err != nil {
		Logger.Fatal(err)
	}
	config.Token = t.Token
}

func loginDB() {
	var err error
	config.Ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	config.Client, err = mongo.Connect(config.Ctx, options.Client().ApplyURI(config.DBURL))
	if err != nil {
		Logger.Fatal(err)
	}
	err = config.Client.Ping(config.Ctx, readpref.Primary())
	if err != nil {
		Logger.Fatal(err)
	}
	config.Collection = config.Client.Database("off").Collection("products")
}

func sendData(data bson.M) {
	json, _ := json.Marshal(data)
	Logger.Printf("%s\n\n\n", string(json))
	http.Post(config.APIURL+"/api/consumable/add",
		"application/json", strings.NewReader(string(json)))
	time.Sleep(config.Timing * time.Millisecond)
}

func sendAllData() {
	cur, err := config.Collection.Find(config.Ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(config.Ctx)

	for cur.Next(config.Ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		sendData(result)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	config = Init()
	loginAPI()
	loginDB()
	sendAllData()
}
