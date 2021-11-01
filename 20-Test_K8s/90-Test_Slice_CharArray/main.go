package main

// #include "test.c"
import "C"
import "fmt"

func main() {
	fmt.Println("start")

	var sargs = []string{"one", "two", "three"}
	for _, v := range sargs {
		fmt.Println(v)
	}

	cargs := C.makeCharArray(C.int(len(sargs)))
	defer C.freeCharArray(cargs, C.int(len(sargs)))
	for i, s := range sargs {
		C.setArrayString(cargs, C.CString(s), C.int(i))
	}

	fmt.Println("-end-")
}
