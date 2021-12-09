package waf

import (
	"reflect"
	"testing"
)

const directory = "/etc/waf"

// 03.
func TestLoadModSecurityCoreRuleSet(t *testing.T) {

	//expectFilenames := []string{"/etc/waf/REQUEST-942-APPLICATION-ATTACK-SQLI.conf", "/etc/waf/crs-setup.conf", "/etc/waf/modsecdefault.conf"}
	filenames := []string{"/etc/waf/crs-setup.conf"}

	expect := -2
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
