package main

import (
	"fmt"
	"runtime"
)

func main() {
	c := make(chan struct{})
	ci := make(chan int, 100)
	go func(i chan struct{}, j chan int) {
		for i := 0; i < 10; i++ {
			j <- i
		}
		close(j)
		i <- struct{}{}
	}(c, ci)
	fmt.Println("NumGoroutine=", runtime.NumGoroutine())
	<-c
	fmt.Println("NumGoroutine=", runtime.NumGoroutine())
	for v := range ci {
		fmt.Println("v=", v)
	}

}
