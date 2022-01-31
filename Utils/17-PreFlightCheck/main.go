package main

import (
	"fmt"
	"waftesting/waf"
)

func main() {
	rulesetDirectory := "/etc/waf/"
	fmt.Println(rulesetDirectory)
	waf.PreFlightCheck(rulesetDirectory)

	if waf.IsEnabled() {
		fmt.Println("WAF baby!")
	}
	fmt.Println("hello")
}
