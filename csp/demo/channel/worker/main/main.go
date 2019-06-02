package main

import (
	"encoding/json"
	"fmt"
	"github.com/Yq2/demo/csp/demo/channel/worker"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	//MaxWorker = os.Getenv("MAX_WORKERS")
	//MaxQueue = os.Getenv("MAX_QUEUE")
	MaxWorker = 5
	MaxQueue  = 100
	MaxLength = 10
)

func main() {
	service := worker.NewService(MaxWorker, MaxQueue)
	service.Start()
	defer service.Stop()

	// 处理海量任务
	http.HandleFunc("/jobs", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			_, _ = w.Write([]byte("POST FORM方式添加Job"))
			return
		}
		// Job 以JSON格式提交
		var jobs []interface{}
		err := json.NewDecoder(io.LimitReader(r.Body, MaxLength)).Decode(jobs)
		if err != nil {
			w.Header().Set("Content-Type", "application/json;charset")
			w.WriteHeader(http.StatusBadRequest)
		}
		// 处理任务
		for job := range jobs {
			service.AddJob(job)
		}
		// OK
		w.WriteHeader(http.StatusOK)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}

func run(service *worker.Service) {
	for i := 0; i < 90000; i++ {
		service.AddJob(fmt.Sprintf("任务:%v", i))
	}
	time.Sleep(10 * time.Second)
	service.Stop()
}
