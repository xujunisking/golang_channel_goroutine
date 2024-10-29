package main

import (
	"fmt"
	"sync"
	"time"
)

//需求:现在要计算 1-208 的各个数的阶乘，并且把各个数的阶乘放入到map中。
//最后显示出来。要求使用goroutine完成

// 思路
// 1.编写一个函数，来计算各个数的阶乘，并放入到 map中,
// 2.我们启动的协程多个，统计的将结果放入到map中
// 3.map 应该做出一个全局的.
var (
	myMap = make(map[int]int, 10)
	//声明一个全局的互斥锁
	//sync: synchornized同步
	//Mutex: 互斥
	lock sync.Mutex
)

func test(n int) {

	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}

	//这里我们将res 放入到myMap
	lock.Lock() //加锁
	myMap[n] = res
	lock.Unlock() //解锁
}

func main() {

	//开启多个协程完成这个任务[20个]
	for i := 1; i <= 20; i++ {
		go test(i) //这里会出现资源竞争的问题
	}

	time.Sleep(time.Second * 5)

	//这里我们输出结果,变量这个结果
	lock.Lock() //加锁
	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
	lock.Unlock() //解锁
}

//在编译该程序时，增加一个参数 -race即可知道是否存在资源竞争问题
//go run -race main.go
