package main

import (
	"math/rand"
)

//func init() {
//	rand.Seed(time.Now().UnixNano())
//}
const letterBytesMask = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
)

// 掩码方案
func RandStringBytesMask(n int) string {
	b := make([]byte, n)
	for i := 0; i < n; {
		if idx := int(rand.Int63() & letterIdxMask); idx < len(letterBytesMask) {
			b[i] = letterBytesMask[idx]
			i++
		}
	}
	return string(b)
}

//func main() {
//	fmt.Println(RandStringBytesMask(32))
//
//}
