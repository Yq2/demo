package main

import (
	"fmt"
	"io"
	"net"
	"strings"
)

func main() {

	// 从输入流中读取至少min个字节到buf中，直到buf读取满或者遇到error包括EOF
	// 如果n < min, 肯定返回错误
	// 没有数据可读返回n=0和EOF
	// n < min,并且遇到EOF,返回n和ErrUnexpectedEOF
	// 如果参数min>len(buf),返回0,ErrShortBuffer
	// 如果读取了至少min个字节，即使遇到错误也会返回err=nil

	conn, err := net.Dial("tcp", "rpcx.site:80")
	if err != nil {
		fmt.Println("dial error ", err)
		return
	}
	defer conn.Close()

	fmt.Fprintln(conn, "GET / HTTP/1.0\r\n\r\n")
	var sb strings.Builder
	buf := make([]byte, 256)
	for {
		n, err := io.ReadAtLeast(conn, buf, 256)
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
