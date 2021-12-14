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
	//C.pass_GoNum()
	//C.pass_GoText()
	C.pass_GoConst()
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
func GoText(x *C.char) {
	fmt.Println("Go GoText beg")

	var y string
	//y = "hello"
	y = C.GoString(x)
	//fmt.Printf("Go GoText X '%s'", x)
	fmt.Printf("Go GoText X '%s'", y)
	fmt.Println()

	//fmt.Printf("Go GoText bar '%s'", bar)
	fmt.Println()
	fmt.Println("Go GoText end")
}

//export GoConst
func GoConst(x *C.char, y *C.char) {
	fmt.Println("Go GoConst beg")

	//var y string
	//y = "hello"
	a := C.GoString(x)
	b := C.GoString(y)
	//fmt.Printf("Go GoText X '%s'", x)
	fmt.Printf("Go GoConst A '%s'", a)
	fmt.Println()
	fmt.Printf("Go GoConst B '%s'", b)
	fmt.Println()

	//fmt.Printf("Go GoText bar '%s'", bar)
	//fmt.Println()
	fmt.Println("Go GoConst end")
}
