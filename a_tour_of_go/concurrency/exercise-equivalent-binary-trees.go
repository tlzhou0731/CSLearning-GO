package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	} else {
		Walk(t.Left, ch)
		ch <- t.Value
		Walk(t.Right, ch)
	}
}
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := 0; i < 10; i++ {
		x, y := <-ch1, <-ch2
		if x != y {
			return false
		}
	}
	return true
}
func main() {
	tree1, tree2 := tree.New(2), tree.New(2)
	fmt.Println(Same(tree1, tree2))
	fmt.Println(tree1)
	fmt.Println(tree2)
}
