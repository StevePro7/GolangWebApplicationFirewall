package main

//#include "hello.c"
import "C"
import (
	"errors"
	"fmt"
	"log"
)

func main() {
	fmt.Println("beg")
	// Call to void function without params
	err := Hello()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("end")
}

func Hello() error {
	_, err := C.Hello()
	if err != nil {
		return errors.New("error calling Hello function: " + err.Error())
	}

	return nil
}
