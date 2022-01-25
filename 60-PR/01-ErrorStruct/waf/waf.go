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

func LoadModSecurityCoreRuleSet(filename string) int {

	Cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(Cfilename))

	//imgInfo := C.struct_ImgInfo{imgPath: C.CString()}
	imgInfo := C.struct_ImgInfo{}
	defer C.free(unsafe.Pointer(imgInfo.imgPath))
	C.printStruct(&imgInfo)

	//msg := "test"	//C.GoString()
	//Cpayload *C.char :=
	msg := C.GoString(imgInfo.imgPath)
	if len(msg) > 0 {
		log.Errorf("Msg : '%s'", msg)
	}

	//C.LoadModSecurityCoreRuleSet(Cfilename)

	log.Info(filename)
	return 7
}
