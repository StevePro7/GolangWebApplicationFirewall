package waf

import (
	"reflect"
	"testing"
)

const directory = "/etc/waf"

func TestExtractRulesSetFilenames(t *testing.T) {
	InitializeModSecurity(directory)

	expectFilenames := []string{"/etc/waf/REQUEST-942-APPLICATION-ATTACK-SQLI.conf", "/etc/waf/crs-setup.conf", "/etc/waf/modsecdefault.conf"}
	actualFilenames := ExtractRulesSetFilenames()

	test := reflect.DeepEqual(expectFilenames, actualFilenames)
	if !test {
		t.Errorf("Expect '%s' Actual '%s'", expectFilenames, actualFilenames)
	}
}

//func TestInitializeModSecurity(t *testing.T) {
//	exp := "/etc/waf/"
//	InitializeModSecurity(directory)
//	act := GetRulesDirectory()
//	if exp != act {
//		t.Errorf("exp '%s' act '%s'", exp, act)
//	}
//}

//func
//TestLoadModSecurityCoreRuleSet2(t * testing.T)
//{
//	exp := 5
//	act := Add(1, 2)
//	if exp != act {
//		t.Errorf("exp %d act %d", exp, act)
//	}
//}
