package main

import (
	"fmt"

	"github.com/projectcalico/app-policy/waf"
)

func main() {
	x := 3
	y := 4

	a := waf.MyGoAdd(x, y)
	fmt.Println("MyGoAdd4 = ", a)

	b := waf.MyCGoAdd(x, y)
	fmt.Println("MyCGoAdd4 = ", b)
}
