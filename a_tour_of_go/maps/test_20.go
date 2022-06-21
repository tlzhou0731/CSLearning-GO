package main

import (
	"fmt"
	_ "fmt"
)

type Vertex20 struct {
	Lat, Long float64
}

var m20 = map[string]Vertex20{
	"Bell Labs": {88.88, 99.99},
	"Google":    {77.77, 22.22},
}

func main() {
	fmt.Println(m20)
}
