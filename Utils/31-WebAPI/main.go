package main

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"html"
	"net/http"
	"net/url"
)

func main() {

	bind := ":8081"
	log.Println("Start web server on port", bind)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		if err != nil {
			log.Fatal(err)
			return
		}
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		// Get HTTP GET parameters as a map
		getParameters, _ := url.ParseQuery(r.URL.RawQuery)
		getParametersJson, err := json.Marshal(getParameters)

		log.WithFields(log.Fields{
			"getParameters": string(getParametersJson),
		}).Info("Request received!")

		_, err = fmt.Fprintf(w, "Hi")
		if err != nil {
			log.Fatal(err)
			return
		}
	})

	log.Fatal(http.ListenAndServe(bind, nil))
}
