package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { //文件或者目录存在
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func main() {
	//将d:/abc.txt 文件内容导入到 e:/kkk.txt

	file1Path := "d:/abc.txt"
	file2Path := "e:/kkk.txt"
	data, err := ioutil.ReadFile(file1Path)
	if err != nil {
		//说明读取文件有错误
		fmt.Printf("read file err=%v\n", err)
		return
	}

	err = ioutil.WriteFile(file2Path, data, 0666)
	if err != nil {
		fmt.Printf("write file error=%v\n", err)
	}
}
