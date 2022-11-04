// 使用两个goroutine轮流打印100以内的数
//
// 比如一个goroutine打印1，另一个打印2，(即一个打印奇数，一个打印偶数)，这样打印下去到100
//
// 注意，要保证顺序
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg1 sync.WaitGroup
var channl = make(chan int)
var channl1 = make(chan int)

func even() { //打印奇数//
	for j := 1; j <= 99; j += 2 {
		<-channl
		fmt.Printf("odd number:%v\n", j)
		channl1 <- 1
	}
	wg1.Done()
}

func odd() { //打印偶数//
	for q := 2; q <= 100; q += 2 {
		channl <- 1
		time.Sleep(time.Millisecond * 10) //0.01秒延迟，防止出现奇偶反转//
		fmt.Printf("even number:%v\n", q)
		<-channl1
	}
	wg1.Done()
}

func main() {

	wg1.Add(2)
	//go print()
	//go print()
	go even() //1
	go odd()  //2
	wg1.Wait()
	fmt.Println("程序结束...\n") //3
}
