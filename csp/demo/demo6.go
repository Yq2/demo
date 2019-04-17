package main

import (
	"fmt"
	"math/rand"
)

// 带退出通知机制的生成器
func GenerateIntC(done chan struct{}) chan int {
	ch := make(chan int)
	go func() {
	Lable:
		for {
			select {
			case ch <- rand.Int():
			case <-done:
				break Lable
			}
		}
		close(ch)
	}()
	return ch
}

func main() {
	done := make(chan struct{})
	ch := GenerateIntC(done)
	fmt.Println("<-ch=", <-ch)
	fmt.Println("<-ch=", <-ch)
	close(done)
	for v := range ch {
		fmt.Println("v=", v)
	}

}
