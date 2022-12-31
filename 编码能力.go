package main

import (
	"fmt"
)

//代码可能有点小问题，导致输入不完整，在此尽量展示思路//

func main() {
	var n, m int
	var a [1000][100]int
	var b [1000][100]int
	var c [100]int
	fmt.Scanf("%d %d", &n, &m)
	for i := 1; i <= n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scanf("%d %d", &a[i][j], &b[i][j])
			//计算第i层梯子个数//
			c[i] += a[i][j]
		}
	}
	j2, d, N := 0, 0, 0
	for i := 1; i < n+1; i++ {
		d = d + b[i][j2]
		M := b[i][j2] % c[i]
		//找到M个梯子房间//
		if M == 0 {
			M = c[i]
		}
		for {
			if a[i][j2] == 1 {
				N++
				if N == M {
					break
				}
			}
			//模拟走到下一个房间//
			j2 = (j2 + 1) % m
		}
	}
	//输出结果//
	fmt.Printf("%v\n", d)
}
