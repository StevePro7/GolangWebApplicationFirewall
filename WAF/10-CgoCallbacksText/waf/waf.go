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
	//C.pass_GoAdd()
	C.pass_GoNum()
	//C.pass_GoText()
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

//export GoNum
func GoNum(x C.int) {
	fmt.Println("Go GoNum beg")
	fmt.Printf("Go GoNum X '%d'", x)
	fmt.Println()
	fmt.Println()
	fmt.Println("Go GoNum end")
}

//export GoText
func GoText(x C.int) {
	fmt.Println("Go GoText beg")
	fmt.Printf("Go GoText X '%d'", x)
	fmt.Println()

	//fmt.Printf("Go GoText bar '%s'", bar)
	fmt.Println()
	fmt.Println("Go GoText end")
}
