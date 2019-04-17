package main

import (
	"fmt"
	"math/rand"
)

func GenerateInt1() chan int {
	ch := make(chan int, 10)
	go func() {
		for {
			ch <- rand.Int()
		}
	}()
	return ch
}

func GenerateInt2() chan int {
	ch := make(chan int, 10)
	go func() {
		for {
			ch <- rand.Int()
		}
	}()
	return ch
}

// 多个goroutine增强型生成器
func GenerateIntGroup() chan int {
	ch := make(chan int, 20)
	go func() {
		for {
			select {
			case ch <- <-GenerateInt1():
			case ch <- <-GenerateInt2():
			}
		}
	}()
	return ch
}

func main() {
	ch := GenerateIntGroup()
	for i := 0; i < 100; i++ {
		fmt.Println("<-ch =", <-ch)
	}
}
