package main

import "time"

var limit = make(chan int, 3)
var work = [50]int{}

func main() {
	for range work {
		go func() {
			limit <- 1
			//...
			time.Sleep(2)
			<-limit
		}()
	}
	// 阻塞main
	// all goroutines are asleep - deadlock!
	select {}
}
