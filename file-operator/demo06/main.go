package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	filePath := "d:/abc.txt"
	//注意这里是必须存在的文件 打开文件并在文件末尾追加文件内容
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 8666)

	if err != nil {
		fmt.Printf("open file err=%v\r\n", err)
		return
	}

	defer file.Close()

	str := "ABC,ENGLISH！！！\r\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}

	writer.Flush()
}
