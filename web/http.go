package main

import (
	"net/http"
	"sync"
)

type ServeMux struct {
	mu      sync.RWMutex        // 锁，由于请求涉及并发处理，因此需要一个锁
	m       map[string]muxEntry // 路由规则，一个string对应一个mux实体，这里的string就是注册的路由表达式
	pattern string              // 是否在任意的规则中带host信息
}

type muxEntry struct {
	explicit bool         // 是否精确匹配
	h        http.Handler // 这个路由表达式对应哪个handler
	pattern  string       // 匹配字符串
}

type Handler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}
