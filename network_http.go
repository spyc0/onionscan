package main

import (
	"h12.me/socks"
)

func ScanHTTP(onion string) bool {
	conn, err := socks.DialSocks5("127.0.0.1:9050", onion)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}
