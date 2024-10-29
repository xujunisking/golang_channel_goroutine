package main

import (
	"fmt"
	"reflect"
)

type Cal struct {
	Num1 int
	Num2 int
}

func (c Cal) GetSub(name string) {
	res := c.Num1 - c.Num2

	fmt.Printf("%v 完成了减法运算，%d - %d = %d", name, c.Num1, c.Num2, res)
}

func main() {
	model := &Cal{}

	sv := reflect.ValueOf(model).Elem()

	kd := sv.Kind()
	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}

	num := sv.NumField()
	fmt.Printf("struct has %d fields\n", num)
	fmt.Printf("rVal=%v rVal type=%T\n", sv, sv)

	// resFound := sv.FieldByName("num1")
	// if !resFound.IsValid() {
	// 	fmt.Printf("没有找到num1字段！")
	// 	return
	// }
	//fmt.Printf("value=%v rVal type=%T\n", sv.FieldByName("num1"), sv.FieldByName("num1"))

	//field1Value, found := sv.FieldByName("Num1")
	found := sv.FieldByName("Num1")
	if found.IsValid() {
		sv.FieldByName("Num1").SetInt(8)
	}
	found = sv.FieldByName("Num2")
	if found.IsValid() {
		sv.FieldByName("Num2").SetInt(3)
	}

	var params []reflect.Value
	params = append(params, reflect.ValueOf("tom"))
	sv.Method(0).Call(params)
}
