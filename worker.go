package main

import (
	"sync"
)

func Worker(id int, cfg Config, jobs <-chan int, results chan<- ScanResult, wg *sync.WaitGroup) {
	defer wg.Done()

	for port := range jobs {
		res := ScanPort(cfg.Target, port, cfg.Timeout)
		results <- res
	}
}
