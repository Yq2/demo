package singleton

import (
	"sync"
	"sync/atomic"
)

type singleton struct {
}

var (
	instance    *singleton
	initialized uint32
	mu          sync.Mutex
)

// 单例模式
func Instance() *singleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}
	mu.Lock()
	defer mu.Unlock()
	if instance == nil {
		defer atomic.StoreUint32(&initialized, 1)
		instance = &singleton{}
	}
	return instance
}

// 看看标准库里面的Once实现
type Once struct {
	m    sync.Mutex
	done uint32
}

func (o *Once) Do(f func()) {
	if atomic.LoadUint32(&o.done) == 1 {
		return
	}
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}

// 下面是基于sync.Once重新实现单利模式

var (
	instanceOnce *singleton
	once         sync.Once
)

func InstanceOfsyncOnce() *singleton {
	once.Do(func() {
		instanceOnce = &singleton{}
	})
	return instanceOnce
}
