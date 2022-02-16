package waf

// #cgo CFLAGS: -I/usr/local/modsecurity/include
// #cgo LDFLAGS: -L/usr/local/modsecurity/lib/ -Wl,-rpath -Wl,/usr/local/modsecurity/lib/ -lmodsecurity
// #include "waf.h"
import "C"
import "fmt"

func LoadModSecurityCoreRuleSet(reqHeaderKeys, reqHeaderVals []string) int {

	size := len(reqHeaderKeys)

	CreqHeaderSize := C.int(size)
	CreqHeaderKeys := C.makeCharArray(CreqHeaderSize)
	defer C.freeCharArray(CreqHeaderKeys, CreqHeaderSize)
	CreqHeaderVals := C.makeCharArray(CreqHeaderSize)
	defer C.freeCharArray(CreqHeaderVals, CreqHeaderSize)

	for index, reqHeaderKey := range reqHeaderKeys {
		C.setArrayString(CreqHeaderKeys, C.CString(reqHeaderKey), C.int(index))
	}

	for index, reqHeaderVal := range reqHeaderVals {
		C.setArrayString(CreqHeaderVals, C.CString(reqHeaderVal), C.int(index))
	}

	result := C.AddRequestHeaders(CreqHeaderKeys, CreqHeaderVals, CreqHeaderSize)
	answer := int(result)

	fmt.Println(len(reqHeaderKeys))
	fmt.Println(len(reqHeaderVals))

	return answer
}
