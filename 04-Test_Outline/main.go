package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HomeFunc(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data := "Home ModSec!"
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		_, err := fmt.Fprintf(w, "%s", err.Error())
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}

func AdminFunc(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := map[string]interface{} {"status": "smart reverse proxy hello!"}
	out, err := json.Marshal(resp)
	if err != nil {
		log.Fatal(err)
	}
	_, err = fmt.Fprintf(w, string(out))
	if err != nil {
		log.Fatal(err)
		return
	}
}

func main() {
	fmt.Println("start")
	bind := ":3080"
	gmux := mux.NewRouter()
	gmux.HandleFunc("/", HomeFunc).Methods("GET")

	admin := mux.NewRouter()
	admin.HandleFunc("/admin", AdminFunc).Methods("GET")
	go func() {
		if err := http.ListenAndServe(":3081", admin); err != nil {
			log.Fatalf("unable to start admin server: %s", err.Error())
		}
	}()

	fmt.Println("listen...")
	log.Fatal(http.ListenAndServe(bind, gmux))
}