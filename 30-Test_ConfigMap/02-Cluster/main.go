package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	bind := ":8081"
	log.Println("Start web server on port", bind)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		url := fmt.Sprintf("%q", html.EscapeString(r.URL.Path))
		log.Println("endpoint : ", url)
		_, err := fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		if err != nil {
			log.Fatal(err)
			return
		}
	})
	log.Fatal(http.ListenAndServe(bind, nil))
}
