package main

import (
	"fmt"
	"waftesting/waf"
)

func main() {
	fmt.Println("begin")
	waf.InitializeModSecurity()
	rulesetDirectory := "/etc/waf/"
	waf.DefineRulesSetDirectory(rulesetDirectory)
	//filenames := waf.ExtractRulesSetFilenames()

	filenames := "REQUEST-942-APPLICATION-ATTACK-SQLI.conf"
	//filenames := "/etc/waf/REQUEST-942-APPLICATION-ATTACK-SQLI.conf"

	waf.LoadModSecurityCoreRuleSet(filenames)
	fmt.Println("-end-")
}
