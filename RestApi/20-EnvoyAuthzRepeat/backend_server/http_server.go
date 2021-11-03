package main

import (
	"fmt"
	"golang.org/x/net/http2"
	"log"
	"net/http"
	"net/http/httputil"
)

//var ()

func fronthandler(w http.ResponseWriter, r *http.Request) {

	log.Println("/ called")
	fmt.Println("/ called")

	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(requestDump))
	fmt.Println(string(requestDump))

	w.Header().Set("X-Custom-Header-From-Backend", "from backend")
	log.Println(w, "ok")
	_, err = fmt.Fprint(w, "ok")
	if err != nil {
		log.Fatalln(err)
		return
	}
}
func main() {
	http.HandleFunc("/", fronthandler)

	var server *http.Server
	server = &http.Server{
		Addr: ":11000",
	}
	err := http2.ConfigureServer(server, &http2.Server{})
	if err != nil {
		log.Fatalln(err)
		return
	}

	fmt.Println("Starting backend server...")
	log.Println("Starting backend server...")

	err = server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
		return
	}
}
