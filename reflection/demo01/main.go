package main

import (
	"fmt"
	"reflect"
)

func reflectTest01(b interface{}) {

	//通过反射获取的传入的变量的type，kind，值
	//1.先获取到 reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType=", rTyp)

	//2.获取到 reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal=%v rVal type=%T\n", rVal, rVal)

	n2 := 2 + rVal.Int()
	fmt.Println("n2=", n2)

	//下面我们将 ryal 转成 interface{}
	iV := rVal.Interface()
	//将 interface{} 通过断言转成需要的类型
	num2 := iV.(int)
	fmt.Println("num2=", num2)
}

func reflectTest02(b interface{}) {
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType=", rTyp)

	rVal := reflect.ValueOf(b)
	//3.获取 变量对应的Kind
	//(1)rval.Kind()==>
	kind1 := rVal.Kind()
	//(2)rTyp.Kind()==>
	kind2 := rTyp.Kind()
	fmt.Printf("kind =%v kind=%v", kind1, kind2)

	iV := rVal.Interface()
	fmt.Printf("iV=%v iV type=%T\n", iV, iV)
	//将 interface{} 通过断言转成需要的类型
	stu, ok := iV.(Student)
	if ok {
		fmt.Printf("stu.Name=%v stu.Age=%v\n", stu.Name, stu.Age)
	}
}

type Student struct {
	Name string
	Age  int
}

func main() {

	//请编写一个案例，
	//演示对(基本数据类型、interface{}、reflect.Value)进行反射的基本操作

	var num int = 100
	reflectTest01(num)

	stu := Student{
		Name: "tom",
		Age:  20,
	}
	reflectTest02(stu)
}
