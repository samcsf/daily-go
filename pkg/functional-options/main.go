package main

import "fmt"

type ExampleFetcher struct {
	timeout int
	retry   int
}

func NewFetcher(options ...func(*ExampleFetcher)) *ExampleFetcher {
	ef := &ExampleFetcher{}
	for _, op := range options {
		op(ef)
	}
	return ef
}

func main() {
	timeout := func(e *ExampleFetcher) {
		e.timeout = 8080
	}
	retry := func(e *ExampleFetcher) {
		e.retry = 3
	}
	ef := NewFetcher(timeout, retry)
	fmt.Println(ef)
}
