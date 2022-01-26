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

func LoadModSecurityCoreRuleSet(filenames []string) {

	size := len(filenames)
	load := 0

	log.Infof("WAF Attempt load %d Core Rule Set files", size)

	for _, filenameX := range filenames {
		log.Infof("filname '%v'", filenameX)
	}

	filename := filenames[0]
	success := loadModSecurityCoreRuleSetImpl(filename)
	if success {
		load++
	}

	log.Infof("WAF Process load %d Core Rule Set files  SUCCESS", load)
	if size != load {
		log.Infof("WAF Process load %d Core Rule Set files  FAILURE", size-load)
	}

}

func loadModSecurityCoreRuleSetImpl(filename string) bool {

	// Assume core rule set file will load OK.
	success := true

	Cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(Cfilename))

	imgInfo := C.struct_ImgInfo{}
	defer C.free(unsafe.Pointer(imgInfo.imgPath))

	C.LoadModSecurityCoreRuleSet(&imgInfo, Cfilename)

	msg := C.GoString(imgInfo.imgPath)
	if len(msg) > 0 {
		log.Errorf("WAF issue loading file '%s'", msg)
		success = false
	}

	log.Info(filename)
	return success
}
