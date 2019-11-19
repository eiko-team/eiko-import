package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"crypto/tls"

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
		config.APIEmail, config.APIPass)
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	req, err := http.NewRequest("POST", config.APIURL+"/api/login",
		strings.NewReader(body))
	if err != nil {
		Logger.Fatal(err)
	}
	got, err := config.HClient.Do(req)
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
