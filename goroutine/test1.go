// go协程//
// 请你用Channel实现对课件中并发示例代码的改写(即用Channel操作代替加锁操作)
package main

import (
	"fmt"
	"sync"
)

var x int64
var wg sync.WaitGroup
var mu sync.Mutex
var buffer = make(chan int)
var recive, i int

func add() {
	for i := 0; i < 5000; i += 1 {
		x = x + 1
	}
	buffer <- i
	wg.Done()
}
func main() {
	wg.Add(2)
	go add()
	recive = <-buffer
	go add()
	recive = <-buffer
	wg.Wait()
	fmt.Println(x)
	fmt.Println("程序结束\n")
}
