package main

import (
	"log"
	"os"
)

func main() {
	demo()
}

// defer 函数在退出时才能执行，在for循环执行defer会导致资源延迟释放
func demo() {
	for i := 0; i < 5; i++ {
		f, err := os.Open("demo1.go")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
	}
}

// 解决办法是在for循环中构造一个局部函数，在局部函数中执行defer
func demo2() {
	for i := 0; i < 5; i++ {
		func() {
			f, err := os.Open("demo1.go")
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
		}()
	}
}
