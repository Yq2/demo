package main

import (
	"fmt"
	"io"
	"net"
	"strings"
)

func main() {
	// 返回一个内容长度受限的Reader，当从这个Reader中读取了n个字节一会就会遇到EOF
	conn, err := net.Dial("tcp", "rpcx.site:80")
	if err != nil {
		fmt.Println("dial error", err)
		return
	}
	defer conn.Close()

	// 发送请求 HTTP1.0协议
	fmt.Fprintln(conn, "GET / HTTP/1.0\r\n\r\n")
	var sb strings.Builder
	buf := make([]byte, 256)
	rr := io.LimitReader(conn, 10000)
	for {
		n, err := io.ReadAtLeast(rr, buf, 256)
		if err != nil {
			if err != io.EOF && err != io.ErrUnexpectedEOF {
				fmt.Println("read error", err)
			}
			break
		}
		sb.Write(buf[:n])
	}
	fmt.Println("response ", sb.String())
	fmt.Println("total response size ", sb.Len())
}
