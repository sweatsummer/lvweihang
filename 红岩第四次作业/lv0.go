// 聊天屏蔽敏感词//
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//创建敏感词库//
	var m map[int]string = map[int]string{1: "cxk", 2: "sb", 3: "邮专", 4: "cnm"}
	fmt.Println("请输入你想说的话：")
	reader := bufio.NewReader(os.Stdin)
	sentence, _ := reader.ReadString('\n')
	//建立循环判别//
	for i := 1; i <= len(m); i++ {
		start := strings.Contains(sentence, m[i])
		if start {
			fmt.Printf("该字段含有敏感词%v\n", m[i])
			//替换相应敏感词//
			sentence = strings.Replace(sentence, m[i], "**", -1)
		}
	}
	fmt.Println(sentence)
}
