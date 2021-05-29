package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	var wg sync.WaitGroup
	for i := 21; i < 65535; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			address := fmt.Sprintf("81.70.203.195:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				fmt.Printf("端口: %s 已关闭\n", address)
				return
			}
			_ = conn.Close()
			fmt.Printf("端口: %s 已打开\n", address)
		}(i)
	}
	wg.Wait()

	elapsed := time.Since(start) / 1e9
	fmt.Printf("\n\n%d seconds", elapsed)
}
