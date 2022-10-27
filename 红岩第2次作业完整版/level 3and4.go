//level 3 （冒泡排序）//

package main

import (
	"fmt"
)

func array() {
	var array = []int{9, 6, -8, 5, 2, 1, 11, 0, 54, 2, 1, 34, 1, 9}
	fmt.Println(array)
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			tmp := 0
			//比较相邻两数的大小，判断是否交换//
			if array[j] > array[j+1] {
				//引入变量tmp实现两数数值交换//
				tmp = array[j]
				array[j] = array[j+1]
				array[j+1] = tmp
			}
		}
	}
	fmt.Println(array)
}

// leve 5//
// ○ 定义一个coder结构体，包括以下信息：姓名，学号，手机号，代码行数，出生日期等一切你能想到属性
type coder struct {
	name        string
	idnumber    int
	phonenumber int
	num         int
	birthtime   int
}

//  建立结构体数组或链表，实现对多个同学的信息输入，输出 //

func main() {
	S := [4]coder{
		{name: "lihua", idnumber: 100, phonenumber: 1234, num: 160, birthtime: 2004},
		{name: "xiaohong", idnumber: 101, phonenumber: 1235, num: 160, birthtime: 2004},
		{name: "pengpeng", idnumber: 102, phonenumber: 1236, num: 140, birthtime: 2004},
		{name: "xiaoma", idnumber: 103, phonenumber: 1237, num: 200, birthtime: 2003},
	}
	var m int
	//统计做题数大于150的人，打印完整信息//
	for m = 0; m <= len(S)-1; m++ {
		if S[m].num > 150 {
			fmt.Println(S[m])
		}
	}
	//按学号查询，打印完整信息//
	var a, n int
	fmt.Println("please input what do you seek(idname):")
	fmt.Scanf("%v", &a)
	for n = 0; n <= len(S)-1; n++ {
		if S[n].idnumber == a {
			fmt.Printf("Information:%v\n", S[n])
		}
	}
}

//○ 实现简单的统计功能，比如统计做题数大于150的同学并输出其完整信息
//○ 实现查找功能，包括按姓名、学号查找
//○ 实现信息修改功能（肝不动了。。。）
