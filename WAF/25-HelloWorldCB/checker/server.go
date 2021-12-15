package checker

import (
	"log"
	"waftesting/waf"
)

func Check() {
	//log.Println("check")

	waf.InitializeModSecurity()
	waf.DefineRulesSetDirectory("/etc/waf")

	filenames := waf.ExtractRulesSetFilenames()
	waf.LoadModSecurityCoreRuleSet(filenames)

	//url := "/"
	//url := "/foo.com"
	url := "/http://localhost:3080/test/artists.php?artist=0+div+1+union%23foo*%2F*bar%0D%0Aselect%23foo%0D%0A1%2C2%2Ccurrent_user"
	httpMethod := "GET"
	httpProtocol := "HTTP"
	httpVersion := "1.1"
	clientLink := "echo-a"
	clientPort := 80
	serverLink := "echo-a"
	serverPort := 80
	detection := waf.ProcessHttpRequest(url, httpMethod, httpProtocol, httpVersion, clientLink, clientPort, serverLink, serverPort)

	log.Println("Detection = ", detection)
	/**/
}
