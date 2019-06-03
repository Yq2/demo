package main

import (
	"errors"
	"fmt"
	"runtime"
)

func foo() (err error) {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = fmt.Errorf("Unknown panic: %v", r)
			}
		}
	}()
	panic("TODO")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			switch r.(type) {
			case runtime.Error:
				// 这是运行时错误类型异常
			case error:
				// 普通错误类型异常
			default:
				// 其他类型异常
			}
		}
	}()

}
