// lv2: 试图使用接口来描述一个系统该有的行为，然后实现这些接口，例如：动物都有发声这个行为，狗会吠吠叫，猫会喵喵叫。 //

// level 2//

package main

import "fmt"

// 结构体（学校的年龄，工作，娱乐）//

type school struct {
	age   int
	job   string
	enjoy string
}

//建立接口 workIn //

type workIn interface {
	work()
	party()
}

// 接口1//
func (T *school) work() {
	fmt.Println("school can " + T.job)
}

// 接口2//
func (T *school) party() {
	fmt.Println("school also can " + T.enjoy)
}

func main() {
	// 结构体初始化 //
	T := school{
		age:   70,
		job:   "teach",
		enjoy: "hold activities",
	}
	fmt.Println(T)
	T.work()
	T.party()
}
