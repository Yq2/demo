package main

import (
	"golang.org/x/tools/godoc/vfs"
	"golang.org/x/tools/godoc/vfs/gatefs"
)

func main() {
	fs := gatefs.New(vfs.OS("/path"), make(chan bool, 8))
	fs.String()
}

// 利用缓冲通道不仅可以控制最大并发数目，
// 而且可以通道带缓存的channel的利用量和最大容量来判断程序运行的并发率
// 当管道为空时可以任务时空闲状态
// 当管道满了时任务时繁忙状态，这对于后台一些低级任务的运行时有参考价值的
