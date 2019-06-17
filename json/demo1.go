package main

import (
	"encoding/json"
	"fmt"
)

type UploadSomething struct {
	Type   string
	Object interface{}
}

type File struct {
	FileName string
}

type Png struct {
	Wide   float64 `json:"wide"`
	Height float64 `json:"height"`
}

func main() {

	var (
		input  string
		object json.RawMessage
		err    error
	)
	input = `{"type":"Png","object":{"wide":3.15,"height":5.13}}`
	ss := UploadSomething{
		Object: &object,
	}
	if err = json.Unmarshal([]byte(input), &ss); err != nil {
		panic(err)
	}
	fmt.Printf("ss.Object: %v\n", ss.Object)

	switch ss.Type {
	case "File":
		var f File
		if err := json.Unmarshal(object, &f); err != nil {
			panic(err)
		}
		fmt.Printf("filename: %v", f.FileName)
	case "Png":
		var p Png
		if err := json.Unmarshal(object, &p); err != nil {
			panic(err)
		}
		fmt.Printf("wide: %v", p.Wide)
	}
}
