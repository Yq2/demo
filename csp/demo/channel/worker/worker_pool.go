package worker

type Pool struct {
	workers []*Worker
	pool    chan *Worker
}

// 构建工作者池
func NewWorkerPool(maxWorkers int) *Pool {
	p := &Pool{
		workers: make([]*Worker, maxWorkers),
		pool:    make(chan *Worker, maxWorkers),
	}
	// 初始化工作者
	for i, _ := range p.workers {
		worker := NewWorker(maxWorkers)
		p.workers[i] = worker
		p.pool <- worker
	}
	return p
}

// 启动工作者
func (p *Pool) Start() {
	for _, worker := range p.workers {
		worker.Start()
	}
}

// 停止工作者
func (p *Pool) Stop() {
	for _, worker := range p.workers {
		worker.Stop()
	}
}

// 获取工作者 阻塞
func (p *Pool) Get() *Worker {
	return <-p.pool
}

func (p *Pool) Put(w *Worker) {
	p.pool <- w
}
