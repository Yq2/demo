package main

import (
	"math/rand"
)

//func init() {
//	rand.Seed(time.Now().UnixNano())
//}
//func main() {
//	fmt.Println(RandStringBytes(32))
//}

const letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// 字节替换rune
// 如果需求是只使用英语字母字符(包括大小写)，那么我们可以使用byte替换rune,
// 因为UTF-8编码中英语字母和byte是一一对应的。

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
