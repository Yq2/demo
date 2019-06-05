package main

import "fmt"

func main() {

	channel := func() <-chan int {
		ch := make(chan int)
		go func() {
			for i := 0; ; i++ {
				fmt.Println("<- i=", i)
				ch <- i
			}
		}()
		return ch
	}()
	for v := range channel {
		fmt.Println(v)
		if v == 5 {
			// 主goroutine退出之后，channel所在goroutine无法自动退出
			break
		}
	}

}
