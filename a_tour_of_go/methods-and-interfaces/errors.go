package main

import (
	"fmt"
	"strconv"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("<nil>")
	}

	i, err2 := strconv.Atoi("42")
	if err2 != nil {
		fmt.Printf("couldn't convert number: %v\n", err2)
		return
	}
	fmt.Println("Converted integer:", i)
}
