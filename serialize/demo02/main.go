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

// 将字符串反序列化为结构体
func deserializeStruct() {
	//str := "{\"Name\":\"牛魔王\",\"Age\":500,\"Birthday\":\"2011-11-11\",\"Sal\":8000,\"Ski11\":\"牛魔拳\"}"
	str := "{\"monster_name\":\"牛魔王\",\"monster_age\":500,\"monster_birthday\":\"2011-11-11\",\"monster_sal\":8000,\"monster_skill\":\"牛魔拳\"}"

	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Printf("反序列化错误 err%v\n", err)
	}
	fmt.Printf("字符串反序列化后monster=%v\n", monster)
}

// 将字符串反序列化为map
func deserializeMap() {
	str := "{\"address\":\"火云洞\",\"age\":30,\"name\":\"红孩儿\"}"

	var a map[string]interface{}

	//反序列化
	//注意:反序列化map,不需要make,因为make操作被封装到 Unmarshal函数中
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Printf("反序列化错误 err%v\n", err)
	}
	fmt.Printf("字符串反序列化后a=%v\n", a)
}

// 将字符串反序列化为slice
func deserializeSlice() {
	str := "[{\"address\":\"北京\",\"age\":\"7\",\"name\":\"jack\"}," +
		"{\"address\":[\"墨西哥\",\"夏威夷\"],\"age\":\"20\",\"name\":\"tom\"}]"

	var slice []map[string]interface{}

	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Printf("反序列化错误 err%v\n", err)
	}
	fmt.Printf("字符串反序列化后slice=%v\n", slice)

}

func main() {
	deserializeStruct()

	deserializeMap()

	deserializeSlice()
}
