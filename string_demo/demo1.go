package main

// 字符串的结构化定义
type StringHeader struct {
	Data uintptr
	Len  int
}
