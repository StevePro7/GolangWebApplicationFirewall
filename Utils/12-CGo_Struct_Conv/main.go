package main

// #include "test.h"
import "C"

import "fmt"

func main() {
	fmt.Println("beg")
	cx := C.int(1)
	cy := C.int(2)
	sum := C.int(C.MyCAdd(cx, cy))
	fmt.Println(sum)
	fmt.Println("end")
}
