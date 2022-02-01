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
