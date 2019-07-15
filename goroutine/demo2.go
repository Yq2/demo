package main

import (
	"fmt"
	"runtime"
)

func main() {

	runtime.GOMAXPROCS(1)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}()

	select {
	// 或者通过阻塞的方式避免CPU占用
	// 但是这种方式因为没有可用的channel抛出错误 all goroutines are asleep - deadlock!
	}
}
