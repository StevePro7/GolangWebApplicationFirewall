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

func InitializeModSecurity() error {
	log.Printf("WAF Initialize Mod Security.")
	_, err := C.InitializeModSecurity()
	if err != nil {
		return err
	}

	return nil
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

func CleanupModSecurity() error {
	_, err := C.CleanupModSecurity()
	if err != nil {
		return err
	}

	log.Printf("WAF Cleanup Mod Security.")
	return nil
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

	retVal, err := C.ProcessHttpRequest(Cid, Curi, ChttpMethod, ChttpProtocol, ChttpVersion, CclientHost, CclientPort, CserverHost, CserverPort)
	if err != nil {
		errMsg2 := fmt.Sprintf("ModSec error '%v'", err.Error())
		return errors.New(errMsg2)
	}

	detection := int(retVal)

	log.Println("")
	log.Printf("URL '%s' Detection=%d", url, detection)
	log.Println("")
	if detection > 0 {
		errMsg := fmt.Sprintf("ERROR URL '%s' Detection=%d", url, detection)
		return errors.New(errMsg)
	}

	return nil
}

func StevePro() error {
	result, err := C.StevePro()
	if err != nil {
		return err
	}
	verdict := int(result)
	log.Printf("Verdict %d\n", verdict)
	return nil
}
