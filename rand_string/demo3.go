package main

import (
	"math/rand"
)

const letterBytesRmndr = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

//func main()  {
//
//	fmt.Println(RandStringBytesRmndr(32))
//}

//func init() {
//	rand.Seed(time.Now().UnixNano())
//}

// 使用余数
func RandStringBytesRmndr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytesRmndr[rand.Int63()%int64(len(letterBytesRmndr))]
	}
	return string(b)
}
