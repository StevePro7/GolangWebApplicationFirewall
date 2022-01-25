package main

// #include "test.h"
import "C"

import (
	"fmt"
	"unsafe"
)

//type X struct {
//	Y, Z int32
//}
//
//type Bob struct {
//	Sgb int32
//}

/*type Fish struct {
	Temp string
}
*/

func main() {
	//fmt.Println("beg")

	//fmt.Println("beg II")
	//a := &X{Y: 2, Z: 3}
	//s := C.sum(*((*C.struct_x)(unsafe.Pointer(a))))
	//fmt.Println(s)
	//fmt.Println("end II")

	//fmt.Println("beg III")
	//b := Bob{ Sgb: 9 }
	//C.test((*C.struct_bob)(unsafe.Pointer(&b)))
	//sgb := b.Sgb
	//fmt.Println(sgb)
	//fmt.Println("end III")

	fmt.Println("beg IV")

	imgInfo := C.struct_ImgInfo{imgPath: C.CString("stevepro")}
	defer C.free(unsafe.Pointer(imgInfo.imgPath))
	C.printStruct(&imgInfo)

	//ff := Fish{Temp: "blah"}
	//C.fishing((*C.struct_fush)(unsafe.Pointer(&ff)))
	////sgb := b.Sgb
	//fmt.Println(ff.Temp)
	fmt.Println("end IV")

	//cx := C.int(1)
	//cy := C.int(2)
	//sum := C.int(C.MyCAdd(cx, cy))
	//fmt.Println(sum)

	//	fmt.Println("end")
}
