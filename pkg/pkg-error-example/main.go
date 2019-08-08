package main

import (
	"fmt"
	"github.com/pkg/errors"
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func main() {
	// Wrap 方法将每一层的信息包裹起来，包括stack
	e := a()
	if e != nil {
		fmt.Println(e)
	}

	// 因为返回的是error interface所以对自定义error进行方法调用之前要使用type assert
	// StackTrace + "%+v" 能输出具体到文件行数的信息
	err, ok := errors.Cause(e).(stackTracer)
	if !ok {
		panic("err does not implement stackTrace()")
	}
	st := err.StackTrace()
	fmt.Printf("%+v", st)
}

func a() error {
	e1 := b()
	if e1 != nil {
		return errors.Wrap(e1, "error in a() calling b()")
	}
	e2 := c()
	if e2 != nil {
		return errors.Wrap(e2, "error in a() calling c()")
	}
	return nil
}

func b() error {
	return nil
}

func c() error {
	e := d()
	return errors.Wrap(e, "error in c()")
}

func d() error {
	return errors.Wrap(errors.New("my error"), "error in d()")
}
