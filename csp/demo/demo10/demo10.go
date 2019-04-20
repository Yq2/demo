package main

import (
	"fmt"
)

// 工作池的goutine数目
const (
	NUMBER = 10
)

// 多个task任务
type task struct {
	begin  int
	end    int
	result chan<- int
}

// 任务执行，计算begin 到end的和
func (t *task) do() {
	sum := 0
	for i := t.begin; i < t.end; i++ {
		sum += i
	}
	t.result <- sum
}

func main() {
	// 固定worker工作池
	workers := NUMBER

	// 创建任务通道
	taskchan := make(chan task, 10)

	// 创建结果通道
	resultchan := make(chan int, 10)

	// worker信号通道
	// 主要用来同步控制goroutine
	done := make(chan struct{}, 10)

	// 初始化task的goroutine，计算100个自然数之和
	// 该方法是阻塞的
	go InitTask(taskchan, resultchan, 100)

	// 分发任务到NUMBERS个goroutine池
	// 推荐在主goroutine里面开启多个工作goroutine
	DistributeTask(taskchan, workers, done)

	// 获取各个goroutine处理完成任务的通知，并关闭结果通道
	// CloseResult是阻塞的，需要等所有go完全退出后才算执行完
	go CloseResult(done, resultchan, workers)

	// 通过结果通道获取结果并汇总
	sum := ProcessResult(resultchan)
	fmt.Println("sum =", sum)

}

// 构建task并写入task通道
func InitTask(taskchan chan<- task, r chan int, p int) {
	qu := p / 10
	mod := p % 10
	high := qu * 10
	for j := 0; j < qu; j++ {
		b := 10*j + 1
		e := 10 * (j + 1)
		tsk := task{
			begin:  b,
			end:    e,
			result: r,
		}
		taskchan <- tsk
	}
	if mod != 0 {
		tsk := task{
			begin:  high + 1,
			end:    p,
			result: r,
		}
		taskchan <- tsk
	}
	close(taskchan)
}

// 读取task chan 并分发到workers goroutine处理，总的数量是workers
func DistributeTask(taskchan <-chan task, workers int, done chan struct{}) {
	for i := 0; i < workers; i++ {
		go ProcessTask(taskchan, done)
	}
}

// 工作goroutine 处理具体工作，并将处理结果发送到结果chan
func ProcessTask(taskchan <-chan task, done chan struct{}) {
	for t := range taskchan {
		t.do()
	}
	done <- struct{}{}
}

func CloseResult(done chan struct{}, resultchan chan int, workers int) {
	for i := 0; i < workers; i++ {
		<-done
	}
	close(done)
	close(resultchan)
}

// 读取结果通道，汇总结果
func ProcessResult(resultchan chan int) int {
	sum := 0
	for r := range resultchan {
		sum += r
	}
	return sum
}
