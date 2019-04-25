package main

import "fmt"

// 命名返回值
func f1() (r int) {
	defer func() {
		r++
	}()
	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func main() {

	// 对于defer的函数返回整体上有三个步骤
	// 1 指定return 的值拷贝，将return语句返回的值复制到函数返回值栈区（如果只有一个return，不带任何变量或值，则此步骤不做任何动作）
	// 2 执行defer语句，多个defer按照FILO先进后出规则顺序执行
	// 3 执行调整RET指令
	fmt.Println("f1=", f1())
	fmt.Println("f2=", f2())
	fmt.Println("f3=", f3())
}
