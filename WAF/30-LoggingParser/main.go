package main

import (
	"fmt"
	"waftesting/waf"
)

func main() {
	fmt.Println("hello")

	x := 1
	y := 2
	res := waf.GoAdd(x, y)
	fmt.Println(res)
}
