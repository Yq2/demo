package main

import (
	"fmt"
	"runtime"
)

var msg string
var done bool = false
var doneChan = make(chan bool)

func main() {
	//a()
	b()
}

func a() {
	runtime.GOMAXPROCS(1)
	go func() {
		msg = " hello world"
		done = true
	}()

	for {
		// 不同goroutine之间不满足一致性内存模型
		if done {
			println(msg)
			break
		}
	}
}

func b() {
	runtime.GOMAXPROCS(1)
	go func() {
		msg = " hello world"
		// 通过显示同步
		doneChan <- true
	}()

	<-doneChan
	fmt.Println(msg)
}
