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

	id := "7ce62288-d6dd-4be0-8b31-ae27876aeeea"
	url := "/foo.com"
	httpMethod := "GET"
	httpProtocol := "HTTP"
	httpVersion := "1.1"
	clientHost := "http://localhost"
	clientPort := uint32(80)
	serverHost := "http://localhost"
	serverPort := uint32(80)

	err = waf.ProcessHttpRequest(id, url, httpMethod, httpProtocol, httpVersion, clientHost, clientPort, serverHost, serverPort)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("hello!!")
	waf.CleanupModSecurity()
	fmt.Println("bye")
}
