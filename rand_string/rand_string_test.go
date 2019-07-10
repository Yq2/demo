package main

import (
	"math/rand"
	"testing"
	"time"
)

const n = 32

func init() {
	rand.Seed(time.Now().UnixNano())
}

func BenchmarkRandStringBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStringBytes(n)
	}
}

func BenchmarkRandStringRunes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStringRunes(n)
	}
}

func BenchmarkRandStringBytesRmndr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStringBytesRmndr(n)
	}
}

func BenchmarkRandStringBytesMask(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStringBytesMask(n)
	}
}

func BenchmarkRandStringBytesMaskImpr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStringBytesMaskImpr(n)
	}
}

func BenchmarkRandStringBytesMaskImprSrc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStringBytesMaskImprSrc(n)
	}
}

//BenchmarkRandStringBytesMaskImprSrc-4           10000000               206 ns/op              64 B/op          2 allocs/op
//BenchmarkRandStringBytesMaskImpr-4               5000000               246 ns/op              64 B/op          2 allocs/op
//BenchmarkRandStringBytesRmndr-4                  2000000               792 ns/op              64 B/op          2 allocs/op
//BenchmarkRandStringBytesMask-4                   2000000               901 ns/op              64 B/op          2 allocs/op
//BenchmarkRandStringBytes-4                       2000000               950 ns/op              64 B/op          2 allocs/op
//BenchmarkRandStringRunes-4                       1000000              1237 ns/op             176 B/op          2 allocs/op
