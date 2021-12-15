package waf

import (
	"testing"
)

const directory = "/etc/waf"

func TestGenerateModSecurtityID(t *testing.T) {

	id := GenerateModSecurtityID()
	expectLength := 36
	actualLength := len(id)

	if expectLength != actualLength {
		t.Errorf("ID '%s' expect length: %d actual length: %d", id, expectLength, actualLength)
	}
}

func TestInitializeModSecurity(t *testing.T) {

	InitializeModSecurity()
}

func TestDefineRulesSetDirectory(t *testing.T) {

	InitializeModSecurity()
	DefineRulesSetDirectory(directory)
	expect := "/etc/waf/"
	actual := GetRulesDirectory()
	if expect != actual {
		t.Errorf("Expect: '%s' Actual: '%s'", expect, actual)
	}
}

func TestExtractRulesSetFilenames(t *testing.T) {

	InitializeModSecurity()
	DefineRulesSetDirectory(directory)

	expectFilenames := []string{
		"/etc/waf/crs-setup.conf",
		"/etc/waf/modsecdefault.conf",
		"/etc/waf/REQUEST-942-APPLICATION-ATTACK-SQLI.conf",
	}
	actualFilenames := ExtractRulesSetFilenames()

	test := len(expectFilenames) == len(actualFilenames)
	if !test {
		t.Errorf("Expect '%s' Actual '%s'", expectFilenames, actualFilenames)
	}
}

func TestLoadModSecurityCoreRuleSet(t *testing.T) {

	InitializeModSecurity()
	filenames := []string{
		"/etc/waf/crs-setup.conf",
		"/etc/waf/modsecdefault.conf",
		"/etc/waf/REQUEST-942-APPLICATION-ATTACK-SQLI.conf",
	}

	expect := len(filenames)
	actual := LoadModSecurityCoreRuleSet(filenames)

	if expect != actual {
		t.Errorf("Expect: %d Actual: %d", expect, actual)
	}
}

func TestProcessHttpRequest_ValidURL_OK(t *testing.T) {

	InitializeModSecurity()
	filenames := []string{
		"/etc/waf/crs-setup.conf",
		"/etc/waf/modsecdefault.conf",
		"/etc/waf/REQUEST-942-APPLICATION-ATTACK-SQLI.conf",
	}
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

func TestProcessHttpRequest_InvalidURL_BlockDueToWarning(t *testing.T) {

	InitializeModSecurity()
	filenames := []string{
		"/etc/waf/crs-setup.conf",
		"/etc/waf/modsecdefault.conf",
		"/etc/waf/REQUEST-942-APPLICATION-ATTACK-SQLI.conf",
	}
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

func TestProcessHttpRequest_InvalidURL_NoRulesLoad_OK(t *testing.T) {

	InitializeModSecurity()
	var filenames []string
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

func TestProcessHttpRequest_InvalidURL_CustomRulesLoad_BadRequest(t *testing.T) {

	InitializeModSecurity()
	DefineRulesSetDirectory("/etc/waf_custom")
	filenames := ExtractRulesSetFilenames()
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
