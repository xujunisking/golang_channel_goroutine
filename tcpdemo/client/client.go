package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	fmt.Println("客户端启动，开始链接服务器端...")
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("client dial err=", err)
		return
	}

	//fmt.Println("链接成功 conn=", conn)
	reader := bufio.NewReader(os.Stdin) //os.Stdin 代表标准输入[终端]

	for {

		fmt.Println("客户端等待输入信息...")
		//从终端读取一行用户输入， 并准备发送给服务器
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readString err=", err)
		}

		if line == "exit\r\n" { //注意客户端不仅仅是输入了字符，还要加上回车换行符(\r\n)
			fmt.Printf("客户端退出...")
			return
		}

		//将line发送给服务器
		n, err := conn.Write([]byte(line))
		if err != nil {
			fmt.Println("conn.Write err=", err)
			return
		}
		fmt.Printf("客户端发送了 %d 字节的数据\n", n)
	}
}
