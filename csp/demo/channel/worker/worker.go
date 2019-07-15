package worker

import (
	"fmt"
	"sync"
)

type Worker struct {
	job  chan interface{}
	quit chan struct{}
	wg   sync.WaitGroup
}

// 构建工作者
func NewWorker(maxJobs int) *Worker {
	return &Worker{
		job:  make(chan interface{}, maxJobs),
		quit: make(chan struct{}),
	}
}

// 启动任务
func (w *Worker) Start() {
	w.wg.Add(1)
	go func() {
		defer w.wg.Done()
		for {
			// 接收任务
			// 此时工作中已经从工作池中取出
			select {
			case job := <-w.job:
				// 处理任务
				fmt.Println("[Worker] job:%v", job)
			case <-w.quit:
				return
			}
		}
	}()
}

//关闭任务
func (w *Worker) Stop() {
	w.quit <- struct{}{}
	w.wg.Wait()
}

func (w *Worker) AddJob(job interface{}) {
	w.job <- job
}
