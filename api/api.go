package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/eiko-team/eiko-import/config"
	"github.com/eiko-team/eiko/misc/log"
)

var (
	// Logger used to log output
	Logger = log.New(os.Stdout, "api",
		log.Ldate|log.Ltime|log.Lshortfile)
)

func Login(config *config.Configuration) {
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
	Logger.Printf(`connected","token":"%s`, t.Token)
}
