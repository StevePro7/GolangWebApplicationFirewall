package main

import "fmt"

// #include "hello.c"
import "C"

func main() {
	fmt.Println("start")

	fmt.Println("-end-")
}
