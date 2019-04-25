package main

import (
	"fmt"
	"math/rand"
)

// done 接收退出信号
func GenerateIntD(done chan struct{}) chan int {
	ch := make(chan int, 5)
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

// done 接收退出信号
func GenerateIntE(done chan struct{}) chan int {
	ch := make(chan int, 10)
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

// 通过select 执行扇入Fan in 操作
func GenerateIntFanIn(done chan struct{}) chan int {
	ch := make(chan int, 10)
	send := make(chan struct{})
	go func() {
	Lable:
		for {
			select {
			case ch <- <-GenerateIntD(send):
			case ch <- <-GenerateIntE(send):
			case <-done:
				send <- struct{}{}
				send <- struct{}{}
				break Lable
			}
		}
		close(ch)
	}()
	return ch
}

func main() {
	// 创建一个作为接收退出信号的chan
	done := make(chan struct{})

	// 启动生成器
	ch := GenerateIntFanIn(done)

	// 获取生成器资源
	for i := 0; i < 50; i++ {
		fmt.Println("<-ch=", <-ch)
	}

	// 通知生产者停止生产
	done <- struct{}{}
	fmt.Println("stop generate...")
}
