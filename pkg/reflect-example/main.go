package main

import (
	"fmt"
	"reflect"
)

/*
	go 中每个interface类型都会保存它对应的type, value
	有了这个前提reflect中的各种api的意义将比较明确用途
*/

type animal struct{}

func (a *animal) SayHi() {
	fmt.Println("hi")
}

func (a *animal) SayYo() {
	fmt.Println("yo")
}

func main() {
	i := 10
	// typeOf / valueOf 查看当前type和value
	fmt.Println("Type is :", reflect.TypeOf(i))
	fmt.Println("Value is :", reflect.ValueOf(i))

	// 动态调用方法
	ani := &animal{}
	val := reflect.ValueOf(ani)
	val.MethodByName("SayHi").Call(nil)
	val.MethodByName("SayYo").Call(nil)
}
