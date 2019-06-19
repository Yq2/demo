package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("价值一亿AI核心代码")
	r := bufio.NewReader(os.Stdin)
	for {
		c, err := r.ReadString('\n')
		if err != nil {
			continue
		}
		c = strings.Replace(c, "吗", "", -1)
		c = strings.Replace(c, "? ", "!", -1)
		c = strings.Replace(c, "?", "!", -1)
		c = strings.Replace(c, "？", "", -1)
		fmt.Println(c)
	}
}
