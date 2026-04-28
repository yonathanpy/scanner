package main

import (
	"net"
)

func ResolveTarget(target string) string {
	ips, err := net.LookupIP(target)
	if err != nil || len(ips) == 0 {
		return target
	}
	return ips[0].String()
}
