package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// #include "test.h"
import "C"

func main() {
	fmt.Println("----------------Go to C---------------")
	fmt.Println(C.char('Y'))
	//fmt.Printf("%s\n", C.str('Y'))
	fmt.Println(C.uchar('C'))
	fmt.Println(C.short(254))
	fmt.Println(C.long(11112222))
	goi := 2
	//unsafe.Pointer --> void *
	cpi := unsafe.Pointer(&goi)
	C.printI(cpi)

	fmt.Println("----------------C to Go---------------")
	fmt.Println(C.ch)
	fmt.Println(C.uch)
	fmt.Println(C.st)
	fmt.Println(C.i)
	fmt.Println(C.lt)

	f := float32(C.f)
	fmt.Println(reflect.TypeOf(f))
	fmt.Println(C.f)
	db := float64(C.db)
	fmt.Println(reflect.TypeOf(db))
	fmt.Println(C.db)

	//The difference between constant string and char array is different when converted to Go type
	str := C.GoString(C.str)
	fmt.Println(str)

	fmt.Println(reflect.TypeOf(C.str1))
	var charray []byte
	for i := range C.str1 {
		if C.str1[i] != 0 {
			charray = append(charray, byte(C.str1[i]))
		}
	}

	fmt.Println(charray)
	fmt.Println(string(charray))

	imgInfo := C.struct_ImgInfo{imgPath: C.CString("../images/xx.jpg"), format: 0, width: 500, height: 400}
	defer C.free(unsafe.Pointer(imgInfo.imgPath))
	C.printStruct(&imgInfo)
}
