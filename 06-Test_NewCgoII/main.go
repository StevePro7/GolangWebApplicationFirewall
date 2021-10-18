package main

// #cgo CPPFLAGS: -I/usr/local/modsecurity/include
// #cgo LDFLAGS: /usr/local/modsecurity/lib/libmodsecurity.so
// #include "modsec.c"
import "C"

import (
	"log"
	"time"
	"unsafe"
)

func InitModSec() {
	//log.Println("initModSec start")
	C.MyCInit()
	//log.Println("initModSec -end-")
}

func modsec(url string) int {
	log.Println("modsec start ", url)
	Curi := C.CString(url)
	defer C.free(unsafe.Pointer(Curi))
	start := time.Now()
	inter := int(C.MyCProcess(Curi))
	elapsed := time.Since(start)
	log.Printf("modsec()=%i, elapsed: %s", inter, elapsed)
	log.Println("modsec -end-")
	return inter
}

func LimitMiddleware() {
	//log.Println("LimitMiddleware start")
	var url string
	//url = "http://localhost:3080/test/artists.php"
	url = "http://localhost:3080/test/artists.php?artist=0+div+1+union%23foo*%2F*bar%0D%0Aselect%23foo%0D%0A1%2C2%2Ccurrent_user"

	inter := modsec(url)
	log.Println("Intervention??:", inter)
	//log.Println("LimitMiddleware -end-")
}

func main() {
	//log.Println("start")
	InitModSec()
	LimitMiddleware()
	//log.Println("-end-")
}
