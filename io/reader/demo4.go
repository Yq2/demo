package main

import (
	"fmt"
	"io"
	"net"
	"strings"
)

func main() {
	// 读取到0字节，并且遇到EOF, 返回EOF
	// 读取到0<n<len(buf)字节，并且遇到EOF, 返回ErrUnexpectedEOF
	// 读取到n==len(buf),即使遇到error也返回err=nil
	// 返回err != nil则肯定n<len(buf)

	conn, err := net.Dial("tcp", "rpcx.site:80")
	if err != nil {
		fmt.Println("dial error ", err)
		return
	}
	defer conn.Close()

	// 发送请求 http 1.0协议
	fmt.Fprintln(conn, "GET / HTTP/1.0\r\n\r\n")

	// 读取response
	var sb strings.Builder
	buf := make([]byte, 256)
	for {
		n, err := io.ReadFull(conn, buf)
		if err != nil {
			if err != io.EOF && err != io.ErrUnexpectedEOF {
				fmt.Println("read error ", err)
			}
			break
		}
		sb.Write(buf[:n])
	}
	fmt.Println("response ", sb.String())
	fmt.Println("total response size ", sb.Len())
}
