package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
)

func main() {

	var a = [...]int{1, 2, 3}
	var b = &a
	for i, v := range b {
		fmt.Println(i, v)
	}
	// for range
	for i := range a {
		fmt.Printf("b[%d]: %d\n", i, b[i])
	}
	for i, v := range b {
		fmt.Printf("b[%d]: %d\n", i, v)
	}
	for i := 0; i < len(a); i++ {
		fmt.Printf("b[%d]: %d\n", i, b[i])
	}

	var times [5][0]int
	for range times {
		fmt.Println("hello")
	}
	//
	var line1 [2]image.Point
	var line2 = [...]image.Point{
		image.Point{X: 0, Y: 0}, image.Point{1, 1},
	}
	//
	var decoder1 [2]func(reader io.Reader) (image.Image, error)
	var decoder2 = [...]func(reader io.Reader) (image.Image, error){
		png.Decode,
		jpeg.Decode,
	}
	var unknown1 [2]interface{}
	var unknown2 = [...]interface{}{123, "你好"}

	//
	var chanList = [2]chan int{}
	_ = line1
	_ = line2
	_ = decoder1
	_ = decoder2
	_ = unknown1
	_ = unknown2
	_ = chanList
}
