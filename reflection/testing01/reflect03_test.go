package testing01

import (
	"reflect"
	"testing"
)

type person struct {
	Userld string
	Name   string
}

func TestReflectStructPtr(t *testing.T) {
	var (
		model *person
		st    reflect.Type
		elem  reflect.Value
	)
	st = reflect.TypeOf(model)
	t.Log("reflect.TypeOf", st.Kind().String()) //struct
	st = st.Elem()                              //st指向的类型
	t.Log("reflect.TypeOf.Elem", st.Kind().String())
	//New返回一个Value类型值,该值持有持有一个指向类型为typ的新申请的零值的指针
	elem = reflect.New(st)
	t.Log("reflect.New", elem.Kind().String())             //ptr
	t.Log("reflect.New.Elem", elem.Elem().Kind().String()) //struct
	//model就是创建的user结构体变量(实例)
	model = elem.Interface().(*person)               //model是*person它的指向和elem是一样的.
	elem = elem.Elem()                               //职得elem指向的值
	elem.FieldByName("Userld").SetString("12345678") //赋值..
	elem.FieldByName("Name").SetString("nickname")
	t.Log("model model.Name", model, model.Name)
}
