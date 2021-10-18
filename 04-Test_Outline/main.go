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

func TestFunc(w http.ResponseWriter, r *http.Request) {
	log.Print("testHandler")
	urlx, err := url.Parse("http://www.lightbase.io/freeforlife")
	if err != nil {
		log.Println(err)
	}
	log.Print(urlx)
	proxy := httputil.NewSingleHostReverseProxy(urlx)
	director := proxy.Director
	proxy.Director = func(req * http.Request) {
		director(req)
		req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
		req.Host = req.URL.Host
		req.URL.Path = urlx.Path
	}
	proxy.ServeHTTP(w, r)
}

func main() {
	log.Println("testing")
	bind := ":3080"
	gmux := mux.NewRouter()
	gmux.HandleFunc("/", HomeFunc).Methods("GET")
	gmux.HandleFunc("/test", TestFunc).Methods("GET")

	admin := mux.NewRouter()
	admin.HandleFunc("/admin", AdminFunc).Methods("GET")
	go func() {
		if err := http.ListenAndServe(":3081", admin); err != nil {
			log.Fatalf("unable to start admin server: %s", err.Error())
		}
	}()

	log.Printf("starting smart reverse proxy on [%s]", bind)
	if err := http.ListenAndServe(bind, gmux); err != nil {
		log.Fatalf("unable to start web server: %s", err.Error())
	}
}