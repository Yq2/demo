package main

import (
	"context"
	"fmt"
	"time"
)

type otherContext struct {
	context.Context
}

func work(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s get msg to cancel\n", name)
			return
		default:
			fmt.Printf("%s is running \n", name)
			time.Sleep(1 * time.Second)
		}
	}
}

func workWithValue(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s get msg to cancel\n", name)
			return
		default:
			value := ctx.Value("key").(string)
			fmt.Printf("%s is running value=%s \n", name, value)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {

	// 使用context.Background() 构建一个WithCancel类型的上下文
	ctxa, cancel := context.WithCancel(context.Background())

	// work 模拟运行并检测前端的退出通知
	go work(ctxa, "work1")

	// 使用WithDeadline 包装前面的上下文ctxa
	tm := time.Now().Add(3 * time.Second)
	ctxb, _ := context.WithDeadline(ctxa, tm)

	go work(ctxb, "work2")

	// 使用WithValue包装前面的上下文ctxb
	oc := otherContext{ctxb}
	ctxc := context.WithValue(oc, "key", "andes,pass from main")
	go workWithValue(ctxc, "work3")

	// 故意sleep 10秒，让work2，work3超时退出
	time.Sleep(10 * time.Second)

	// 显示调用work1 的cancel方法通知其退出
	fmt.Println("read cacel")
	cancel()
	fmt.Println("caceled")

	// 等待work1 打印退出通知
	time.Sleep(5 * time.Second)
	fmt.Println("Main Stop")
	// 退出顺序 work3 -> work2 -> work1
}
