package config

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	// Logger used to log output
	Logger = log.New(os.Stdout, "config: ",
		log.Ldate|log.Ltime|log.Lshortfile)
)

type Configuration struct {
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

func Init() Configuration {
	c := Configuration{}
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

func (config Configuration) GetAPIEmail() string {
	return config.APIEmail
}

func (config Configuration) GetAPIPass() string {
	return config.APIPass
}

func (config Configuration) GetAPIHost() string {
	return config.APIHost
}

func (config Configuration) GetAPIPort() string {
	return config.APIPort
}

func (config Configuration) GetDBHost() string {
	return config.DBHost
}

func (config Configuration) GetDBPort() string {
	return config.DBPort
}

func (config Configuration) GetTiming() time.Duration {
	return config.Timing
}

func (config Configuration) GetToken() string {
	return config.Token
}

func (config Configuration) SetToken(token string) {
	config.Token = token
}

func (config Configuration) GetAPIURL() string {
	return config.APIURL
}

func (config Configuration) GetDBURL() string {
	return config.DBURL
}

func (config Configuration) GetCtx() context.Context {
	return config.Ctx
}

func (config Configuration) SetCtx(ctx context.Context) {
	config.Ctx = ctx
}

func (config Configuration) GetClient() *mongo.Client {
	return config.Client
}

func (config Configuration) SetClient(client *mongo.Client) {
	config.Client = client
}

func (config Configuration) GetCollection() *mongo.Collection {
	return config.Collection
}

func (config Configuration) SetCollection(collection *mongo.Collection) {
	config.Collection = collection
}
