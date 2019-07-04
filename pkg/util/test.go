package util

import (
	"errors"
	"fmt"
	"strings"
)

func ChkEqual(expect, actual interface{}) error {
	strA, _ := expect.(string)
	strB, _ := actual.(string)
	if strings.Compare(strA, strB) != 0 {
		return errors.New(fmt.Sprintf("Expect string: \"%s\", got: \"%s\"", expect, actual))
	}
	return nil
}
