package worker

type WorkerPool struct {
	workers []*Worker
	pool    chan *Worker
}

// 构建工作池
func NewWorkerPool(maxWorkers int) *WorkerPool {
	p := &WorkerPool{
		workers: make([]*Worker, maxWorkers),
		pool:    make(chan *Worker, maxWorkers),
	}
	// 初始化工作者
	for i, _ := range p.workers {
		worker := NewWorker(0)
		p.workers[i] = worker
		p.pool <- worker
	}
	return p
}

// 启动工作者
func (p *WorkerPool) Start() {
	for _, worker := range p.workers {
		worker.Start()
	}
}

// 停止工作者
func (p *WorkerPool) Stop() {
	for _, worker := range p.workers {
		worker.Stop()
	}
}

// 获取工作者 阻塞
func (p *WorkerPool) Get() *Worker {
	return <-p.pool
}

func (p *WorkerPool) Put(w *Worker) {
	p.pool <- w
}
