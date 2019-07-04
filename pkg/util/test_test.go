package util

import "testing"

func TestChkEqual(t *testing.T) {
	err := ChkEqual("a", "b")
	if err == nil {
		t.Error("Expect error when string not match")
	}
	err = ChkEqual("aaa", "aaa")
	if err != nil {
		t.Error(err)
	}
}
