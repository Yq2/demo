package main

import (
	"math/rand"
	"time"
)

// 按照概率发布

func init() {
	// 基于当前时间纳秒数初始化随机数种子
	rand.Seed(time.Now().UnixNano())
}

func isPassed(rate int) bool {
	if rate >= 100 {
		return true
	}
	if rate > 0 && rand.Int() > rate {
		return true
	}
	return false
}
