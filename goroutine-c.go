package main

import (
	"fmt"
	"time"
)

var ch = make(chan int)

// (c)//
func main() {
	for i := 0; i <= 9; i++ {
		go Print(i)
	}
	time.Sleep(time.Second)
	ch <- 1
	fmt.Println("主程序结束...")
}

func Print(a int) {
	fmt.Println(a)
	<-ch
}
