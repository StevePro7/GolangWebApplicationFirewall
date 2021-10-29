package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func adminHelloHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := map[string]interface{}{"status": "smart reverse proxy admin access - hello!"}
	out, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
	}
	_, err = fmt.Fprintf(w, string(out))
	if err != nil {
		return
	}
}

func testHandler(w http.ResponseWriter, req *http.Request) {
	log.Print("testHandler")
	urlx, err := url.Parse("http://www.lightbase.io/freeforlife")
	if err != nil {
		log.Println(err)
	}
	log.Print(urlx)
	proxy := httputil.NewSingleHostReverseProxy(urlx)
	director := proxy.Director
	proxy.Director = func(req *http.Request) {
		director(req)
		req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
		req.Host = req.URL.Host
		req.URL.Path = urlx.Path
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

	log.Printf("starting smart reverse proxy on [%s], admin endpoint: [%s] with /test", bind, adminbind)

	//TODO
	//initmodsec()
	//if err := http.ListenAndServe(bind, limitMiddelware(gmux)); err != nil {
	//	log.Fatalf("unable to start server: %s", err.Error())
	//}
	return
}
