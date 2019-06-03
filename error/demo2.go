package main

func MyRecover() interface{} {
	return recover()
}

func main() {
	// 可以正常捕获异常
	defer MyRecover()
	panic(1)
}
