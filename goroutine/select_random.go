package main

import "fmt"

func random(n int) <-chan string {
	c := make(chan string)
	go func() {
		defer close(c)
		// 利用select随机选择通道的特性，在特定序列中产生随机数
		for i := 0; i < n; i++ {
			select {
			case c <- "选a":
			case c <- "选b":
			case c <- "选c":
			case c <- "选d":
			}
		}
	}()
	return c
}

func main() {

	for i := range random(10) {
		fmt.Println("i=", i)
	}
}
