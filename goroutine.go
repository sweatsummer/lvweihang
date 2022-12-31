package main

import "fmt"

// a原因：在匿名函数内第一次打印“有点强人锁男”后该进程上锁，导致无法打印后续//
// b 修改如下//
func main() {
	cha := make(chan int)

	go func() {
		fmt.Println("下山的路又堵起了")
		<-cha
	}()
	cha <- 1

}
