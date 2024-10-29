package main

import "fmt"

func main() {

	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 200
	close(intChan) //close
	//这是不能够再写入数到channe1
	//intchan <- 300 //panic: send on closed channel
	fmt.Println("okook~'")

	//当管道关闭后，读取数据是可以的
	n1 := <-intChan
	fmt.Println("n1=", n1)
	n2 := <-intChan
	fmt.Println("n2=", n2)

	//遍历管道
	intChan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan2 <- i * 2 //放[入100个数据到管道
	}

	//遍历(不能用这种方式遍历管道，因为遍历后管道的长度是减少的)
	// for i :=0;i<len(intchan2); i++ {}
	//在遍历时，如果channel没有关闭，则会出现deadlock的错误
	//关闭管道
	close(intChan2)
	for v := range intChan2 {
		fmt.Println("v=", v)
	}
}
