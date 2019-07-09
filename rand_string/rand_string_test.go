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

//BenchmarkRandStringBytes-4                       2000000               968 ns/op
//BenchmarkRandStringRunes-4                       1000000              1310 ns/op
//BenchmarkRandStringBytesRmndr-4                  2000000               843 ns/op
//BenchmarkRandStringBytesMask-4                   2000000               911 ns/op
//BenchmarkRandStringBytesMaskImpr-4               5000000               249 ns/op
//BenchmarkRandStringBytesMaskImprSrc-4           10000000               212 ns/op
