package main

// #cgo CPPFLAGS: -I/usr/local/modsecurity/include
// #cgo LDFLAGS: /usr/local/modsecurity/lib/libmodsecurity.so
// #include "modsec.c"
import "C"

import "log"

func InitModSec() {
	log.Println("initModSec start")
	C.MyCInit()
	log.Println("initModSec -end-")
}

func main() {
	log.Println("start")
	InitModSec()
	log.Println("-end-!")
}