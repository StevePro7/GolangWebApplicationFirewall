package waf

import (
	"testing"
)

const testErrorRulesetDirectory = "test_files/error-rules"

func TestLoadModSecurityCoreRuleSetErrorDirectory2(t *testing.T) {

	InitializeModSecurity()
	DefineRulesSetDirectory(testErrorRulesetDirectory)

	filenames := []string{
		"test_files/error-rules/REQUEST-941-APPLICATION-ATTACK-XSS.conf",
		"test_files/error-rules/REQUEST-942-APPLICATION-ATTACK-SQLI.conf",
	}

	expect := 1
	actual := LoadModSecurityCoreRuleSet2(filenames)

	if expect != actual {
		t.Errorf("Expect: %d Actual: %d", expect, actual)
	}
}
