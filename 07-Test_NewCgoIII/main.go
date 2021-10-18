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
	log.Print("Serving request")
	proxy.ServeHTTP(w, r)
}

func InitModSec() {
	//log.Println("initModSec start")
	C.MyCInit()
	//log.Println("initModSec -end-")
}

func modsec(url string) int {
	//log.Println("modsec start ", url)
	Curi := C.CString(url)
	defer C.free(unsafe.Pointer(Curi))
	start := time.Now()
	inter := int(C.MyCProcess(Curi))
	elapsed := time.Since(start)
	log.Printf("modsec()=%d, elapsed: %s", inter, elapsed)
	//log.Println("modsec -end-")
	return inter
}

func LimitMiddleware(next http.Handler) http.Handler {
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

//func LimitMiddleware() {
//	//log.Println("LimitMiddleware start")
//	var url string
//	//url = "http://localhost:3080/test/artists.php"
//	url = "http://localhost:3080/test/artists.php?artist=0+div+1+union%23foo*%2F*bar%0D%0Aselect%23foo%0D%0A1%2C2%2Ccurrent_user"
//
//	inter := modsec(url)
//	log.Println("Intervention??:", inter)
//	//log.Println("LimitMiddleware -end-")
//}

func main() {
	//log.Println("start")
	//InitModSec()
	//LimitMiddleware()
	//log.Println("-end-")

	log.Println("testing")
	bind := ":3080"
	gmux := mux.NewRouter()
	gmux.HandleFunc("/", HomeFunc).Methods("GET")
	//gmux.HandleFunc("/test/artists.php", TestFunc).Methods("GET")
	gmux.HandleFunc("/test", TestFunc).Methods("GET")

	admin := mux.NewRouter()
	admin.HandleFunc("/admin", AdminFunc).Methods("GET")
	go func() {
		if err := http.ListenAndServe(":3081", admin); err != nil {
			log.Fatalf("unable to start admin server: %s", err.Error())
		}
	}()

	log.Printf("starting smart reverse proxy on [%s]", bind)

	log.Printf("initialize mod sec")
	InitModSec()

	log.Printf("listening!!")
	if err := http.ListenAndServe(bind, LimitMiddleware(gmux)); err != nil {
		log.Fatalf("unable to start web server: %s", err.Error())
	}
}
