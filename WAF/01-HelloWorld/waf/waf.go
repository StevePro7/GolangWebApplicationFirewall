package waf

// #cgo CFLAGS: -I/usr/local/modsecurity/include
// #cgo LDFLAGS: -L/usr/local/modsecurity/lib/ -Wl,-rpath -Wl,/usr/local/modsecurity/lib/ -lmodsecurity
// #include "waf.h"
import "C"
import (
	"io/ioutil"
	"log"
	"strings"
	"time"
	"unsafe"
)

// Directory where the Core Rules Set are stored.
var rulesetDirectory string

// 03.
func LoadModSecurityCoreRuleSet(filenames []string) int {

	size := len(filenames)
	log.Printf("WAF Attempt load %d Core Rule Sets", size)

	index := loadModSecurityCoreRuleSetImpl(filenames, size)
	if index == size {
		log.Printf("WAF Process load %d Core Rule Sets  SUCCESS", size)
	} else {
		badFile := filenames[index]
		log.Fatalf("WAF Process load %d Core Rule Sets  FAILED!  Bad File: '%s'", size, badFile)
	}

	return index
}

func loadModSecurityCoreRuleSetImpl(filenames []string, size int) int {

	// Transfer core rule set file names to WAF wrapper code.
	csize := C.int(size)
	carray := C.makeCharArray(csize)
	defer C.freeCharArray(carray, csize)
	for index, filename := range filenames {
		C.setArrayString(carray, C.CString(filename), C.int(index))
	}

	// Finally, load ModSecurity core rule set from WAF wrapper code.
	return int(C.LoadModSecurityCoreRuleSet(carray, csize))
}

// 02.
func ExtractRulesSetFilenames() []string {

	// Read all core rule set file names from rules directory.
	var files []string
	items, _ := ioutil.ReadDir(rulesetDirectory)

	log.Printf("WAF Found %d Core Rules Sets", len(items))
	for _, item := range items {

		file := rulesetDirectory + item.Name()
		files = append(files, file)
		log.Printf("WAF Found Rule('%s')", file)
	}

	return files
}

// 01.
func InitializeModSecurity(directory string) {
	rulesetDirectory = directory

	// Ensure rules directory ends with trailing slash.
	if !strings.HasSuffix(rulesetDirectory, "/") {
		rulesetDirectory = rulesetDirectory + "/"
	}

	log.Printf("WAF Core Rules Set directory: '%s'", rulesetDirectory)
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

func GetRulesDirectory() string {
	return rulesetDirectory
}
