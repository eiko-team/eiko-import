package config

import (
	"encoding/json"
	"os"
	"time"

	"github.com/eiko-team/eiko/misc/log"
)

type Configuration struct {
	APIEmail string        `json:"api_email"`
	APIPass  string        `json:"api_pass"`
	APIHost  string        `json:"api_host"`
	APIPort  string        `json:"api_port"`
	OffFile  string        `json:"off_filepath"`
	Timing   time.Duration `json:"timing"`
	Token    string
	APIURL   string
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
	return &config
}

func (config *Configuration) Print() {
	Logger.Printf("%+v", config)
}

func (config *Configuration) SetToken(token string) {
	config.Token = token
}
