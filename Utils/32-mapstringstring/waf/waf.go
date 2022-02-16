package waf

// #cgo CFLAGS: -I/usr/local/modsecurity/include
// #cgo LDFLAGS: -L/usr/local/modsecurity/lib/ -Wl,-rpath -Wl,/usr/local/modsecurity/lib/ -lmodsecurity
// #include "waf.h"
import "C"
import "fmt"

func LoadModSecurityCoreRuleSet(reqHeaderKeys, reqHeaderValues []string) int {

	fmt.Println(len(reqHeaderKeys))
	fmt.Println(len(reqHeaderValues))

	return 7
}
