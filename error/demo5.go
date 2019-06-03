package main

func main() {
	defer func() {
		if r := recover(); r != nil {

		}
		// 虽然总是返回nil, 但是可以恢复异常状态
	}()
	// 警告: 用`nil`为参数抛出异常
	panic(nil)
}
