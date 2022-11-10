// 要求在使用ReleaseSkill函数的情况下，编写一个程序，程序流程如下：
// 1.  用户选择释放技能时输出文字的模板（这就要求你在程序中编写数个匿名函数/不匿名函数来使用）
// 2.  用户输入技能名，程序调用ReleaseSkill函数，输出技能释放时文字
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var number int
	var i int
	for i == 0 {
		fmt.Println("请输入你想使用的技能名：\n") //技能名输入//
		reader := bufio.NewReader(os.Stdin)
		s1, _ := reader.ReadString('\n')
		fmt.Println("排版格式\n")
		option1() //各三个模板//
		option2()
		option3()
		fmt.Println("请选择你的技能排版(1,2,3):\n")
		fmt.Scanf("%v\n", &number)
		if number == 1 {
			ReleaseSkill(s1, func(skillName string) {
				fmt.Println("尝尝我的厉害吧! ", skillName)
			})
		}
		if number == 2 {
			ReleaseSkill(s1, func(skillName string) {
				fmt.Println("\n尝尝我的厉害吧!\n", skillName)
			})
		}
		if number == 3 {
			func() {
				fmt.Printf(" %v\n尝尝我的厉害吧！\n", s1)
			}()
		}
	tmp:
		fmt.Println("还想尝试其他的内容吗（继续0 or 结束1):") //可选择是否退出//
		fmt.Scanf("%v", &i)
		if i != 0 && i != 1 {
			fmt.Println("命令未知...请重新输入:")
			goto tmp
		}
	}
	fmt.Println("程序结束...\n")
}
func ReleaseSkill(skillNames string, releaseSkillFunc func(string)) {
	releaseSkillFunc(skillNames)
}
func option1() {
	fmt.Println("1:技能名+尝尝我的厉害吧！\n")
}
func option3() {
	fmt.Println("2:技能名\n  尝尝我的厉害吧!\n\n")
}

func option2() {
	fmt.Println("3:尝尝我的厉害吧!\n  技能名\n\n")
}
