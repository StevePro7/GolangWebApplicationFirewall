package main

// #include "test.c"
import "C"
import "fmt"

func main() {
	fmt.Println("start")

	var sargs = []string{"one", "two", "three", "four", "five"}
	for _, v := range sargs {
		fmt.Println(v)
	}

	csize := C.int(len(sargs))
	//cargs := C.makeCharArray(C.int(len(sargs)))
	cargs := C.makeCharArray(csize)
	//defer C.freeCharArray(cargs, C.int(len(sargs)))
	defer C.freeCharArray(cargs, csize)
	for i, s := range sargs {
		C.setArrayString(cargs, C.CString(s), C.int(i))
	}

	C.processArrayString(cargs, csize)
	fmt.Println("-end-")
}
