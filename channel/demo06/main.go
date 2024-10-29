package main

import "fmt"

func main() {
	//使用select可以解决从管道取数据的阻塞问题
	//1.定义一个管道 10个数据int
	intchan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intchan <- i
	}

	//2.定义一个管道 5个数据string
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}

	//传统的方法在遍历管道时，如果不关闭会阻塞而导致 deadlock

	//问题，在实际开发中，可能我们不好确定什么关闭该管道，
	//可以使用select 方式可以解决
	//label:
	for {
		select {
		//注意:这里，如果intchan一直没有关闭，不会一直阻塞而deadlock
		case v := <-intchan:
			fmt.Printf("从intChan读取的数据%d\n", v)
		case v := <-stringChan:
			fmt.Printf("从stringChan读取的数据%s\n", v)
		default:
			fmt.Printf("都取不到了！！！")
			return
			//break label
		}
	}
}
