package main

import "fmt"

func main() {

	//演示一下管道的使用
	//1.创建一个可以存放3个int类型的管道
	var intChan chan int = make(chan int, 3)

	//2.intchan是什么
	fmt.Printf("intChan的值=%v intChan本身的地址=%p\n", intChan, &intChan)

	//3.向管道写入数据
	intChan <- 10
	num := 211
	intChan <- num
	intChan <- 50
	//intChan <- 98 //注意点，当我们给管写入数据时，不能超过其容量

	//4.查看管道的长度和cap(容量)
	fmt.Printf("channel len=%v cap=%v \n", len(intChan), cap(intChan))

	//5.从管道中读取数据
	var num2 = <-intChan
	fmt.Println("num2=", num2)
	fmt.Printf("channel len=%v cap=%v \n", len(intChan), cap(intChan))

	//6.在没有使用协程的情况下，如果我们的管道数据已经全部取出，再取就会报告 deadlock
	num3 := <-intChan
	num4 := <-intChan
	fmt.Println("num3=", num3, "num4=", num4)

}
