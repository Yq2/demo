package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"
)

func main() {

	// 成功的读取到头不会返回io.EOF, 而是返回数据和nil。
	conn, err := net.Dial("tcp", "rpcx.site:80")
	if err != nil {
		fmt.Println("dial error ", err)
		return
	}
	defer conn.Close()

	// 发送请求，HTTP1.0协议
	conn.Write([]byte("GET / HTTP/1.0\r\n\r\n"))
	//fmt.Fprintln(conn,"GET / HTTP/1.0\r\n\r\n")

	// 读取response
	data, err := ioutil.ReadAll(conn)
	if err != nil {
		if err != io.EOF {
			fmt.Println("read error ", err)
		}
		panic(err)
	}
	// 显示结果
	fmt.Println("response ", string(data))
	fmt.Println("total response size ", len(data))

}
