package main

import "sync"

var mutex sync.Mutex

func main() {
	Map := make(map[int]int)
	for i := 0; i < 100000; i++ {
		i := i
		go writeMap(Map, i, i)
		go readMap(Map, i)
	}
}

func readMap(Map map[int]int, key int) int {
	mutex.Lock()
	defer mutex.Unlock()
	return Map[key]
}

func writeMap(Map map[int]int, key int, value int) {
	mutex.Lock()
	defer mutex.Unlock()
	Map[key] = value
}
