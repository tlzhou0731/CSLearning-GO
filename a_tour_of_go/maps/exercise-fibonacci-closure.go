package main

import "fmt"

// 返回一个“返回int的函数”
func fibonacci() func() int {
	pre := 0
	cur := 0
	cnt := 0
	return func() int {
		if cnt == 0 {
			cnt++
		} else if cnt == 1 {
			cur = 1
			cnt++
		} else {
			cur = pre + cur
			pre = cur - pre
			cnt++
		}
		return cur
	}

}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
