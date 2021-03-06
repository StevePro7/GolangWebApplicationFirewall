package main

import (
	"waftesting/waf"
)

func runServer() {
	rulesetDirectory := "/etc/waf"

	// Initialize WAF and load OWASP Core Rule Sets.
	waf.InitializeModSecurity()
	waf.DefineRulesSetDirectory(rulesetDirectory)
	filenames := waf.ExtractRulesSetFilenames()
	waf.LoadModSecurityCoreRuleSet(filenames)
}

func main() {
	runServer()
	checker.Check()
}
