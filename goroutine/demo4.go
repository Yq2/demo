package main

import "fmt"

func main() {
	fmt.Println("********a1*******")
	a1()
	fmt.Println("********a2*******")
	a2()
	fmt.Println("********a3*******")
	a3()
}

func a1() {
	for i := 0; i < 5; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}

func a2() {
	for i := 0; i < 5; i++ {
		i := i
		defer func() {
			fmt.Println(i)
		}()
	}
}

func a3() {
	for i := 0; i < 5; i++ {
		defer func(i int) {
			fmt.Println(i)
		}(i)
	}
}
