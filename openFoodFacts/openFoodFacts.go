package openFoodFacts

import (
	"encoding/csv"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	conf "github.com/eiko-team/eiko-import/config"
	"github.com/eiko-team/eiko-import/formating"
	"github.com/eiko-team/eiko/misc/log"
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

func sendData(names []string, data []string, i int64) {
	str, _ := formating.ProductToString(names, data)
	// Logger.Printf("%q -> %q -> %s", names, data, str)
	// TODO: set token cookie
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
		sendData(names, rec, i)
		i++
	}
}

func Run(c *conf.Configuration) {
	config = c
	sendAllData()
}
