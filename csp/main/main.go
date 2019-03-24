package main

func main() {

	origin, wait := make(chan int, 100), make(chan struct{})
	Processor(origin, wait)
	for num := 2; num < 100000; num++ {
		origin <- num
	}
	close(origin)
	<-wait

}

// 求素数
func Processor(seq chan int, wait chan struct{}) {
	go func() {
		prime, ok := <-seq
		if !ok {
			close(wait)
			return
		}
		println("prime:%d", prime)
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
