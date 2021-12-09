package waf

// #cgo CFLAGS: -I/usr/local/modsecurity/include
// #cgo LDFLAGS: -L/usr/local/modsecurity/lib/ -Wl,-rpath -Wl,/usr/local/modsecurity/lib/ -lmodsecurity
// #include "waf.h"
import "C"
import (
	"log"
	"strings"
	"time"
	"unsafe"
)

// Root directory where the Core Rules Set are stored.
var rulesetDirectory string

func InitializeModSecurity(directory string) {
	rulesetDirectory = directory

	// Ensure rules directory ends with trailing slash.
	if !strings.HasSuffix(rulesetDirectory, "/") {
		rulesetDirectory = rulesetDirectory + "/"
	}
}

func InitModSec() {
	C.InitModSec()
}

func ProcessHttpRequest(url, httpMethod, httpProtocol, httpVersion string, clientLink string, clientPort int, serverLink string, serverPort int) int {
	log.Println("stevepro modsec start ", url)
	Curi := C.CString(url)
	ChttpMethod := C.CString(httpMethod)
	ChttpProtocol := C.CString(httpProtocol)
	ChttpVersion := C.CString(httpVersion)
	CclientLink := C.CString(clientLink)
	CclientPort := C.int(clientPort)
	CserverLink := C.CString(serverLink)
	CserverPort := C.int(serverPort)

	defer C.free(unsafe.Pointer(Curi))
	defer C.free(unsafe.Pointer(ChttpMethod))
	defer C.free(unsafe.Pointer(ChttpProtocol))
	defer C.free(unsafe.Pointer(ChttpVersion))
	defer C.free(unsafe.Pointer(CclientLink))
	defer C.free(unsafe.Pointer(CserverLink))
	start := time.Now()
	detection := int(C.ProcessHttpRequest(Curi, ChttpMethod, ChttpProtocol, ChttpVersion, CclientLink, CclientPort, CserverLink, CserverPort))

	elapsed := time.Since(start)
	log.Printf("stevepro modsec()=%d, elapsed: %s", detection, elapsed)
	log.Println("stevepro modsec -end-")
	return detection
}

// TODO delete this code... eventually!
func Add(x, y int) int {
	return x + y
}

func GetRulesDirectory() string {
	return rulesetDirectory
}
