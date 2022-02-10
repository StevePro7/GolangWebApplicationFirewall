package main

import (
	"fmt"
	"log"
)

// #include "hello.h"
import "C"

func main() {
	fmt.Println("start!")

	_, err := C.Hello()
	if err != nil {
		log.Fatal(err)
	}

	obj, err2 := C.msc_intervention()
	if err2 != nil {
		log.Fatal(err2)
	}
	val := int(obj)
	fmt.Printf("VAL = %d\n", val)
	fmt.Println("-end-")
}
