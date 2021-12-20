package waf

import (
	"testing"
)

const testCoreRulesetDirectory = "rules/core-rules"
const testCustomRulesetDirectory = "rules/custom-rules"

func TestInitializeModSecurity(t *testing.T) {

	InitializeModSecurity()
}

func TestDefineRulesSetDirectory(t *testing.T) {

	InitializeModSecurity()
	DefineRulesSetDirectory(testCoreRulesetDirectory)
	expect := "rules/core-rules/"
	actual := GetRulesDirectory()
	if expect != actual {
		t.Errorf("Expect: '%s' Actual: '%s'", expect, actual)
	}
}

func TestExtractRulesSetFilenames(t *testing.T) {

	InitializeModSecurity()
	DefineRulesSetDirectory(testCoreRulesetDirectory)

	expectFilenames := []string{
		"rules/core-rules/crs-setup.conf",
		"rules/core-rules/modsecdefault.conf",
		"rules/core-rules/REQUEST-942-APPLICATION-ATTACK-SQLI.conf",
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
		"rules/core-rules/crs-setup.conf",
		"rules/core-rules/modsecdefault.conf",
		"rules/core-rules/REQUEST-942-APPLICATION-ATTACK-SQLI.conf",
	}

	expect := len(filenames)
	actual := LoadModSecurityCoreRuleSet(filenames)

	if expect != actual {
		t.Errorf("Expect: %d Actual: %d", expect, actual)
	}
}

func TestGenerateModSecurtityID(t *testing.T) {

	id := GenerateModSecurtityID()
	expectLength := 36
	actualLength := len(id)

	if expectLength != actualLength {
		t.Errorf("ID '%s' expect length: %d actual length: %d", id, expectLength, actualLength)
	}
}

func TestProcessHttpRequest_ValidURL_OK(t *testing.T) {

	InitializeModSecurity()
	DefineRulesSetDirectory(testCoreRulesetDirectory)
	filenames := ExtractRulesSetFilenames()
	LoadModSecurityCoreRuleSet(filenames)

	id := "7ce62288-d6dd-4be0-8b31-ae27876aeeea"
	url := "/foo.com"
	httpMethod := "GET"
	httpProtocol := "HTTP"
	httpVersion := "1.1"
	clientLink := "http://localhost"
	clientPort := 80
	serverLink := "http://localhost"
	serverPort := 80

	expect := 0
	actual := ProcessHttpRequest(id, url, httpMethod, httpProtocol, httpVersion, clientLink, clientPort, serverLink, serverPort)

	if expect != actual {
		t.Errorf("Expect: %d Actual: %d", expect, actual)
	}
}

func TestProcessHttpRequest_InvalidURL_BlockDueToWarning(t *testing.T) {

	InitializeModSecurity()
	DefineRulesSetDirectory(testCoreRulesetDirectory)
	filenames := ExtractRulesSetFilenames()
	LoadModSecurityCoreRuleSet(filenames)

	id := "7ce62288-d6dd-4be0-8b31-ae27876aeeea"
	url := "/test/artists.php?artist=0+div+1+union%23foo*%2F*bar%0D%0Aselect%23foo%0D%0A1%2C2%2Ccurrent_user"
	httpMethod := "GET"
	httpProtocol := "HTTP"
	httpVersion := "1.1"
	clientLink := "http://localhost"
	clientPort := 80
	serverLink := "http://localhost"
	serverPort := 80

	expect := 0
	actual := ProcessHttpRequest(id, url, httpMethod, httpProtocol, httpVersion, clientLink, clientPort, serverLink, serverPort)

	if expect != actual {
		t.Errorf("Expect: %d Actual: %d", expect, actual)
	}
}

func TestProcessHttpRequest_InvalidURL_NoRulesLoad_OK(t *testing.T) {

	InitializeModSecurity()
	var filenames []string
	LoadModSecurityCoreRuleSet(filenames)

	id := "7ce62288-d6dd-4be0-8b31-ae27876aeeea"
	url := "/test/artists.php?artist=0+div+1+union%23foo*%2F*bar%0D%0Aselect%23foo%0D%0A1%2C2%2Ccurrent_user"
	httpMethod := "GET"
	httpProtocol := "HTTP"
	httpVersion := "1.1"
	clientLink := "http://localhost"
	clientPort := 80
	serverLink := "http://localhost"
	serverPort := 80

	expect := 0
	actual := ProcessHttpRequest(id, url, httpMethod, httpProtocol, httpVersion, clientLink, clientPort, serverLink, serverPort)

	if expect != actual {
		t.Errorf("Expect: %d Actual: %d", expect, actual)
	}
}

func TestProcessHttpRequest_InvalidURL_CustomRulesLoad_BadRequest(t *testing.T) {

	InitializeModSecurity()
	DefineRulesSetDirectory(testCustomRulesetDirectory)
	filenames := ExtractRulesSetFilenames()
	LoadModSecurityCoreRuleSet(filenames)

	id := "7ce62288-d6dd-4be0-8b31-ae27876aeeea"
	url := "/test/artists.php?artist=0+div+1+union%23foo*%2F*bar%0D%0Aselect%23foo%0D%0A1%2C2%2Ccurrent_user"
	httpMethod := "GET"
	httpProtocol := "HTTP"
	httpVersion := "1.1"
	clientLink := "http://localhost"
	clientPort := 80
	serverLink := "http://localhost"
	serverPort := 80

	expect := 1
	actual := ProcessHttpRequest(id, url, httpMethod, httpProtocol, httpVersion, clientLink, clientPort, serverLink, serverPort)

	if expect != actual {
		t.Errorf("Expect: %d Actual: %d", expect, actual)
	}
}
