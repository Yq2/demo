package main

import (
	"fmt"
	"log"
)

func main() {
	defer func() {
		// 无法捕获异常
		if r := MyRecover2(); r != nil {
			fmt.Println(r)
		}
	}()
	panic(1)
}

func MyRecover2() interface{} {
	log.Println("trace...")
	return recover()
}
