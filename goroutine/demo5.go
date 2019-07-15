package main

import (
	"errors"
	"fmt"
)

func returnsErr() error {
	var p error = nil
	p = errors.New("error demo")
	return p
}

func main() {

	fmt.Println(returnsErr())
}
