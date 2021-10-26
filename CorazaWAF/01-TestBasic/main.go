package main

import "C"
import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jptosso/coraza-waf"
	"github.com/jptosso/coraza-waf/seclang"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
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

func modsec(url string, buf string) int {
	log.Println("modsec start ", url)
	log.Println("URL  ", url)
	log.Println("Body ", buf)
	start := time.Now()

	log.Println("initModSec start")
	waf := coraza.NewWaf()
	parser, _ := seclang.NewParser(waf)

	// Now we parse our rules
	//err := parser.FromFile("modsecdefault.conf")
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	//err = parser.FromFile("crs-setup.conf")
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	err := parser.FromFile("rules/REQUEST-942-APPLICATION-ATTACK-SQLI.conf")
	if err != nil {
		log.Fatal(err)
		return 0
	}
	log.Println("initModSec -end-")

	inter := 0

	// Then we create a transaction and assign some variables
	tx := waf.NewTransaction()
	tx.ProcessConnection("127.0.0.1", 80, "127.0.0.1", 80)
	tx.ProcessUri(url, "GET", "1.1")
	tx.ProcessRequestHeaders()
	_, err = tx.ProcessRequestBody()
	if err != nil {
		log.Fatalln(err)
		return 0
	}

	if it := tx.Interruption; it != nil {
		var action = it.Action
		log.Println(action)
		inter = 1
	}

	// Finally, we check the transaction status
	if tx.Interrupted() {
		fmt.Println("Transaction was interrupted")
	}

	elapsed := time.Since(start)
	log.Printf("modsec()=%d, elapsed: %s", inter, elapsed)
	log.Println("modsec -end-")
	return inter
}

func LimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

	log.Printf("starting smart reverse proxy on [%s]", bind)

	log.Printf("initialize mod sec")
	//InitModSec()

	log.Printf("listening [deny] ??")
	if err := http.ListenAndServe(bind, LimitMiddleware(gmux)); err != nil {
		log.Fatalf("unable to start web server: %s", err.Error())
	}
}
