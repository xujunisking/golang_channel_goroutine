package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	file, err := os.Open("d:/test.txt")
	if err != nil {
		fmt.Println("open file err=", err)
	}

	//当函数退出时，要及时的关闭file
	defer file.Close()

	// 创建一个，是带缓冲的 *Reader
	/*
		const (
			defaultBufsize = 4096 //默认的缓冲区为4096
		)
	*/
	reader := bufio.NewReader(file)
	//循环的读取文件的内容
	for {
		str, err := reader.ReadString('\n') // 读到一个换行就结束
		if err == io.EOF {                  // i0.EOF表示文件的末尾
			break
		}
		//输出内容
		fmt.Print(str)
	}

	fmt.Println("文件读取结束...")
}
