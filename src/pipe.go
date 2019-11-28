package main

import "fmt"

func test_pipe() {
	pipe := make(chan int, 3)
	pipe <- 4
	pipe <- 5
	pipe <- 6
	var t1, t2 int
	t1 = <-pipe
	t2 = <-pipe
	pipe <- 4
	fmt.Println(len(pipe), t1, t2)
}
