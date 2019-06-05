package main

import "fmt"

type Inter interface {
	Ping()
	Pang()
}

type St struct {
	Name string
}

func (St) Ping() {
	println("ping")
}

func (*St) Pang() {
	println("pang")
}

func main() {

	var st *St = nil
	var it Inter = st
	fmt.Printf("%p\n", st)
	fmt.Printf("%p\n", it)
	if it != nil {
		it.Pang()

		// 下面语句会导致panic
		// 方法转换函数调用，第一个参数是St类型，由于*St是nil，无法获取指针所指的对象值，所以导致panic
		it.Ping()
	}
}
