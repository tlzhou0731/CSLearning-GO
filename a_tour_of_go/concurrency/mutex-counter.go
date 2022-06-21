package main

import (
	"fmt"
	_ "fmt"
	"sync"
	_ "sync"
	"time"
	_ "time"
)

type SafeCounter struct {
	v     map[string]int
	mutex sync.Mutex
}

func (c SafeCounter) Increase(key string) {
	c.mutex.Lock()
	c.v[key]++
	c.mutex.Unlock()
}
func (c SafeCounter) Value(key string) int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 100; i++ {
		go c.Increase("someKey")
		fmt.Println(c.Value("someKey"))
	}
	time.Sleep(time.Second)
	fmt.Println(c.Value("someKey"))
}
