package waf

import "testing"

func TestParser01(t *testing.T) {

	expect := 4
	actual := GoAdd(1, 2)
	if expect != actual {
		t.Errorf("Expect: %d Actual: %d", expect, actual)
	}
}
