// lv1: 使用结构体描述一个事物的属性（结构体字段）和行为（方法），
// 作业里需要给出描述的对象信息注释（例如：某种动物，网站等

// level 0 略//

// level 1//
package main

import "fmt"

// 网站结构体（年龄，id，风格,状态）//

type Website struct {
	age       int
	id        string
	style     string
	condition string
}

// 函数方法（bilbil 的状态）//
func (bilbil Website) open() {
	bilbil.condition = "fascinating"
	fmt.Printf("The condition of bilbil is %s\n", bilbil.condition)
}

// 结构体方法 //
func open(bilbil Website) {
	bilbil.condition = "great"
	fmt.Printf("The condition of bilbil is %s\n", bilbil.condition)

}

func main() {
	// website的初始化//
	var bilbil = Website{
		age:       13,
		id:        "bilbil",
		style:     "young and creative",
		condition: "normal",
	}
	bilbil.open()
	open(bilbil)
	fmt.Println(bilbil)
}
