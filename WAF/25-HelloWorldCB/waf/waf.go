package waf

// #cgo CFLAGS: -I/usr/local/modsecurity/include
// #cgo LDFLAGS: -L/usr/local/modsecurity/lib/ -Wl,-rpath -Wl,/usr/local/modsecurity/lib/ -lmodsecurity
// #include "waf.h"
import "C"
import (
	"github.com/google/uuid"
	"io/ioutil"
	"strings"
	"time"
	"unsafe"

	log "github.com/sirupsen/logrus"
)

// Directory where the Core Rules Set are stored.
var rulesetDirectory string

func InitializeModSecurity() {
	C.InitializeModSecurity()
}

func DefineRulesSetDirectory(directory string) {
	rulesetDirectory = directory
	log.Printf("WAF Core Rules Set directory: '%s'", rulesetDirectory)

	// Ensure rules directory ends with trailing slash.
	if !strings.HasSuffix(rulesetDirectory, "/") {
		rulesetDirectory = rulesetDirectory + "/"
	}
}

func ExtractRulesSetFilenames() []string {

	// Read all core rule set file names from rules directory.
	var files []string
	items, _ := ioutil.ReadDir(rulesetDirectory)
	count := 1
	for _, item := range items {

		// Ignore symbolically linked files!
		filename := item.Name()
		if strings.HasPrefix(filename, "..") {
			continue
		}

		file := rulesetDirectory + filename
		files = append(files, file)
		log.Infof("WAF Found Rule[%d]('%s')", count, file)
		count++
	}

	log.Infof("WAF Total Core Rules Sets: %d", len(files))
	return files
}

func LoadModSecurityCoreRuleSet(filenames []string) int {

	size := len(filenames)
	log.Infof("WAF Attempt load %d Core Rule Sets", size)

	index := loadModSecurityCoreRuleSetImpl(filenames, size)
	if index == size {
		log.Infof("WAF Process load %d Core Rule Sets  SUCCESS", size)
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

func GenerateModSecurtityID() string {
	return uuid.New().String()
}

func ProcessHttpRequest(id, url, httpMethod, httpProtocol, httpVersion string, clientLink string, clientPort int, serverLink string, serverPort int) int {
	log.Printf("WAF Process Http Request URL '%s'", url)
	Cid := C.CString(id)
	Curi := C.CString(url)
	ChttpMethod := C.CString(httpMethod)
	ChttpProtocol := C.CString(httpProtocol)
	ChttpVersion := C.CString(httpVersion)
	CclientLink := C.CString(clientLink)
	CclientPort := C.int(clientPort)
	CserverLink := C.CString(serverLink)
	CserverPort := C.int(serverPort)

	defer C.free(unsafe.Pointer(Cid))
	defer C.free(unsafe.Pointer(Curi))
	defer C.free(unsafe.Pointer(ChttpMethod))
	defer C.free(unsafe.Pointer(ChttpProtocol))
	defer C.free(unsafe.Pointer(ChttpVersion))
	defer C.free(unsafe.Pointer(CclientLink))
	defer C.free(unsafe.Pointer(CserverLink))

	start := time.Now()
	detection := int(C.ProcessHttpRequest(Cid, Curi, ChttpMethod, ChttpProtocol, ChttpVersion, CclientLink, CclientPort, CserverLink, CserverPort))
	elapsed := time.Since(start)

	log.Infof("WAF Process Http Request URL '%s' Detection=%d Time elapsed: %s", url, detection, elapsed)
	return detection
}

// GetRulesDirectory public helper function for testing.
func GetRulesDirectory() string {
	return rulesetDirectory
}

//export GoModSecurityLoggingCallback
func GoModSecurityLoggingCallback(x *C.char) {
	log.Info("WAF Go GoText beg")

	var y string
	y = C.GoString(x)
	//fmt.Printf("Go GoText X '%s'", x)
	log.Infof("WAF Go GoText X '%s' [oof]", y)
	log.Info()

	//fmt.Printf("Go GoText bar '%s'", bar)
	log.Info("WAF Go GoText end")
}
