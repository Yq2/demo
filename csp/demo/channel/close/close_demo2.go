package main

import (
	"fmt"
	"sync"
	"time"
)

func workerWG(wg *sync.WaitGroup, cancel chan bool) {
	defer wg.Done()
	for {
		select {
		default:
			fmt.Println("working...")
			// 正常工作
		case <-cancel:
			fmt.Println("exit")
			return
			// 退出
		}
	}
}

func main() {
	cancel := make(chan bool)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go workerWG(&wg, cancel)
	}
	time.Sleep(time.Second)
	close(cancel)
	wg.Wait()
}
