package main

import (
	"encoding/json"
	"fmt"
)

type Animal struct {
	Name  string
	Order string
}

func main() {
	var jsonBlob = []byte(`[
        {"Name": "Platypus", "Order": "Monotremata"},
        {"Name": "Quoll",    "Order": "Dasyuromorphia"}
    ]`)

	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", animals)

	var f interface{}
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	json.Unmarshal(b, &f)
	for k, v := range f.(map[string]interface{}) {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int ", vv)
		case float64:
			fmt.Println(k, "is float64 ", vv)
		case []interface{}:
			fmt.Println(k, "is array:")
			for i, j := range vv {
				fmt.Println(i, j)
			}
		}
	}
}
