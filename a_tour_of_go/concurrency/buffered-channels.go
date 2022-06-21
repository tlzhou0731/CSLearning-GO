package main

import "fmt"

func main() {
	ch := make(chan int, 100)
	ch <- 1
	ch <- 2
	for val, states := <-ch; states != false; val, states = <-ch {
		fmt.Println(val, states)
	}

}
