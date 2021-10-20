package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HelloFunc(w http.ResponseWriter, _ *http.Request) {
	log.Println("API Hello -start")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data := "API Hello function"
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		_, err := fmt.Fprintf(w, "%s", err.Error())
		if err != nil {
			log.Fatalln(err)
		}
	}
	log.Println("API Hello --end-")
}

func main() {
	log.Println("API start")
	bind := ":3080"
	gmux := mux.NewRouter()
	gmux.HandleFunc("/", HelloFunc).Methods("GET")
	log.Printf("API listen on [%s]\n", bind)
	log.Fatal(http.ListenAndServe(bind, gmux))
}
