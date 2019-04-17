package main

import (
	"fmt"
	"math/rand"
)

// 简单生成器
func GenerateIntA() chan int {
	ch := make(chan int, 10)
	go func() {
		for {
			ch <- rand.Int()
		}
	}()
	return ch
}

func main() {
	ch := GenerateIntA()
	fmt.Println("<-ch=", <-ch)
	fmt.Println("<-ch=", <-ch)
}
