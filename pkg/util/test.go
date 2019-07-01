package test

import (
	"log"
	"strings"
	"testing"
)

func ChkEqual(a, b interface{}, t *testing.T) {
	strA, err := a.(string)
	ChkErr(err)
	strB, err := b.(string)
	ChkErr(err)
	if strings.Compare(a, b) != 0 {
		log.Println("Compare failed")
		t.Fail()
	}

}
