package main

import (
	"os"

	"github.com/eiko-team/eiko-import/api"
	"github.com/eiko-team/eiko-import/config"
	"github.com/eiko-team/eiko-import/openFoodFacts"
	"github.com/eiko-team/eiko/misc/log"
)

var (
	// Logger used to log output
	Logger = log.New(os.Stdout, "main",
		log.Ldate|log.Ltime|log.Lshortfile)
)

func main() {
	Logger.Println("Starting")
	conf := config.Init()
	api.Login(conf)
	openFoodFacts.Login(conf)
	openFoodFacts.Run()
}
