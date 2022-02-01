package main

import (
	"fmt"
	"log"
	"waftesting/waf"
)

func main() {
	var err error
	waf.InitializeModSecurity()
	_ = waf.LoadModSecurityCoreRuleSet("waf/test_files/core-rules/crs-setup.conf")
	_ = waf.LoadModSecurityCoreRuleSet("waf/test_files/core-rules/modsecdefault.conf")
	_ = waf.LoadModSecurityCoreRuleSet("waf/test_files/core-rules/REQUEST-942-APPLICATION-ATTACK-SQLI.conf")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("hello!!")
	waf.CleanupModSecurity()
}
