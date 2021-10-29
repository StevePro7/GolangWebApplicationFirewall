package main

import "C"
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
	resp := map[string]interface{}{"status": "smart reverse proxy hello!"}
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
	proxy.ServeHTTP(w, r)
}

func modsec(url string) int {
	log.Printf("mod.sec : \"%s", url)
	inter := 0
	return inter
}

func initModSec() {
	log.Println("initModSec start")
	//C.init()
	log.Println("initModSec -end-")
}

func limitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("req.URL : \"%s\"", r.URL)

		urlx := r.URL.String()
		inter := modsec(urlx)
		//var inter int
		//inter = 0
		//if r.RequestURI == "/test/artists.php" {
		//	inter = 1
		//}

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

	admin := mux.NewRouter()
	admin.HandleFunc("/admin", AdminFunc).Methods("GET")
	go func() {
		if err := http.ListenAndServe(":3081", admin); err != nil {
			log.Fatalf("unable to start admin server: %s", err.Error())
		}
	}()

	log.Printf("starting smart reverse proxy on [%s]", bind)

	log.Printf("initialize mod sec")
	initModSec()

	log.Printf("listening!")
	if err := http.ListenAndServe(bind, limitMiddleware(gmux)); err != nil {
		log.Fatalf("unable to start web server: %s", err.Error())
	}
}
