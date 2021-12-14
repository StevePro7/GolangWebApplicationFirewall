package main

import (
	"fmt"
	"waftesting/waf"
)

func main() {
	fmt.Println("beg")
	callback.Run()
	fmt.Println("end")
}
