package main

import (
	"fmt"
	"net"
)

func main() {

	for i := 21; i < 30; i++ {
		address := fmt.Sprintf("81.70.203.195:%d", i)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			fmt.Printf("端口: %s 已关闭\n", address)
			continue
		}
		_ = conn.Close()
		fmt.Printf("端口: %s 已打开\n", address)
	}
}
