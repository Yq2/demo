package main

import (
	"fmt"
	"io"
	"net"
	"strings"
)

func main() {

	// io.Copy使用方法
	// 内部使用CopyBuffer，可能会用到一个指定的buf来复制数据
	// 在两种特殊情况下buf并不会被使用，而是使用接口的方法进行复制
	// src 实现了WriterTo接口
	// dst 实现了ReadFrom接口
	//if wt, ok := src.(WriterTo); ok {
	//	return wt.WriteTo(dst)
	//}
	//// Similarly, if the writer has a ReadFrom method, use it to do the copy.
	//if rt, ok := dst.(ReaderFrom); ok {
	//	return rt.ReadFrom(src)
	//}

	conn, err := net.Dial("tcp", "rpcx.site:80")
	if err != nil {
		fmt.Println("dial error ", err)
		return
	}
	defer conn.Close()

	_, _ = fmt.Fprintln(conn, "GET / HTTP/1.0\r\n\r\n")

	var sb strings.Builder
	_, err = io.Copy(&sb, conn)
	if err != nil {
		fmt.Println("read error", err)
		return
	}
	fmt.Println("response ", sb.String())
	fmt.Println("total response size ", sb.Len())
}
