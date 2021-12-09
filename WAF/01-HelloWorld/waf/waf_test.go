package waf

import (
	"reflect"
	"testing"
)

const directory = "/etc/waf"

// 04.
func TestProcessHttpRequest_InvalidURL_CustomRulesLoad_BadRequest(t *testing.T) {

	filenames := []string{"/etc/waf/stevepro-REQUEST-942-APPLICATION-ATTACK-SQLI.conf", "/etc/waf/crs-setup.conf", "/etc/waf/modsecdefault.conf"}
	LoadModSecurityCoreRuleSet(filenames)

	url := "/test/artists.php?artist=0+div+1+union%23foo*%2F*bar%0D%0Aselect%23foo%0D%0A1%2C2%2Ccurrent_user"
	httpMethod := "GET"
	httpProtocol := "HTTP"
	httpVersion := "1.1"
	clientLink := "http://localhost"
	clientPort := 80
	serverLink := "http://localhost"
	serverPort := 80

	expect := 1
	actual := ProcessHttpRequest(url, httpMethod, httpProtocol, httpVersion, clientLink, clientPort, serverLink, serverPort)

	if expect != actual {
		t.Errorf("Expect: %d Actual: %d", expect, actual)
	}
}

func TestProcessHttpRequest_InvalidURL_NoRulesLoad_OK(t *testing.T) {

	filenames := []string{}
	LoadModSecurityCoreRuleSet(filenames)

	url := "/test/artists.php?artist=0+div+1+union%23foo*%2F*bar%0D%0Aselect%23foo%0D%0A1%2C2%2Ccurrent_user"
	httpMethod := "GET"
	httpProtocol := "HTTP"
	httpVersion := "1.1"
	clientLink := "http://localhost"
	clientPort := 80
	serverLink := "http://localhost"
	serverPort := 80

	expect := 0
	actual := ProcessHttpRequest(url, httpMethod, httpProtocol, httpVersion, clientLink, clientPort, serverLink, serverPort)

	if expect != actual {
		t.Errorf("Expect: %d Actual: %d", expect, actual)
	}
}

func TestProcessHttpRequest_InvalidURL_BlockDueToWarning(t *testing.T) {

	filenames := []string{"/etc/waf/REQUEST-942-APPLICATION-ATTACK-SQLI.conf", "/etc/waf/crs-setup.conf", "/etc/waf/modsecdefault.conf"}
	LoadModSecurityCoreRuleSet(filenames)

	url := "/test/artists.php?artist=0+div+1+union%23foo*%2F*bar%0D%0Aselect%23foo%0D%0A1%2C2%2Ccurrent_user"
	httpMethod := "GET"
	httpProtocol := "HTTP"
	httpVersion := "1.1"
	clientLink := "http://localhost"
	clientPort := 80
	serverLink := "http://localhost"
	serverPort := 80

	expect := 0
	actual := ProcessHttpRequest(url, httpMethod, httpProtocol, httpVersion, clientLink, clientPort, serverLink, serverPort)

	if expect != actual {
		t.Errorf("Expect: %d Actual: %d", expect, actual)
	}
}

func TestProcessHttpRequest_ValidURL_OK(t *testing.T) {

	filenames := []string{"/etc/waf/REQUEST-942-APPLICATION-ATTACK-SQLI.conf", "/etc/waf/crs-setup.conf", "/etc/waf/modsecdefault.conf"}
	LoadModSecurityCoreRuleSet(filenames)

	url := "/foo.com"
	httpMethod := "GET"
	httpProtocol := "HTTP"
	httpVersion := "1.1"
	clientLink := "http://localhost"
	clientPort := 80
	serverLink := "http://localhost"
	serverPort := 80

	expect := 0
	actual := ProcessHttpRequest(url, httpMethod, httpProtocol, httpVersion, clientLink, clientPort, serverLink, serverPort)

	if expect != actual {
		t.Errorf("Expect: %d Actual: %d", expect, actual)
	}
}

// 03.
func TestLoadModSecurityCoreRuleSet(t *testing.T) {

	filenames := []string{"/etc/waf/REQUEST-942-APPLICATION-ATTACK-SQLI.conf", "/etc/waf/crs-setup.conf", "/etc/waf/modsecdefault.conf"}

	expect := len(filenames)
	actual := LoadModSecurityCoreRuleSet(filenames)

	if expect != actual {
		t.Errorf("Expect: %d Actual: %d", expect, actual)
	}
}

// 02.
func TestExtractRulesSetFilenames(t *testing.T) {
	InitializeModSecurity(directory)

	expectFilenames := []string{"/etc/waf/REQUEST-942-APPLICATION-ATTACK-SQLI.conf", "/etc/waf/crs-setup.conf", "/etc/waf/modsecdefault.conf"}
	actualFilenames := ExtractRulesSetFilenames()

	test := reflect.DeepEqual(expectFilenames, actualFilenames)
	if !test {
		t.Errorf("Expect '%s' Actual '%s'", expectFilenames, actualFilenames)
	}
}

// 01.
func TestInitializeModSecurity(t *testing.T) {

	InitializeModSecurity(directory)
	expect := "/etc/waf/"
	actual := GetRulesDirectory()
	if expect != actual {
		t.Errorf("Expect: '%s' Actual: '%s'", expect, actual)
	}
}
