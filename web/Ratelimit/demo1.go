package main

import (
	"fmt"
	"time"
)

func main() {
	var fillInterval = time.Millisecond * 50
	var capacity = 100
	var tokenBucket = make(chan struct{}, capacity)
	fillToken := func() {
		ticker := time.NewTicker(fillInterval)
		for {
			select {
			case <-ticker.C:
				select {
				case tokenBucket <- struct{}{}:
				default:

				}
				fmt.Println("current token cnt: ", len(tokenBucket), "time:", time.Now().String())
			}
		}
	}
	takeAvailable := func(block bool) {
		if block {
			for {
				select {
				case <-tokenBucket:
					fmt.Println("get one")
				}
				time.Sleep(time.Millisecond * 50)
			}
		} else {
			for {
				select {
				case <-tokenBucket:
					fmt.Println("get one")
				default:
					fmt.Println("blocking")
				}
				time.Sleep(time.Millisecond * 50)
			}
		}
	}

	go fillToken()
	go takeAvailable(false)

	time.Sleep(time.Hour)
}
