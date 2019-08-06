package errors

import (
// "github.com/sirupsen/logrus"
)

const KindUnexpected = "KindUnexpected"

type Op string // unique mark

type Kind string

type Error struct {
	Op   Op
	Kind Kind
	Err  error
	// Severity logrus.Level
}

func E(args ...interface{}) error {
	e := &Error{}
	for _, arg := range args {
		switch arg := arg.(type) {
		case Op:
			e.Op = arg
		case Kind:
			e.Kind = arg
		case error:
			e.Err = arg
		default:
			panic("bad call to E")
		}
	}
	return e
}

func Ops(e *Error) []Op {
	res := []Op{e.Op}

	subErr, ok := err.Err.(*Error)
	if !ok {
		return res
	}

	res = append(res, Ops(subErr)...)
	return res
}

func Kind(err error) string {
	e, ok := err.(*Error)
	if !ok {
		return KindUnexpected
	}

	if e.Kind != "" {
		return e.Kind
	}

	return Kind(e.Err)
}
