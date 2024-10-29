package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string  `json:"monster_name"`
	Age      int     `json:"monster_age"`
	Birthday string  `json:"monster_birthday"`
	Sal      float64 `json:"monster_sal"`
	Skill    string  `json:"monster_skill"`
}

// 将结构体进行序列化
func testStruct() {
	monster := Monster{
		Name:     "牛魔王",
		Age:      500,
		Birthday: "2011-11-11",
		Sal:      8000.0,
		Skill:    "牛魔拳",
	}

	//将monster 序列化
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("序列化错误 err%v\n", err)
	}
	fmt.Printf("monster序列化后=%v\n", string(data))
}

// 将map进行序列化
func testMap() {
	//定义一个map key为string  value为任意类型
	//使用map之前需要make
	a := make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 30
	a["address"] = "火云洞"

	data, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("序列化错误 err%v\n", err)
	}
	fmt.Printf("a map序列化后=%v\n", string(data))
}

// 将切片进行序列化
func testSlice() {
	var slice []map[string]interface{}
	m1 := make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = "7"
	m1["address"] = "北京"
	slice = append(slice, m1)

	m2 := make(map[string]interface{})
	m2["name"] = "tom"
	m2["age"] = "20"
	m2["address"] = "上海"
	slice = append(slice, m2)

	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("序列化错误 err%v\n", err)
	}
	fmt.Printf("slice 序列化后=%v\n", string(data))
}

// 将基本数据类型序列化,没有实际意义
func testFloat64() {
	var num1 float64 = 2345.67

	data, err := json.Marshal(num1)
	if err != nil {
		fmt.Printf("序列化错误 err%v\n", err)
	}
	fmt.Printf("num1 序列化后=%v\n", string(data))
}

func main() {
	testStruct()

	testMap()

	testSlice()

	testFloat64()
}
