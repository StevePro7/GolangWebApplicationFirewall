package waf

import "testing"

func TestLoadModSecurityCoreRuleSet(t *testing.T) {
	exp := "test"
	InitializeModSecurity("/tmp")
	act := GetBob()
	if exp != act {
		t.Errorf("exp '%s' act '%s'", exp, act)
	}
}

func TestLoadModSecurityCoreRuleSet2(t *testing.T) {
	exp := 5
	act := Add(1, 2)
	if exp != act {
		t.Errorf("exp %d act %d", exp, act)
	}
}
