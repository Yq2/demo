package main

import (
	"fmt"
	"sync"
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
	// 创建任务通道
	taskPool := make(chan task, 10)
	// 创建结果通道
	resultPool := make(chan int, 10)

	// wait用于同步等待任务的执行
	wait := &sync.WaitGroup{}

	// 初始化task的goroutine，计算100个自然数之和
	go InitTask(taskPool, resultPool, 100)

	// 每个task 启动一个go处理
	go DistributeTask(taskPool, wait, resultPool)

	// 通过结果通道获取结果并汇总
	sum := ProcessResult(resultPool)
	fmt.Println("sum =", sum)

}

// 构建task并写入task通道
func InitTask(resultPool chan<- task, r chan int, p int) {
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
		resultPool <- tsk
	}
	if mod != 0 {
		tsk := task{
			begin:  high + 1,
			end:    p,
			result: r,
		}
		resultPool <- tsk
	}
	close(resultPool)
}

// 读取task chan ，每个task 启动一个worker goroutine 进行处理
// 等待每个task运行完，关闭结果通道
func DistributeTask(resultPool <-chan task, wait *sync.WaitGroup, result chan int) {
	for v := range resultPool {
		wait.Add(1)
		go ProcessTask(v, wait)
	}
	wait.Wait()
	close(result)
}

// goutine 处理具体工作，并将处理结果发送给结果通道
func ProcessTask(t task, wait *sync.WaitGroup) {
	t.do()
	wait.Done()
}

// 读取结果通道，汇总结果
func ProcessResult(resultPool chan int) int {
	sum := 0
	for r := range resultPool {
		sum += r
	}
	return sum
}
