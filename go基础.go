package main

import "fmt"

type people struct {
	name string
}

// (a)错误原因 变量名首字母小写未能转化为json格式; 还应当引入"encoding/json"包//
// (b)原因：2/0在前，分母做0出错，将两打印交换顺序即可打印出1//
func main() {
	a := 2
	b := 0

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	fmt.Println("3/2 =", 3/2)
	fmt.Println("2/0=", a/b)

}
