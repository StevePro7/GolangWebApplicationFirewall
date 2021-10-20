package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HomeFunc(w http.ResponseWriter, _ *http.Request) {
	log.Println("API Home -start")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data := "API Home function"
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		_, err := fmt.Fprintf(w, "%s", err.Error())
		if err != nil {
			log.Fatalln(err)
		}
	}
	log.Println("API Home --end-")
}

func TestFunc(w http.ResponseWriter, _ *http.Request) {
	log.Println("API Test -start")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data := "API Test function"
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		_, err := fmt.Fprintf(w, "%s", err.Error())
		if err != nil {
			log.Fatalln(err)
		}
	}
	log.Println("API Test --end-")
}

func main() {
	log.Println("API start")
	bind := ":3080"
	gmux := mux.NewRouter()
	gmux.HandleFunc("/", HomeFunc).Methods("GET")
	gmux.HandleFunc("/test/artists.php", TestFunc).Methods("GET")
	log.Printf("API listen on [%s]\n", bind)
	log.Fatal(http.ListenAndServe(bind, gmux))
}
