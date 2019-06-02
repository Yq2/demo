package main

import (
	"fmt"
	"time"
)

func worker(cancel chan bool) {
	for {
		select {
		default:
			fmt.Println("working...")
			// 正常工作
		case <-cancel:
			fmt.Println("exit")
			// 退出
		}
	}
}

func main() {
	cancel := make(chan bool)
	for i := 0; i < 10; i++ {
		go worker(cancel)
	}
	time.Sleep(time.Second)
	close(cancel)
}
