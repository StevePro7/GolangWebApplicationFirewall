package main

import (
	"fmt"
	"log"
)

// #include "hello.c"
import "C"

func main() {
	fmt.Println("start!")

	_, err := C.Hello()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("-end-")
}
