package main

import (
	"fmt"
	"waftesting/callback"
)

func main() {
	fmt.Println("beg")
	callback.Run()
	fmt.Println("end")
}
