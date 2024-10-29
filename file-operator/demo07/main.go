package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	filePath := "d:/abc.txt"
	//注意这里是必须存在的文件 打开文件并在文件末尾追加文件内容
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 8666)

	if err != nil {
		fmt.Printf("open file err=%v\r\n", err)
		return
	}

	defer file.Close()

	//先议取原来文件的内容，并显示在终端
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF { //如果读取到文件的末尾
			break
		}
		//显示到终端
		fmt.Print(str)
	}

	str := "Hello,昆明！！！\r\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}

	writer.Flush()
}
