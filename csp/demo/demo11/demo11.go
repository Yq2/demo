package main

import (
	"fmt"
	"time"
)

type query struct {
	sql    chan string
	result chan string
}

func exeQuery(q query) {
	go func() {
		sql := <-q.sql
		// 访问数据库

		// 输出结果通道
		q.result <- "result from " + sql
	}()
}

func main() {
	// 初始化 Query
	q := query{make(chan string, 1), make(chan string, 1)}

	// 执行query
	exeQuery(q)

	// 发送参数
	q.sql <- "select * from table"

	// 做其他事情
	time.Sleep(1 * time.Second)

	// 获取结果
	fmt.Println("<-q.result = ", <-q.result)
}
