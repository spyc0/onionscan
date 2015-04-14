package main

import (
	"github.com/samuel/go-socks/socks"
)

func ScanHTTP(onion string) bool {
	proxy := &socks.Proxy{Addr: "127.0.0.1:9050"}
	conn, err := proxy.Dial("tcp", onion + ":80")
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}
