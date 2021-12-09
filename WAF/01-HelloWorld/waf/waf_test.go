package waf

import "testing"

const directory = "/etc/waf"

//func TestExtractRulesSetFilenames(t *testing.T) {
//	InitializeModSecurity(directory)
//
//	expect_filenames := []string{"crs-setup.conf"}
//	actual_filenames := ExtractRulesSetFilenames()
//		//act := GetRulesDirectory()
//		if exp != filenames{
//		t.Errorf("exp '%s' act '%s'", exp, filenames)
//	}

func TestInitializeModSecurity(t *testing.T) {
	exp := "/etc/waf/"
	InitializeModSecurity(directory)
	act := GetRulesDirectory()
	if exp != act {
		t.Errorf("exp '%s' act '%s'", exp, act)
	}
}

//func
//TestLoadModSecurityCoreRuleSet2(t * testing.T)
//{
//	exp := 5
//	act := Add(1, 2)
//	if exp != act {
//		t.Errorf("exp %d act %d", exp, act)
//	}
//}
