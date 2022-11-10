// ● 用户可以自定义每个技能使用的文字模板，并且用户在释放技能前不再需要指定输出文字的模板
// ●  用户可以添加自定义的技能名字，判断用户输入的技能名字中是否有敏感词（敏感词词库自行设置），如果包含敏感词则输出对应的提示文字并取消本次添加。添加后也可以自定义每个技能使用的文字模板
// ●  用户可以添加自定义的模板，并且同样需要对敏感词鉴权，添加后用户可以使用该模板。
// 希望模板的自定义程度不要仅仅局限于只能自定义文字，也要可以自定义技能名出现的位置和次数，如："你死定了！"+技能名+"我不知道你该怎么从"+技能名+"活下来~"
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	//自定义敏感词库//
	var data map[int]string = map[int]string{1: "鸡你太美", 2: "cxk", 3: "114514", 4: "sb", 5: "邮专", 6: "cnm"}
	var mod map[int]string = map[int]string{}
	var T, k int
	for q := 1; q >= 1; q++ {
		fmt.Printf("目前有%v个技能模板\n", q-1)
		if q >= 1 {
			fmt.Println("请选择添加自定义模板或使用释放已储存的技能模板(添加1 释放2）：")
			fmt.Scanf("%v", &T)
			if T == 2 {
				fmt.Println("请输入你要释放的模板号：")
				time.Sleep(time.Millisecond * 100)
				fmt.Scanf("%d", &k)
				releaseskill(mod[k])
			}
		}
	tmp:
		fmt.Println("按任意键继续...")
		fmt.Scanf("%v")
		fmt.Println("请输入你的技能名：")
		reader2 := bufio.NewReader(os.Stdin)
		n1, _ := reader2.ReadString('\n')
		//筛查敏感词//
		for i := 1; i <= len(data); i++ {
			of := strings.Contains(n1, data[i])
			if of == true {
				fmt.Printf("在技能名中检测到敏感词：%v，请重新输入:", data[i])
				goto tmp
			}
		}
		fmt.Println("请自定义你的技能文字模版(用”技能名“代替技能名的名字,可自定义技能名出现的次数和其他语句，如“技能名，尝尝我的厉害吧！”）:")
	tmq:
		reader1 := bufio.NewReader(os.Stdin)
		mod1, _ := reader1.ReadString('\n')
		mod1 = strings.Replace(mod1, "技能名", n1, -1)
		//筛查敏感词//
		for j := 1; j <= len(data); j++ {
			off := strings.Contains(mod1, data[j])
			if off == true {
				fmt.Printf("在模板中检测到敏感词：%v，请重新输入:", data[j])
				goto tmq
			}
		}
		mod[q] = mod1
		fmt.Printf("已保存有关技能:%v 到模板%v\n", n1, q)
	}
}
func releaseskill(m string) {
	fmt.Println(m)
	fmt.Println("技能已释放")
}
