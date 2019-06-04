package hash

import (
	"testing"
)

func BenchmarkMD5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		md5Hash()
	}
}

func BenchmarkSHA1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sha1Hash()
	}
}

func BenchmarkMurmurHash32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		murmur32()
	}
}

func BenchmarkMurmurHash64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		murmur64()
	}
}

//BenchmarkMD5-8                  10000000               133 ns/op
//BenchmarkSHA1-8                 10000000               151 ns/op
//BenchmarkMurmurHash32-8         100000000               18.7 ns/op
//BenchmarkMurmurHash64-8         50000000                26.1 ns/op
