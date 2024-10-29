package main

import (
	"fmt"
	"reflect"
)

func reflect01(b interface{}) {

	//获取到 reflect.Value
	rVal := reflect.ValueOf(b)
	//查看rVal的Kind是什么类型
	fmt.Printf("rVal Kind=%v\n", rVal.Kind())
	//获取到指针指向的值
	rVal.Elem().SetInt(20)

}

func main() {
	var num int = 10
	reflect01(&num)
	fmt.Println("num=", num)
}
