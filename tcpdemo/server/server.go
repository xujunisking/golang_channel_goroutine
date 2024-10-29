package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {

	//循环接收客户端发送的数据
	defer conn.Close() //关闭conn
	for {
		//创建一个新的切片
		buf := make([]byte, 1024)
		//conn.Read(buf)
		//1.等待客户端通过conn发送信息
		//2.如果客户端美欧Write[发送]，协程就阻塞在这里
		fmt.Printf("服务器在等待客户端%s 发送信息\n", conn.RemoteAddr().String())
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("客户端已退出", err)
			return
		}
		//3.显示客户端发送的内容到服务器的终端
		fmt.Print(string(buf[:n]))
	}
}

func main() {

	fmt.Println("服务器开始监听了......")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err=", err)
		return
	}
	defer listen.Close() //延时关闭listen

	//循环等待客户端来链接
	for {
		//等待客户端链接
		fmt.Println("等待客户端链接......")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err=", err)
		} else {
			fmt.Printf("Accept() suc con=%v 客户端IP=%v\n", conn, conn.RemoteAddr().String())
		}
		//这里准备启动一个协程，为客户端服务
		go process(conn)
	}
	//fmt.Printf("listen suc=%v\n", listen)
}
