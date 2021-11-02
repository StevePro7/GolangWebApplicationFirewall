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
	"net/http/httputil"
	"net/url"
	"time"
	"unsafe"
)

func HomeFunc(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data := "Home ModSec!!!!"
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		_, err := fmt.Fprintf(w, "%s", err.Error())
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}

func getFile() string {
	root := "/etc/config/"
	//root := "./config/"
	var fileName string
	fileName = ""
	items, _ := ioutil.ReadDir(root)
	for _, item := range items {
		if !item.IsDir() {
			fileName = root + item.Name()
		}
	}
	return fileName
}

func FileFunc(w http.ResponseWriter, _ *http.Request) {
	log.Print("FileFunc start")

	log.Println("Directory walk start")

	//root := "/etc/config/"
	////root := "./config/"
	//items, _ := ioutil.ReadDir(root)
	//for _, item := range items {
	//	if item.IsDir() {
	//		subFolder := root + item.Name() + "/"
	//		subitems, _ := ioutil.ReadDir(subFolder)
	//		for _, subitem := range subitems {
	//			if !subitem.IsDir() {
	//				// handle file there
	//				subitemName := subFolder + subitem.Name()
	//				fmt.Println(subitemName)
	//			}
	//		}
	//	} else {
	//		// handle file there
	//		//fileName := item.Name()
	//		fileName := root + item.Name()
	//		fmt.Println(fileName)
	//	}
	//}
	fileName := getFile()
	fmt.Println(fileName)

	log.Println("Directory walk -end-")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data := "File Func...!!!!"
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		_, err := fmt.Fprintf(w, "%s", err.Error())
		if err != nil {
			log.Fatal(err)
			return
		}
	}

	log.Print("FileFunc -end-")
}

func TestFunc(w http.ResponseWriter, r *http.Request) {

	log.Printf("Initialize ModSec start")
	fileName := getFile()
	InitModSec(fileName)
	log.Printf("Initialize ModSec -end-")

	log.Printf("req.URL : \"%s\"", r.URL)
	log.Printf("Methods : \"%s\"", r.Method)

	urlx := r.URL.String()
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	log.Println("body = '" + bodyString + "'..!")

	inter := modsec(urlx, bodyString)
	if inter > 0 {
		log.Printf("==== Mod Security Blocked! ====")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	log.Print("Processing request")
	urly, err := url.Parse("https://www.lightbase.io/freeforlife")
	if err != nil {
		log.Println(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(urly)
	director := proxy.Director
	proxy.Director = func(req *http.Request) {
		director(req)
		req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
		req.Host = req.URL.Host
		req.URL.Path = urly.Path
	}
	log.Print("Serving request")
	proxy.ServeHTTP(w, r)
}

func InitModSec(fileName string) {
	log.Println("initModSec start")
	CfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(CfileName))
	C.MyCInit(CfileName)
	log.Println("initModSec -end-")
}

func modsec(url string, buf string) int {
	log.Println("modsec start ", url)
	Curi := C.CString(url)
	Cbuf := C.CString(buf)
	defer C.free(unsafe.Pointer(Curi))
	defer C.free(unsafe.Pointer(Cbuf))
	start := time.Now()
	inter := int(C.MyCProcess(Curi, Cbuf))
	elapsed := time.Since(start)
	log.Printf("modsec()=%d, elapsed: %s", inter, elapsed)
	log.Println("modsec -end-")
	return inter
}

func main() {
	bind := ":3080"
	log.Println("Start web server on port", bind)

	log.Printf("Initialize routes")
	gmux := mux.NewRouter()
	gmux.HandleFunc("/", HomeFunc).Methods("GET")
	gmux.HandleFunc("/file", FileFunc).Methods("GET")
	gmux.HandleFunc("/test/artists.php", TestFunc).Methods("GET")

	log.Printf("listening...")
	if err := http.ListenAndServe(bind, gmux); err != nil {
		log.Fatalf("unable to start web server: %s", err.Error())
	}
}
