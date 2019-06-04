package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	// 基于当前时间纳秒数初始化随机数种子
	rand.Seed(time.Now().UnixNano())
}

// 洗牌算法
func shuffle(n int) []int {
	b := rand.Perm(n)
	return b
}

func main() {
	for i := 1; i < 15; i++ {
		fmt.Println(shuffle(i))
	}
}
