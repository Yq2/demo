package main

import (
	"fmt"
	"github.com/spaolacci/murmur3"
)

var bucketSize = uint64(10)

func main() {
	var bucketMap = map[uint64]int{}
	var start = 150000000
	for i := start; i < start+10000000; i++ {
		hashInt := murmur3.Sum64([]byte(fmt.Sprint(i))) % bucketSize
		bucketMap[hashInt]++
	}
	fmt.Println(bucketMap)
}
