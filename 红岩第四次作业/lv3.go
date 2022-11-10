// ● 程序启动时，从本地文件（文件名随意，建议使用users.data）读入用户名和密码并保存到map中
// ●  用户注册时，判断用户名密码是否符合规范并输出提示信息，规则是用户名不能重复，密码长度大于6位（可自行增加其他规则）。
// ●  用户选择退出程序（不是人为强行终止程序）时，将新注册的用户信息也写入到文件中，然后退出程序
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var user map[int]string = map[int]string{}
var password map[int]string = map[int]string{}
var of bool

func main() {
tmq:
	file := "D:/user.data.txt"
	//一次性读取//
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("%v", err)
	}
	s := string(content)
	a := strings.Split(s, " ")
	//分别保存用户名与密码//
	user[1] = a[1]
	password[1] = a[3]
	//fmt.Printf("用户名：%v,密码：%v", user[1], password[1])
	var name, pass string
tmp:
	fmt.Println("请输入注册用户名：")
	fmt.Scan(&name)
	for i := 0; i <= len(user)-1; i++ {
		if name == user[i] {
			fmt.Println("用户名不可重复")
			goto tmp
		}
	}
tmt:
	fmt.Println("请设置密码：")
	fmt.Scan(&pass)
	if len(pass) < 6 {
		fmt.Println("密码长度过小，应大于六位")
		goto tmt
	}
	filepath := "d:/user.data.txt"
	constent, err := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	defer constent.Close()
	fmt.Println("是否退出程序（true or false)：")
	fmt.Scan(&of)
	if of == false {
		goto tmq
	}
	fmt.Println("注册成功，已写入\n")
	str := name
	str1 := pass
	writer := bufio.NewWriter(constent)
	writer.WriteString(str)
	writer.WriteString(str1)
	writer.Flush()
}

//func login() {
//	var name, pass string
//tmp:
//	fmt.Println("请输入注册用户名：")
//	fmt.Scan(&name)
//	if name == user[1] {
//		fmt.Println("用户名不可重复")
//		goto tmp
//	}
//tmt:
//	fmt.Println("请设置密码：")
//	fmt.Scan(&pass)
//	if len(pass) < 6 {
//		fmt.Println("密码长度过小，应大于六位")
//		goto tmt
//	}
//	filepath:="d:/user.data.txt"
//	constent,err =os.OpenFile(filepath,os.O_WRONLY|os.O_APPEND,0666)
//	if err != nil {
//		fmt.Printf("open file err=%v\n",err)
//	}
//}
