package main

import (
	"fmt"
)

//(d) 素数计算//

func putNum(intChannel chan int) {
	for i := 0; i < 1000000; i++ {
		// 初始化数据放入管道
		intChannel <- i
	}
	close(intChannel)
}

func primeNum(intChannel chan int, primeChannel chan int, exitChannel chan bool) {
	for {
		// 从管道中取数据进行计算
		num, ok := <-intChannel
		if !ok {
			// 管道为空
			break
		}
		flag := true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				continue
			}
		}
		if flag {
			// 素数放入管道
			primeChannel <- num
		}
	}
	exitChannel <- true
}

func main() {
	numCPU := 16

	intChannel := make(chan int, 10000)
	primeChannel := make(chan int, 2000)
	exitChannel := make(chan bool, numCPU)

	// 单线程初始化管道
	go putNum(intChannel)
	for i := 0; i < numCPU; i++ {
		go primeNum(intChannel, primeChannel, exitChannel)
	}

	go func() {
		for i := 0; i < numCPU; i++ {
			<-exitChannel
		}
		close(primeChannel)
	}()

	// 单线程从计算结果中取数据进行打印
	for {
		result, ok := <-primeChannel
		if !ok {
			break
		}
		fmt.Println("素数：", result)
	}
	fmt.Println("主线程结束")
}
