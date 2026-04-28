package main

import (
	"fmt"
	"net"
	"time"
)

type ScanResult struct {
	Port   int
	Open   bool
	Banner string
}

func ScanPort(target string, port int, timeout time.Duration) ScanResult {
	address := fmt.Sprintf("%s:%d", target, port)
	conn, err := net.DialTimeout("tcp", address, timeout)

	if err != nil {
		return ScanResult{Port: port, Open: false}
	}

	defer conn.Close()

	conn.SetReadDeadline(time.Now().Add(1 * time.Second))
	buffer := make([]byte, 1024)
	n, _ := conn.Read(buffer)

	banner := ""
	if n > 0 {
		banner = string(buffer[:n])
	}

	return ScanResult{
		Port:   port,
		Open:   true,
		Banner: banner,
	}
}
