package main

import "C"
import (
	"log"
	"net/http"
	"time"
	"unsafe"
)

func modsec(url string) int {
	Curi := C.CString(url)
	defer C.free(unsafe.Pointer(Curi))
	start := time.Now()
	inter := int(C.proces(Curi))
	elapsed := time.Since(start)
	log.Printf("modsec()=%d, elapsed: %s", inter, elapsed)
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
			log.Printf("Mo9de Security Blocked")
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	})
}
