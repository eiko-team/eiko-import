package config

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"github.com/eiko-team/eiko/misc/log"

	"go.mongodb.org/mongo-driver/mongo"
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

var (
	// Logger used to log output
	Logger = log.New(os.Stdout, "config",
		log.Ldate|log.Ltime|log.Lshortfile)

	config Configuration
)

func Init() *Configuration {
	config := Configuration{}
	configFile := os.Getenv("CONFIG")
	if configFile == "" {
		configFile = "config.json"
	}
	Logger.Printf("Setting configuration to %s", configFile)
	file, err := os.Open(configFile)
	if err != nil {
		Logger.Fatal(err)
	}
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		Logger.Fatal(err)
	}

	config.APIURL = config.APIHost
	if config.APIPort != "" {
		config.APIURL += ":" + config.APIPort
	}
	config.DBURL = "mongodb://" + config.DBHost
	if config.DBPort != "" {
		config.DBURL += ":" + config.DBPort
	}
	return &config
}

func (config *Configuration) Print() {
	Logger.Printf("%+v", config)
}

func (config *Configuration) GetAPIEmail() string {
	return config.APIEmail
}

func (config *Configuration) GetAPIPass() string {
	return config.APIPass
}

func (config *Configuration) GetAPIHost() string {
	return config.APIHost
}

func (config *Configuration) GetAPIPort() string {
	return config.APIPort
}

func (config *Configuration) GetDBHost() string {
	return config.DBHost
}

func (config *Configuration) GetDBPort() string {
	return config.DBPort
}

func (config *Configuration) GetTiming() time.Duration {
	return config.Timing
}

func (config *Configuration) GetToken() string {
	return config.Token
}

func (config *Configuration) SetToken(token string) {
	config.Token = token
}

func (config *Configuration) GetAPIURL() string {
	return config.APIURL
}

func (config *Configuration) GetDBURL() string {
	return config.DBURL
}

func (config *Configuration) GetCtx() context.Context {
	return config.Ctx
}

func (config *Configuration) SetCtx(ctx context.Context) {
	config.Ctx = ctx
}

func (config *Configuration) GetClient() *mongo.Client {
	return config.Client
}

func (config *Configuration) SetClient(client *mongo.Client) {
	config.Client = client
}

func (config *Configuration) GetCollection() *mongo.Collection {
	return config.Collection
}

func (config *Configuration) SetCollection(collection *mongo.Collection) {
	config.Collection = collection
}
