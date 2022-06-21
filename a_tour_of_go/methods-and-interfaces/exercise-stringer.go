package main

import "fmt"

type IPAddr [4]byte

func (ip_addr IPAddr) String() string {
	return fmt.Sprintf("%v,%v,%v,%v", ip_addr[0], ip_addr[1], ip_addr[2], ip_addr[3])
}

// TODO: 给 IPAddr 添加一个 "String() string" 方法

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
