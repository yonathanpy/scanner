package main

import (
	"fmt"
)

func PrintResult(r ScanResult) {
	if r.Open {
		fmt.Printf("[+] Port %d open (%s)\n", r.Port, r.Banner)
	}
}
