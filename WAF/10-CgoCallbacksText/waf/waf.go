package callback

// #cgo CFLAGS: -I/usr/local/modsecurity/include
// #cgo LDFLAGS: -L/usr/local/modsecurity/lib/ -Wl,-rpath -Wl,/usr/local/modsecurity/lib/ -lmodsecurity
// #include "waf.h"
import "C"

import (
	"fmt"
)

func Run() {
	fmt.Println("Go Run beg")
	C.pass_GoAdd()
	fmt.Println("Go Run end")
}

//export GoAdd
func GoAdd(a, b int) int {
	fmt.Printf("Go GoAdd beg (x,y)=(%d,%d)", a, b)
	fmt.Println()
	sum := a + b
	fmt.Printf("Go GoAdd end [%d]", sum)
	fmt.Println()
	return sum
}

//export GoText
func GoText() {
	fmt.Printf("Go GoText beg")
	fmt.Println()
	fmt.Printf("Go GoText end")
	fmt.Println()
}
