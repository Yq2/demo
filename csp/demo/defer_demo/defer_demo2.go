package main

import "fmt"

func f4() int {
	r := 0
	defer func() {
		r++
	}()

	return r
}

func f5() int {
	r := 0
	defer func(i int) {
		i++
	}(r)

	return 0
}

func main() {
	fmt.Println("f4=", f4())
	fmt.Println("f5=", f5())
}
