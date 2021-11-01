package main

// #include "modsec.c"
import "C"

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func InitModSec() {
	log.Println("initModSec start")
	C.MyCInit()
	log.Println("initModSec -end-")
}

func main() {
	bind := ":8081"
	log.Println("Start web server on port", bind)
	InitModSec()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		if err != nil {
			log.Fatal(err)
			return
		}
	})
	log.Fatal(http.ListenAndServe(bind, nil))
}
