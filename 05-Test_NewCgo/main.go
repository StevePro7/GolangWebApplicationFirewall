package main

import (
	"log"
)

// #cgo CPPFLAGS: -I/usr/local/modsecurity/include
// #cgo LDFLAGS: /usr/local/modsecurity/lib/libmodsecurity.so
// #include "modsec.c"
import "C"

func initModSec() {
	log.Println("initModSec! start")
	C.MyCInit()
	log.Println("initModSec! -end-")
}

func main() {
	log.Println("start")
	initModSec()
	log.Println("-end-")
}