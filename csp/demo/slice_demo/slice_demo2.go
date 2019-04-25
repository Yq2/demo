package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	a := []int{0, 1, 2, 3, 4, 5, 6}
	b := a[0:4]

	// 使用反射SliceHeader 来获取切片运行时的数据结构
	as := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	bs := (*reflect.SliceHeader)(unsafe.Pointer(&b))

	// a,b 共享底层数组
	fmt.Printf("a=%v,len=%d,cap=%d,type=%d\n", a, len(a), cap(a), as.Data)
	fmt.Printf("b=%v,len=%d,cap=%d,type=%d\n", b, len(b), cap(b), bs.Data)
	b = append(b, 10, 11, 12)

	// a,b 继续共享底层数组，修改b会影响共享的底层数据，间接影响a
	fmt.Printf("a=%v,len=%d,cap=%d,type=%d\n", a, len(a), cap(a), as.Data)
	fmt.Printf("b=%v,len=%d,cap=%d,type=%d\n", b, len(b), cap(b), bs.Data)

	// len(b)=7 底层数据容量是7，此时需要重新分配数组，并将原数组拷贝到新数组
	b = append(b, 13, 14)

	as = (*reflect.SliceHeader)(unsafe.Pointer(&a))
	bs = (*reflect.SliceHeader)(unsafe.Pointer(&b))

	// 可以看到a和b执行的底层数据组的指针已经不同了
	fmt.Printf("a=%v,len=%d,cap=%d,type=%d\n", a, len(a), cap(a), as.Data)
	fmt.Printf("b=%v,len=%d,cap=%d,type=%d\n", b, len(b), cap(b), bs.Data)
}
