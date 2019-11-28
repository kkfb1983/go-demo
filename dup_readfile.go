package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stdin, "dup file:%v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
		//fmt.Println(counts)
		//return
	}
	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
