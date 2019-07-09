package main

import (
	"math/rand"
)

//func init() {
//	rand.Seed(time.Now().UnixNano())
//}

var letterRunes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

//func main() {
//
//	// 1 通用的方案
//	fmt.Println(RandStringRunes(32))
//}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
