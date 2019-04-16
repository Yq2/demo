package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var urls = []string{
	"http://www.baidu.com",
	"http://www.qq.com",
	"https://developer.android.google.cn/samples",
}

func main() {
	for {
		if len(urls) > 500 {
			break
		}
		urls = append(urls, urls...)
	}
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			resp, err := http.Get(url)
			if err != nil {
				fmt.Printf("url request err: %+v\n", err)
				return
			}
			fmt.Println("url=", url, " status=", resp.Status)
		}(url)
	}
	wg.Wait()
}
