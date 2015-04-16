package main

import (
	"fmt"
	"strings"
)

const (
	userAgent = "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"
)

func ScanHTTP(onion string) bool {
	conn, err := Connect(onion + ":80")
	if err != nil {
		return false
	}
	defer conn.Close()
	request := "GET / HTTP/1.1\r\n" +
	fmt.Sprintf("Host: %s:80\r\n", onion) +
	"Accept: text/html,application/xhtml+xml,application/xml\r\n" +
	"Accept-Language: en-US,en\r\n" +
	fmt.Sprintf("User-Agent: %s\r\n\r\n", userAgent)
	err = conn.Write(request)
	if err != nil {
		return false
	}
	rdata, err := conn.Read(2048)
	if err != nil {
		return false
	}
	if strings.Split(strings.Split(rdata, "\r\n")[0], " ")[0] == "HTTP/1.1" {
		return true
	} else {
		return false
	}
}
