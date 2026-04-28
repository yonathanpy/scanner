package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	cfg := LoadConfig()

	fmt.Println("[*] scanner starting")
	fmt.Printf("[*] Target: %s\n", cfg.Target)

	jobs := make(chan int, 1000)
	results := make(chan ScanResult, 1000)

	var wg sync.WaitGroup

	for i := 0; i < cfg.Workers; i++ {
		wg.Add(1)
		go Worker(i, cfg, jobs, results, &wg)
	}

	go func() {
		for p := cfg.StartPort; p <= cfg.EndPort; p++ {
			jobs <- p
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	start := time.Now()

	for r := range results {
		if r.Open {
			fmt.Printf("[OPEN] %d (%s)\n", r.Port, r.Banner)
		}
	}

	fmt.Println("[*] Done in", time.Since(start))
}
