package main

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name  string `json:"monster_name"`
	Age   int    `json:"monster_age"`
	Score float32
	Sex   string
}

func (s Monster) Print() {
	fmt.Println("----start----")
	fmt.Println(s)
	fmt.Println("----end----")
}

func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

func (s Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

func TestStruct(a interface{}) {
	//获取reflect.Type 类型
	typ := reflect.TypeOf(a)
	//获取reflect.Value 类型
	val := reflect.ValueOf(a)
	//获取Kind 类别
	kd := val.Kind()
	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}

	//获取到该结构体有几个字段
	num := val.NumField()
	fmt.Printf("struct has %d fields\n", num)
	//遍历结构体的所有字段
	for i := 0; i < num; i++ {
		fmt.Printf("Field %d:值为=%v\n", i, val.Field(i))
		//获取到struct标签，注意需要通过reflect.Type来获取tag标签的值
		tagVal := typ.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("Field %d: tag为=%v\n", i, tagVal)
		}
	}
	//获取到该结构体有几个方法
	numofMethod := val.NumMethod()
	fmt.Printf("struct has %d methods\n", numofMethod)

	//var params []reflect.Value
	//val.Method获取到的方法是按照方法的首字母排序的
	val.Method(1).Call(nil)
	//调用结构体的第1个方法Method(e)
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	//传入的参数是[lreflect.Value
	res := val.Method(0).Call(params)
	fmt.Println("res=", res[0].Int())
	//返回结果，返回的结果是[]reflect.Value*
}

func main() {

	var a Monster = Monster{
		Name:  "黄鼠狼",
		Age:   400,
		Score: 30.8,
	}

	TestStruct(a)
}
