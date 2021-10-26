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
		_, err := fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		if err != nil {
			log.Fatal(err)
			return
		}
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Hi")
		if err != nil {
			log.Fatal(err)
			return
		}
	})

	log.Fatal(http.ListenAndServe(bind, nil))
}
