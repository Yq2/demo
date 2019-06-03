package main

func main() {
	// 无法捕获异常
	defer recover()
	panic(1)
}
