package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"waftesting/waf"
)

func main() {
	rulesetDirectory := "/etc/waf2"
	fmt.Println(rulesetDirectory)

	err := waf.Check(rulesetDirectory)
	if err != nil {
		log.Errorf(err.Error())
	}
	//waf.DefineRulesSetDirectory(rulesetDirectory)
	//err := waf.CheckRulesSetDirectoryExists()
	//if err != nil {
	//	log.Printf("WAF Core Rules Set directory: '%s'  does not exist.  WAF not enabled!", rulesetDirectory)
	//}
	//
	//filenames, err := waf.ExtractRulesSetFilenames()
	//if err != nil {
	//	log.Fatalf("WAF Core Rules Set directory: '%s' does not exist!", rulesetDirectory)
	//}

	if waf.IsEnabled() {
		fmt.Println("WAF baby!")
	}
	fmt.Println("hello")
}
