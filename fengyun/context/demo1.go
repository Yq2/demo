package main

import (
	"fmt"
	"net/http"
)

func painc(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%+v\n", r.Context())
}

func main() {
	http.HandleFunc("/", painc)
	if err := http.ListenAndServe(":8090", nil); err != nil {
		fmt.Println(err)
	}
}
