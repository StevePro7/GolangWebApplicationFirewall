package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HelloFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data := "Hello ModSec!"
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func main() {
	fmt.Println("start")
	bind := ":3080"
	gmux := mux.NewRouter()
	gmux.HandleFunc("/", HelloFunc).Methods("GET")
	fmt.Println("listen...")
	log.Fatal(http.ListenAndServe(bind, gmux))
}