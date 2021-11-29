package main

import (
	"fmt"

	"github.com/projectcalico/app-policy/waf"
)

func main() {
	a := waf.MyGoAdd(6, 3)
	fmt.Println("MyGoAdd4 = ", a)

	//b := waf.MyCGoAdd(6, 3)
	//fmt.Println("MyCGoAdd4 = ", b)
}
