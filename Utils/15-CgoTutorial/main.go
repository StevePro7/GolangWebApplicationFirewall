package main

//#include "hello.h"
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

	// Call to int function with two params
	a := 1
	b := 4
	res, err := makeSum(a, b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Sum of %d + %d = %d\n", a, b, res)

	fmt.Println("end")
}

func Hello() error {
	_, err := C.Hello()
	if err != nil {
		return errors.New("error calling Hello function: " + err.Error())
	}
	return nil
}

// makeSum also is a C binding to make a sum
// it returns a results and an error
func makeSum(a, b int) (int, error) {
	// Convert Go ints to C ints
	aC := C.int(a)
	bC := C.int(b)

	sum, err := C.sum(aC, bC)
	if err != nil {
		return 0, errors.New("error calling Sum function: " + err.Error())
	}

	// convert C.int result to Go int
	res := int(sum)
	return res, nil
}
