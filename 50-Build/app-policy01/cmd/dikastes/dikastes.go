package main

import (
	"fmt"
	"github.com/stevepro/app-policy/waf"
)

func main() {
	z := waf.MyAdd(5, 3)
	fmt.Println("MyAdd = ", z)
}
