package main

import (
	"math/rand"
	"time"
)

const letterBytesImprSrc = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBitsImprSrc = 6                           // 6 bits to represent a letter index
	letterIdxMaskImprSrc = 1<<letterIdxBitsImprSrc - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMaxImprSrc  = 63 / letterIdxBitsImprSrc   // # of letter indices fitting in 63 bits
)

// 优化随机数产生
// rand.Seed使用的是全局随机数，并发安全但会有锁的开销
var src = rand.NewSource(time.Now().UnixNano())

func RandStringBytesMaskImprSrc(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMaxImprSrc; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMaxImprSrc
		}
		if idx := int(cache & letterIdxMaskImprSrc); idx < len(letterBytesImprSrc) {
			b[i] = letterBytesImprSrc[idx]
			i--
		}
		cache >>= letterIdxBitsImprSrc
		remain--
	}
	return string(b)
}

//func main() {
//	fmt.Println(RandStringBytesMaskImprSrc(32))
//}
