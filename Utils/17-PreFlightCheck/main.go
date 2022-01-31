package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"waftesting/waf"
)

func main() {
	rulesetDirectory := "/etc/waf2"
	fmt.Println(rulesetDirectory)

	err := waf.CheckRulesSetExists(rulesetDirectory)
	if err != nil {
		log.Errorf("WAF Core Rules Set check: '%s'", err.Error())
	}

	if waf.IsEnabled() {
		log.Info("WAF Web Application Firewal is enabled.")
	} else {
		log.Info("WAF Web Application Firewal is NOT enabled!")
	}
	fmt.Println("hello")
}
