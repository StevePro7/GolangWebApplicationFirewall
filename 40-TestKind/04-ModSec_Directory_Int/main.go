package main

// #include "modsec.c"
import "C"

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func HomeFunc(w http.ResponseWriter, _ *http.Request) {
	log.Print("HomeFunc start")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data := "Home Func!!!!"
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		_, err := fmt.Fprintf(w, "%s", err.Error())
		if err != nil {
			log.Fatal(err)
			return
		}
	}

	log.Print("HomeFunc -end-")
}

func TestFunc(w http.ResponseWriter, _ *http.Request) {
	log.Print("TestFunc start")

	log.Println("Directory walk start")
	var files []string

	root := "/etc/config/"
	items, _ := ioutil.ReadDir(root)
	for _, item := range items {

		fileName := root + item.Name()
		files = append(files, fileName)
	}

	log.Println("Directory walk -end-")

	log.Println("iterate start")
	for _, file := range files {
		fmt.Println("next file = '" + file + "'")
	}
	log.Println("iterate -end-")

	//log.Println("Call C code start")
	//csize := C.int(len(files))
	//cargs := C.makeCharArray(csize)
	//defer C.freeCharArray(cargs, csize)
	//for i, s := range files {
	//	C.setArrayString(cargs, C.CString(s), C.int(i))
	//}
	//C.processArrayString(cargs, csize)
	//log.Println("Call C code -end-")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data := "Test Func...!!!!"
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		_, err := fmt.Fprintf(w, "%s", err.Error())
		if err != nil {
			log.Fatal(err)
			return
		}
	}

	log.Print("TestFunc -end-")
}

func main() {
	bind := ":3080"
	log.Println("Start web server on port", bind)

	log.Println("Init routes")
	gmux := mux.NewRouter()
	gmux.HandleFunc("/", HomeFunc).Methods("GET")
	gmux.HandleFunc("/test/artists.php", TestFunc).Methods("GET")

	log.Printf("Listening...")
	if err := http.ListenAndServe(bind, gmux); err != nil {
		log.Fatalf("unable to start web server: %s", err.Error())
	}
}
