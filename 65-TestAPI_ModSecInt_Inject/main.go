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

func TestFunc(w http.ResponseWriter, r *http.Request) {
	log.Print("Processing request")
	urlx, err := url.Parse("https://www.lightbase.io/freeforlife")
	if err != nil {
		log.Println(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(urlx)
	director := proxy.Director
	proxy.Director = func(req *http.Request) {
		director(req)
		req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
		req.Host = req.URL.Host
		req.URL.Path = urlx.Path
	}
	log.Print("Serving request")
	proxy.ServeHTTP(w, r)
}

func InitModSec() {
	//log.Println("initModSec start")
	C.MyCInit()
	//log.Println("initModSec -end-")
}

func modsec(url string, httpMethod string, httpProtocol string, httpVersion string) int {
	log.Println("modsec start ", url)
	Curi := C.CString(url)
	ChttpMethod := C.CString(httpMethod)
	ChttpProtocol := C.CString(httpProtocol)
	ChttpVersion := C.CString(httpVersion)

	defer C.free(unsafe.Pointer(Curi))
	defer C.free(unsafe.Pointer(ChttpMethod))
	defer C.free(unsafe.Pointer(ChttpProtocol))
	defer C.free(unsafe.Pointer(ChttpVersion))

	start := time.Now()
	inter := int(C.MyCProcess(Curi, ChttpMethod, ChttpProtocol, ChttpVersion))
	elapsed := time.Since(start)
	log.Printf("modsec()=%d, elapsed: %s", inter, elapsed)
	log.Println("modsec -end-")
	return inter
}

func LimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("req.URL : \"%s\"", r.URL)
		log.Printf("Methods : \"%s\"", r.Method)

		uri := r.URL.String()
		httpMethod := "GET"
		//protocol := "HTTP/1.1"
		httpProtocol := "HTTP"
		httpVersion := "1.1"

		//clientSocket := "127.0.0.1:80"
		//serverSocket := "127.0.0.1:80"
		//clientPort := C.int(80)

		inter := modsec(uri, httpMethod, httpProtocol, httpVersion)
		if inter > 0 {
			log.Printf("==== Mod Security Blocked! ====")
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	log.Println("testing")
	bind := ":3080"
	gmux := mux.NewRouter()
	gmux.HandleFunc("/", HomeFunc).Methods("GET")
	gmux.HandleFunc("/test/artists.php", TestFunc).Methods("GET")

	log.Printf("starting smart reverse proxy on [%s]", bind)

	log.Printf("initialize mod sec")
	InitModSec()

	log.Printf("listening [deny]...")
	if err := http.ListenAndServe(bind, LimitMiddleware(gmux)); err != nil {
		log.Fatalf("unable to start web server: %s", err.Error())
	}
}
