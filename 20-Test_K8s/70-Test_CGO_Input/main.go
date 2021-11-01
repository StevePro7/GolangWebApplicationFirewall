package main

// #include "modsec.c"
import "C"

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"unsafe"
)

func InitModSec(url string) {
	log.Println("initModSec start")
	log.Println(url)

	Curi := C.CString(url)
	defer C.free(unsafe.Pointer(Curi))
	C.MyCInit(Curi)

	//Cfile := C.String(file)
	//defer C.free(unsafe.Pointer(Cfile))
	//C.MyCInit(Cfile)
	log.Println("initModSec -end-")
}

func main() {
	bind := ":8081"
	log.Println("Start web server on port", bind)
	file := "modsecdefault.conf"
	InitModSec(file)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		if err != nil {
			log.Fatal(err)
			return
		}
	})
	log.Fatal(http.ListenAndServe(bind, nil))
}
