package main

import (
	"golang.org/x/tools/godoc/vfs"
	"os"
	"time"
)

type gate chan bool

func (g gate) enter()            { g <- true }
func (g gate) leave()            { <-g }
func (g gate) Len() int          { return len(g) }
func (g gate) Cap() int          { return cap(g) }
func (g gate) Idle() bool        { return len(g) == 0 }
func (g gate) Busy() bool        { return len(g) == cap(g) }
func (g gate) Fraction() float64 { return float64(len(g)) / float64(cap(g)) }

type gatefs struct {
	fs vfs.FileSystem
	gate
}

func (fs gatefs) Lstat(p string) (os.FileInfo, error) {
	fs.enter()
	defer fs.leave()
	return fs.Lstat(p)
}

func New(fs vfs.FileSystem, gate chan bool) *gatefs {
	p := &gatefs{fs: fs, gate: gate}
	// 后台监控线程
	go func() {
		for {
			switch {
			case p.gate.Idle():
			// 处理后台任务
			case p.gate.Fraction() >= 0.7:
				// 并发报警
				defer time.Sleep(1 * time.Second)
			}
		}
	}()

	return p
}
