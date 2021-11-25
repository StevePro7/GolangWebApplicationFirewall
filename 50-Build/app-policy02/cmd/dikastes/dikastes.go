package main

import (
	"fmt"
	"github.com/stevepro/app-policy/waf"
)

func main() {
	a := waf.MyGoAdd(5, 3)
	fmt.Println("MyGoAdd = ", a)

	b := waf.MyCGoAdd(5, 3)
	fmt.Println("MyCGoAdd = ", b)
}
