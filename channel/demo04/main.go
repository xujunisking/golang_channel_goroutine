package main

import (
	"fmt"
	"time"
)

//需求:要求统计1-8000的数字中，哪些是素数?这个问题在本章开篇就提出了，
//现在我们有goroutine和channel的知识后，就可以完成了[测试数据:80000]

func putNum(intChan chan int) {

	for i := 1; i <= 8000; i++ {
		intChan <- i
	}
	close(intChan)
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {

	var flag bool //是素数
	for {
		num, ok := <-intChan
		if !ok {
			break
		}

		flag = true //假设是素数
		//判断num是不是素数
		for i := 2; i < num; i++ {
			if num%i == 0 { //说明该num不是素数
				flag = false
				break
			}
		}

		if flag {
			primeChan <- num
		}
	}

	fmt.Println("有一个primeNum协程因为取不到数据而退出")
	//这里还不能关闭primeChan

	exitChan <- true
}

func main() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 2000)
	//标识退出的管理
	exitChan := make(chan bool, 4)

	start := time.Now().Unix()

	//开启一个协程，向intChan放入1-8000个数
	go putNum(intChan)
	//开启4个协程，从 intchan取出数据，并判断是否为素数,如果是，就
	//放入到primechan
	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}

		end := time.Now().Unix()
		fmt.Println("使用协程耗时=", end-start)
		//当我们从exitChan 取出了4个结果，就可以放心的关闭primeNun
		close(primeChan)
	}()

	//primeChan,把结果取出
	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Printf("素数=%d\n", res)
	}

	fmt.Println("主函数退出！")
}
