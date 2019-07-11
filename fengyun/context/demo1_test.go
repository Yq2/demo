package main

import (
	"net/http"
	"testing"
)

func BenchmarkContext(b *testing.B) {
	for j := 0; j < 500; j++ {
		for i := 0; i < b.N; i++ {
			_, _ = http.Get("http://localhost:8090/")
		}
	}
}
