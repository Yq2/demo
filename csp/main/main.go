package main

import "fmt"

var results []int

const END_NUMBER = 90000

func main() {
	origin, wait := make(chan int, 100), make(chan struct{})
	Processor(origin, wait)
	for num := 2; num < END_NUMBER; num++ {
		origin <- num
	}
	close(origin)
	<-wait
	fmt.Printf("%d 以内的素数有: %d 个\n", END_NUMBER, len(results))
	fmt.Println(results)
}

// 求素数
// TODO 算法精髓：从2开始每找到一个素数就标记所有能被该素数整除的所有数，直到没有可标记的数，剩下的就是素数
func Processor(seq chan int, wait chan struct{}) {
	go func() {
		prime, ok := <-seq
		if !ok {
			close(wait)
			return
		}
		results = append(results, prime)
		out := make(chan int, 100)
		Processor(out, wait)
		for num := range seq {
			if num%prime != 0 {
				out <- num
			}
		}
		close(out)
	}()
}
