package checker

import (
	"log"
	"waftesting/waf"
)

func Check() {

	id := waf.GenerateModSecurityID()
	//url := "/"
	//url := "/foo.com"
	url := "/test/artists.php?artist=0+div+1+union%23foo*%2F*bar%0D%0Aselect%23foo%0D%0A1%2C2%2Ccurrent_user"
	httpMethod := "GET"
	httpProtocol := "HTTP"
	httpVersion := "1.1"
	clientHost := "192.168.68.202"
	var clientPort uint32
	clientPort = 60972
	serverHost := "192.168.101.138"
	var serverPort uint32
	serverPort = 80

	detection := waf.ProcessHttpRequest(id, url, httpMethod, httpProtocol, httpVersion, clientHost, clientPort, serverHost, serverPort)

	log.Println("Detection = ", detection)
}
