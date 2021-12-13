package main

import (
	"fmt"
	"waftesting/callback"
)

func main() {
	fmt.Println("hello")
	callback.Run()
	sum := callback.GoAdd(4, 3)
	fmt.Println(sum)
	fmt.Println("-bye-")
}
