//至少编写3个计时器，完成以下功能：
//
//●  在每天凌晨两点，输出"我还能再战4小时！"
//●  在每天早上六点，输出"我要去图书馆开卷！"
//●  自程序运行时起，每过半分钟输出"芜湖！起飞！"
//●  其他你想完成的功能
//
//文案可自行修改，测试时可以把时间间隔缩短。

package main

import (
	"fmt"
	"time"
)

func timer() {
	ticker := time.NewTicker(time.Second * 30)
	for {
		select {
		case <-ticker.C:
			fmt.Println("芜湖！起飞！")
		}
	}
}
func timerss() {
	ticker2 := time.NewTicker(time.Hour * 24)
	for {
		select {
		case <-ticker2.C:
			fmt.Println("我还能再战4小时！")
		}
	}
}
func timers() {
	ticker1 := time.NewTicker(time.Hour * 4)
	for {
		select {
		case <-ticker1.C:
			fmt.Println("我要去图书馆开卷！")
		}
	}
}
func main() {
	//在06：00：00设定计时器即可达到效果//
	go timer()
	go timers()
	timerss()
}
