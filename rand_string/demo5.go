package main

import (
	"math/rand"
)

//func init() {
//	rand.Seed(time.Now().UnixNano())
//}

const letterBytesPlus = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBitsPlus = 6                        // 6 bits to represent a letter index
	letterIdxMaskPlus = 1<<letterIdxBitsPlus - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMaxPlus  = 63 / letterIdxBitsPlus   // # of letter indices fitting in 63 bits
)

// 掩码增强版
func RandStringBytesMaskImpr(n int) string {
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, rand.Int63(), letterIdxMaxPlus; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMaxPlus
		}
		if idx := int(cache & letterIdxMaskPlus); idx < len(letterBytesPlus) {
			b[i] = letterBytesPlus[idx]
			i--
		}
		cache >>= letterIdxBitsPlus
		remain--
	}
	return string(b)
}

//func main() {
//	fmt.Println(RandStringBytesMaskImpr(32))
//}
