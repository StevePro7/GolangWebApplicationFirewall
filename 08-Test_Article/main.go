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

func modsec(url string) int {
	Curi := C.CString(url)
	defer C.free(unsafe.Pointer(Curi))
	start := time.Now()
	inter := int(C.process(Curi))
	elapsed := time.Since(start)
	log.Printf("\n========\nmodsec()=%d, elapsed: %s", inter, elapsed)
	return inter
}

func initModSec() {
	C.init()
}

func limitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("req.URL %s", r.URL)
		inter := modsec(r.URL.String())
		if inter > 0 {
			log.Printf("==== Mod Security Blocked! ====")
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func adminHelloHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := map[string]interface{}{"status": "smart reverse proxy admin access - hello!"}
	out, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
	}
	fmt.Fprintf(w, string(out))
}

func testHandler(w http.ResponseWriter, req *http.Request) {
	log.Print("testHandler")
	url, err := url.Parse("http://www.lightbase.io/freeforlife")
	if err != nil {
		log.Println(err)
	}
	log.Print(url)
	proxy := httputil.NewSingleHostReverseProxy(url)
	director := proxy.Director
	proxy.Director = func(req *http.Request) {
		director(req)
		req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
		req.Host = req.URL.Host
		req.URL.Path = url.Path
	}
	proxy.ServeHTTP(w, req)
}

func main() {
	bind := ":3080"
	adminbind := "localhost:3081"
	gmux := mux.NewRouter()
	gmux.HandleFunc("/test", testHandler).Methods("GET")
	adminmux := mux.NewRouter()
	adminmux.HandleFunc("/test", adminHelloHandler).Methods("GET")
	go func() {
		if err := http.ListenAndServe(adminbind, adminmux); err != nil {
			log.Fatalf("unable to start server: %s", err.Error())
		}
	}()
	log.Printf("starting smart reverse proxy on [%s], administrative endpoint: [%s] with / test ", bind, adminbind)
	initModSec()
	if err := http.ListenAndServe(bind, limitMiddleware(gmux)); err != nil {
		log.Fatalf("unable to start server: %s", err.Error())
	}
	return
}
