package main

import (
	"fmt"
	"log"
	"waftesting/waf"
)

func MyTest01() {
	log.Println("MyTest01")
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

	waf.CleanupModSecurity()
}

func MyTest02() {
	log.Println("MyTest02")
	var err error
	waf.InitializeModSecurity()

	_ = waf.LoadModSecurityCoreRuleSet("waf/test_files/core-rules/crs-setup.conf")
	_ = waf.LoadModSecurityCoreRuleSet("waf/test_files/core-rules/modsecdefault.conf")
	_ = waf.LoadModSecurityCoreRuleSet("waf/test_files/core-rules/REQUEST-942-APPLICATION-ATTACK-SQLI.conf")

	id := "7ce62288-d6dd-4be0-8b31-ae27876aeeea"
	url := "/test/artists.php?artist=0+div+1+union%23foo*%2F*bar%0D%0Aselect%23foo%0D%0A1%2C2%2Ccurrent_user"
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

	waf.CleanupModSecurity()
}

func main() {
	fmt.Println("hello!!")
	MyTest02()
	fmt.Println("bye")
}
