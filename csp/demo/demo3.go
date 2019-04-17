package main

import (
	"fmt"
	"math/rand"
	"runtime"
)

// 退出通知机制
func GenerateInt(done chan struct{}) chan int {
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
	ch := GenerateInt(done)
	fmt.Println("ch: ", <-ch)
	fmt.Println("ch: ", <-ch)
	close(done)
	fmt.Println("ch: ", <-ch)
	fmt.Println("ch: ", <-ch)
	fmt.Println("NumGoroutine=", runtime.NumGoroutine())

}
