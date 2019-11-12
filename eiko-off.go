package main

import (
	"log"
	"os"

	"github.com/eiko-team/eiko-off/api"
	"github.com/eiko-team/eiko-off/config"
	"github.com/eiko-team/eiko-off/openFoodFacts"
)

var (
	// Logger used to log output
	Logger = log.New(os.Stdout, "eiko-OFF: ",
		log.Ldate|log.Ltime|log.Lshortfile)
)

func main() {
	conf := config.Init()
	api.Login(conf)
	openFoodFacts.Login(conf)
	openFoodFacts.Run()
}
