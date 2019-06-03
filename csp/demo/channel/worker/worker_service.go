package worker

import "sync"

type Service struct {
	workers *Pool
	jobs    chan interface{}
	maxJobs int
	wg      sync.WaitGroup
}

func NewService(maxWorkers, maxJobs int) *Service {
	return &Service{
		workers: NewWorkerPool(maxWorkers),
		jobs:    make(chan interface{}, maxJobs),
	}
}

func (p *Service) Start() {
	p.jobs = make(chan interface{}, p.maxJobs)

	p.wg.Add(1)
	p.workers.Start()

	go func() {
		defer p.wg.Done()
		for job := range p.jobs {
			go func(job interface{}) {
				// 从工作者池中取一个工作者
				worker := p.workers.Get()
				// 完成任务后返还给工作者池
				defer p.workers.Put(worker)
				// 提交任务处理（异步）
				worker.AddJob(job)
			}(job)
		}
	}()
}

func (p *Service) Stop() {
	p.workers.Stop()
	close(p.jobs)
	p.wg.Wait()
}

// 提交任务，任务管道带较大的缓存，延缓阻塞时间
func (p *Service) AddJob(job interface{}) {
	p.jobs <- job
}
