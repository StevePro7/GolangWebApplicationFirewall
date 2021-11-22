package main

// #cgo CPPFLAGS: -I/usr/local/modsecurity/include
// #cgo LDFLAGS: /usr/local/modsecurity/lib/libmodsecurity.so
// #include "modsec.c"
import "C"

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"unsafe"
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

func modsec(url, httpMethod, httpProtocol, httpVersion string, clientLink string, clientPort int, serverLink string, serverPort int) int {
	log.Println("modsec start ", url)
	Curi := C.CString(url)
	ChttpMethod := C.CString(httpMethod)
	ChttpProtocol := C.CString(httpProtocol)
	ChttpVersion := C.CString(httpVersion)
	CclientLink := C.CString(clientLink)
	CclientPort := C.int(clientPort)
	CserverLink := C.CString(serverLink)
	CserverPort := C.int(serverPort)

	defer C.free(unsafe.Pointer(Curi))
	defer C.free(unsafe.Pointer(ChttpMethod))
	defer C.free(unsafe.Pointer(ChttpProtocol))
	defer C.free(unsafe.Pointer(ChttpVersion))
	defer C.free(unsafe.Pointer(CclientLink))
	defer C.free(unsafe.Pointer(CserverLink))

	start := time.Now()
	inter := int(C.MyCProcess(Curi, ChttpMethod, ChttpProtocol, ChttpVersion, CclientLink, CclientPort, CserverLink, CserverPort))
	elapsed := time.Since(start)
	log.Printf("modsec()=%d, elapsed: %s", inter, elapsed)
	log.Println("modsec -end-")
	return inter
}

func TestFunc(w http.ResponseWriter, r *http.Request) {
	log.Print("TestFunc start")

	// 01. Directory walk
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

	// 02. Load rules via C code
	log.Println("Call C code start")
	csize := C.int(len(files))
	cargs := C.makeCharArray(csize)
	defer C.freeCharArray(cargs, csize)
	for i, s := range files {
		C.setArrayString(cargs, C.CString(s), C.int(i))
	}
	C.processArrayString(cargs, csize)
	log.Println("Call C code -end-")

	// 03. ModSec check
	log.Printf("req.URL : \"%s\"", r.URL)
	log.Printf("Methods : \"%s\"", r.Method)

	uri := r.URL.String()
	httpMethod := "GET"

	//protocol := "HTTP/1.1"
	httpProtocol := "HTTP"
	httpVersion := "1.1"

	//clientSocket := "127.0.0.1:80"
	clientLink := "127.0.0.1"
	clientPort := 80
	//serverSocket := "127.0.0.1:80"
	serverLink := "127.0.0.1"
	serverPort := 80

	inter := modsec(uri, httpMethod, httpProtocol, httpVersion, clientLink, clientPort, serverLink, serverPort)
	if inter > 0 {
		log.Printf("==== Mod Security Blocked! ====")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data := "Test Func...!!!! ABC"
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
