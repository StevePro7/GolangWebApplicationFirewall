package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HomeFunc(w http.ResponseWriter, r *http.Request) {
	log.Println("xAPI Home -start")
	log.Println("xAPI url:", r.URL)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data := "xAPI Home function"
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		_, err := fmt.Fprintf(w, "%s", err.Error())
		if err != nil {
			log.Fatalln(err)
		}
	}
	log.Println("xAPI Home --end-")
}

func TestFunc(w http.ResponseWriter, r *http.Request) {
	log.Println("xAPI Test -start")
	log.Println("xAPI url:", r.URL)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data := "xAPI Test function"
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		_, err := fmt.Fprintf(w, "%s", err.Error())
		if err != nil {
			log.Fatalln(err)
		}
	}
	log.Println("xAPI Test --end-")
}

func main() {
	log.Println("xAPI start [3.0]")
	bind := ":3080"
	gmux := mux.NewRouter()
	gmux.HandleFunc("/", HomeFunc).Methods("GET")
	gmux.HandleFunc("/test/artists.php", TestFunc).Methods("GET")
	log.Printf("xAPI listen on [%s]\n", bind)
	log.Fatal(http.ListenAndServe(bind, gmux))
}
