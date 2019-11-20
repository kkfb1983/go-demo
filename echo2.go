package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(os.Args)
	s, sep := "", ""
	for idx, arg := range os.Args[1:] {
		s += sep + "参数" + strconv.Itoa(idx) + "->" + arg
		sep = "  "
	}
	fmt.Println(s)
}
