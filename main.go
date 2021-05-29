package main

import (
	"fmt"
	"net"
	"sort"
)

func worker(ports chan int, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("81.70.203.195:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		_ = conn.Close()
		results <- p
	}
}

func main() {
	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int
	var closeports []int

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		// 分配工作
		for i := 1; i < 1024; i++ {
			ports <- i
		}
	}()

	// 收集结果
	for i := 1; i < 1024; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		} else {
			closeports = append(closeports, port)
		}
	}

	close(ports)
	close(results)

	sort.Ints(openports)
	sort.Ints(closeports)

	fmt.Printf("打开的端口有: ")
	for _, port := range openports {
		fmt.Printf("%d ,", port)
	}
}
