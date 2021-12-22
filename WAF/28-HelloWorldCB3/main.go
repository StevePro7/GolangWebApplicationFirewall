package main

import (
	"fmt"
	"waftesting/checker"
	"waftesting/waf"
)

func runServer() {
	//rulesetDirectory := "waf/rules/core-rules"
	rulesetDirectory := "waf/rules/custom-rules"

	// Initialize WAF and load OWASP Core Rule Sets.
	waf.InitializeModSecurity()
	waf.DefineRulesSetDirectory(rulesetDirectory)
	filenames := waf.ExtractRulesSetFilenames()
	waf.LoadModSecurityCoreRuleSet(filenames)
}

func main() {
	detection := 0
	var outcome string
	if detection == 0 {
		outcome = "OK"
	} else {
		outcome = "FAILED"
	}

	fmt.Println(detection)
	fmt.Println(outcome)

	runServer()
	checker.Check()
}
