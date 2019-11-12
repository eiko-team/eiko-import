package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/eiko-team/eiko-off/config"
)

var (
	// Logger used to log output
	Logger = log.New(os.Stdout, "api: ",
		log.Ldate|log.Ltime|log.Lshortfile)
)

func Login(config config.Configuration) {
	body := fmt.Sprintf(`{"user_email":"%s","user_password":"%s"}`,
		config.GetAPIEmail(), config.GetAPIPass())
	got, err := http.Post(config.GetAPIURL()+"/api/login",
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
	config.SetToken(t.Token)
}
