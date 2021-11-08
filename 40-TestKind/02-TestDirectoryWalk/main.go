package main

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

	root := "/etc/config/"
	//root := "./config/"
	items, _ := ioutil.ReadDir(root)
	for _, item := range items {
		//if item.IsDir() {
		//	subFolder := root + item.Name() + "/"
		//	subitems, _ := ioutil.ReadDir(subFolder)
		//	for _, subitem := range subitems {
		//		if !subitem.IsDir() {
		//			// handle file there
		//			subitemName := subFolder + subitem.Name()
		//			//fmt.Println(item.Name() + "/" + subitemName)
		//			fmt.Println(subitemName)
		//		}
		//	}
		//} else {
		// handle file there
		//fileName := item.Name()
		fileName := root + item.Name()
		fmt.Println(fileName)
		//}
	}

	log.Println("Directory walk -end-")

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
