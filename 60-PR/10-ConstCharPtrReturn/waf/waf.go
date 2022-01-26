package waf

// #cgo CFLAGS: -I/usr/local/modsecurity/include
// #cgo LDFLAGS: -L/usr/local/modsecurity/lib/ -Wl,-rpath -Wl,/usr/local/modsecurity/lib/ -lmodsecurity
// #include "waf.h"
import "C"
import (
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

func LoadModSecurityCoreRuleSet2(filenames []string) int {

	size := len(filenames)
	log.Infof("WAF Attempt load %d Core Rule Sets", size)

	payload := loadModSecurityCoreRuleSetImpl2(filenames, size)
	result := len(payload)
	if result > 0 {
		log.Infof("WAF Process load %d Core Rule Sets  SUCCESS", size)
	}

	return result
}
func loadModSecurityCoreRuleSetImpl2(filenames []string, size int) string {

	// Transfer core rule set file names to WAF wrapper code.
	csize := C.int(size)
	carray := C.makeCharArray(csize)
	defer C.freeCharArray(carray, csize)

	for index, filename := range filenames {
		C.setArrayString(carray, C.CString(filename), C.int(index))
	}

	// Finally, load ModSecurity core rule set from WAF wrapper code.
	Cpayload := C.LoadModSecurityCoreRuleSet2(carray, csize)
	message := C.GoString(Cpayload)

	if len(message) > 0 {
		//log.Errorf("WAF Error attempt load file '%s' => '%v'", filename, message)
		log.Errorf("WAF Error attempt load file '%v'", message)
		//success = false
	}

	defer C.free(unsafe.Pointer(Cpayload))
	return message
}

//export GoModSecurityLoggingCallback
func GoModSecurityLoggingCallback(Cpayload *C.char) {

	payload := C.GoString(Cpayload)
	log.Infof(payload)
}

//func CleanupModSecurity() {
//	C.CleanupModSecurity()
//	log.Printf("WAF Cleanup Mod Security.")
//}
