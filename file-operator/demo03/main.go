package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//使用ioutil.ReadFile一次性将文件读取到位
	file := "d:/tesf.txt"
	content, err := ioutil.ReadFile(file)

	if err != nil {
		fmt.Printf("%v", content)
	}

	//把读取到的内容显示到终端
	//fmt.Printf("%v", content) //[]byte
	fmt.Printf("%v", string(content)) //[]byte
	//我们没有显式的open文件，因此也不需要显式的close文件
	//因为，文件的open和close被封装到 ReadFile 函数内部
}
