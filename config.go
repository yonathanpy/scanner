package main

import (
	"flag"
	"time"
)

type Config struct {
	Target    string
	StartPort int
	EndPort   int
	Workers   int
	Timeout   time.Duration
}

func LoadConfig() Config {
	target := flag.String("target", "127.0.0.1", "target host")
	start := flag.Int("start", 1, "start port")
	end := flag.Int("end", 1024, "end port")
	workers := flag.Int("workers", 100, "number of workers")
	timeout := flag.Int("timeout", 500, "timeout ms")

	flag.Parse()

	return Config{
		Target:    ResolveTarget(*target),
		StartPort: *start,
		EndPort:   *end,
		Workers:   *workers,
		Timeout:   time.Duration(*timeout) * time.Millisecond,
	}
}
