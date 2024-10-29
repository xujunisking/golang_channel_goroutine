package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	/*
		const (
			O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
			O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
			O_RDWR   int = syscall.O_RDWR   // 读写模式打开文件
			O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
			O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
			O_EXCL   int = syscall.O_EXCL   // 和O_CREATE配合使用，文件必须不存在
			O_SYNC   int = syscall.O_SYNC   // 打开文件用于同步I/O
			O_TRUNC  int = syscall.O_TRUNC  // 如果可能，打开时清空文件
		)
	*/

	filePath := "d:/abc.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 8666)

	if err != nil {
		fmt.Printf("open file err=%v\r\n", err)
		return
	}

	defer file.Close()

	//准备写入5句"hello，Gardon'
	str := "hello,Gardon"
	//写入时，使用带缓存的*writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	//因为writer是带缓存，因此在调用writerstring方法时，其实
	//内容是先写入到缓存的,所以需要调用Flush方法，将缓冲的数据
	//真正写入到文件中,[否则文件中会没有数据
	writer.Flush()

}
