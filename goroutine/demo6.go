package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

func main() {
	// 谨慎内存地址变化
	var x int = 42
	var p uintptr = uintptr(unsafe.Pointer(&x))

	runtime.GC()
	var px *int = (*int)(unsafe.Pointer(p))
	fmt.Println("px= ", *px)
}
