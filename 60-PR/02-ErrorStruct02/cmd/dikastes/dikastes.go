package main

import (
	"fmt"
	"waftesting/waf"
)

func main() {
	fmt.Println("begin")

	rulesetDirectory := "/etc/waf/"

	waf.InitializeModSecurity()
	waf.DefineRulesSetDirectory(rulesetDirectory)
	filenames := []string{
		"etc/waf/myinvalidfile.conf",
	}
	waf.LoadModSecurityCoreRuleSet(filenames)

	fmt.Println("-end!!")
}
