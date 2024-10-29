package main

import "fmt"

//1)启动一个协程，将 1-2000的数放入到一个channel中,比如numChan
//2)启动8个协程，从numChan取出数(比如n)，并计算 1+..+n的值，并存放到resChan
//3)最后8个协程协同完成工作后，再遍历resChan,显示结果[如res[1]=1.. res[10]=55...]
//4)注意:考虑resChan chan int是否合适?

func writeData(numChan chan int) {
	for i := 1; i <= 2000; i++ {
		numChan <- i
		fmt.Println("writeData=", i)
	}
	close(numChan)
}

func calNum(intChan chan int, resChan chan int, exitChan chan bool) {
	for {
		count, ok := <-intChan
		if !ok {
			break
		}

		var result int
		for i := 1; i <= count; i++ {
			result += i
		}

		resChan <- result
	}

	exitChan <- true
	close(exitChan)
}

func main() {
	numChan := make(chan int, 2000)
	resChan := make(chan int, 2000)
	exitChan := make(chan bool, 1)

	go writeData(numChan)

	for i := 1; i <= 8; i++ {
		go calNum(numChan, resChan, exitChan)
	}

	go func() {
		for i := 1; i <= 8; i++ {
			<-exitChan
		}
		close(resChan)
	}()

	for v := range resChan {
		fmt.Printf("累加结果=%v\n", v)
	}
}
