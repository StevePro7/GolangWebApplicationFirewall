package main

// #cgo CPPFLAGS: -I/usr/local/modsecurity/include
// #cgo LDFLAGS: /usr/local/modsecurity/lib/libmodsecurity.so
// #include "modsec.c"
import "C"

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func InitModSec() {
	log.Println("cAPI initModSec start")
	age := int(C.MyCInit())
	log.Println("cAPI AGE :", age)
	log.Println("cAPI initModSec -end-")
}

func HomeFunc(w http.ResponseWriter, r *http.Request) {
	log.Println("yAPI Home -start")
	log.Println("yAPI url:", r.URL)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data := "yAPI Home function"
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		_, err := fmt.Fprintf(w, "%s", err.Error())
		if err != nil {
			log.Fatalln(err)
		}
	}
	log.Println("yAPI Home --end-")
}

func TestFunc(w http.ResponseWriter, r *http.Request) {
	log.Println("yAPI Test -start")
	log.Println("yAPI url:", r.URL)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data := "yAPI Test function"
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		_, err := fmt.Fprintf(w, "%s", err.Error())
		if err != nil {
			log.Fatalln(err)
		}
	}
	log.Println("yAPI Test --end-")
}

func main() {
	log.Println("yAPI start [3.0]")
	bind := ":3080"
	InitModSec()
	gmux := mux.NewRouter()
	gmux.HandleFunc("/", HomeFunc).Methods("GET")
	gmux.HandleFunc("/test/artists.php", TestFunc).Methods("GET")
	log.Printf("yAPI listen on [%s]\n", bind)
	log.Fatal(http.ListenAndServe(bind, gmux))
}
