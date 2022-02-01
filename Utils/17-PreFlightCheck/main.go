package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"waftesting/waf"
)

func main() {
	rulesetDirectory := ""
	fmt.Println(rulesetDirectory)

	err := waf.CheckRulesSetExists(rulesetDirectory)
	if err != nil {
		log.Errorf("WAF Core Rules Set check: '%s'", err.Error())
	}

	if waf.IsEnabled() {
		log.Info("WAF is enabled...")
	} else {
		log.Info("WAF is NOT enabled!")
	}
}
