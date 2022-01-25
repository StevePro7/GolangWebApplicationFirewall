package main

import (
	"fmt"
	"waftesting/waf"
)

func main() {
	fmt.Println("begin")
	waf.InitializeModSecurity()
	fmt.Println("-end-")
}
