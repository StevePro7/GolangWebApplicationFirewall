package waf

// #cgo CFLAGS: -I/usr/local/modsecurity/include
// #cgo LDFLAGS: -L/usr/local/modsecurity/lib/ -Wl,-rpath -Wl,/usr/local/modsecurity/lib/ -lmodsecurity
// #include "waf.h"
import "C"
import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"unsafe"
)

// Directory where the Core Rules Set are stored.
var rulesetDirectory string

const defaultRulesetDirectory = "/etc/waf/"

func InitializeModSecurity() {
	log.Printf("WAF Initialize Mod Security.")
	C.InitializeModSecurity()
}

func DefineRulesSetDirectory(directory string) {
	rulesetDirectory = directory
}

func LoadModSecurityCoreRuleSet(filenames []string) int {

	size := len(filenames)
	load := 0

	log.Infof("WAF Attempt load %d Core Rule Set files", size)
	for _, filename := range filenames {
		success := loadModSecurityCoreRuleSetImpl(filename)
		if success {
			load++
		}
	}

	log.Infof("WAF Process load %d Core Rule Set files  SUCCESS", load)
	if size != load {
		log.Infof("WAF Process load %d Core Rule Set files  FAILURE", size-load)
	}

	return load
}

func loadModSecurityCoreRuleSetImpl(filename string) bool {

	// Assume core rule set file loads OK.
	success := true

	Cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(Cfilename))

	coreRuleSetErrorObject := C.struct_CoreRuleSetErrorObject{}
	defer C.free(unsafe.Pointer(coreRuleSetErrorObject.error_message))

	// Attempt to load core rule set file;
	// any error generated from ModSec will be populated in struct.
	C.LoadModSecurityCoreRuleSet(&coreRuleSetErrorObject, Cfilename)

	var message = C.GoString(coreRuleSetErrorObject.error_message)
	if len(message) > 0 {
		log.Errorf("WAF Error attempt load file '%s' => '%v'", filename, message)
		success = false
	}

	return success
}

//export GoModSecurityLoggingCallback
func GoModSecurityLoggingCallback(Cpayload *C.char) {

	payload := C.GoString(Cpayload)
	fmt.Println(payload)

}
