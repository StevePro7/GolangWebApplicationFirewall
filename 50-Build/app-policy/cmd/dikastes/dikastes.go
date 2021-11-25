package main

import (
	"fmt"
	"githu.com/stevepro/steven/waf"
)

func main() {
	z := waf.MyAdd(4, 3)
	fmt.Println("MyAdd = ", z)
}
