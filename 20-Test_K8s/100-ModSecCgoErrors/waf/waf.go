package waf

// #cgo CFLAGS: -I/usr/local/modsecurity/include
// #cgo LDFLAGS: -L/usr/local/modsecurity/lib/ -Wl,-rpath -Wl,/usr/local/modsecurity/lib/ -lmodsecurity
// #include "waf.h"
import "C"
import (
	"errors"
	"fmt"
	"log"
	"unsafe"
)

func InitializeModSecurity() {
	log.Printf("WAF Initialize Mod Security.")
	C.InitializeModSecurity()
}

func LoadModSecurityCoreRuleSet(filename string) error {
	Cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(Cfilename))

	// Attempt to load core rule set file;
	// any error generated from ModSec will be returned directly.
	Cpayload := C.LoadModSecurityCoreRuleSet(Cfilename)
	if Cpayload != nil {
		errStr := C.GoString(Cpayload)
		C.free(unsafe.Pointer(Cpayload))

		if len(errStr) > 0 {
			errMsg := fmt.Sprintf("WAF Error attempt load file '%s' => '%v'", filename, errStr)
			return errors.New(errMsg)
		}
	}

	return nil
}

func CleanupModSecurity() {
	C.CleanupModSecurity()
	log.Printf("WAF Cleanup Mod Security.")
}

func ProcessHttpRequest(id, url, httpMethod, httpProtocol, httpVersion string, clientHost string, clientPort uint32, serverHost string, serverPort uint32) error {

	Cid := C.CString(id)
	defer C.free(unsafe.Pointer(Cid))
	Curi := C.CString(url)
	defer C.free(unsafe.Pointer(Curi))
	ChttpMethod := C.CString(httpMethod)
	defer C.free(unsafe.Pointer(ChttpMethod))
	ChttpProtocol := C.CString(httpProtocol)
	defer C.free(unsafe.Pointer(ChttpProtocol))
	ChttpVersion := C.CString(httpVersion)
	defer C.free(unsafe.Pointer(ChttpVersion))
	CclientHost := C.CString(clientHost)
	defer C.free(unsafe.Pointer(CclientHost))
	CserverHost := C.CString(serverHost)
	defer C.free(unsafe.Pointer(CserverHost))
	CclientPort := C.int(clientPort)
	CserverPort := C.int(serverPort)

	result := C.ProcessHttpRequest(Cid, Curi, ChttpMethod, ChttpProtocol, ChttpVersion, CclientHost, CclientPort, CserverHost, CserverPort)
	verdict := int(result)
	if verdict > 0 {
		return errors.New("uh oh")
	}
	return nil
}
