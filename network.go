package main

import (
	"net"
	"github.com/samuel/go-socks/socks"
)

type Network struct {
	Conn  *net.Conn
}

func Connect(host string) (*Network, error) {
	proxy := &socks.Proxy{Addr: "127.0.0.1:9050"}
	conn, err := proxy.Dial("tcp", host)
	if err != nil {
		return &Network{}, err
	}
	return &Network{Conn: conn}, nil
}

func (n *Network) Write(data string) error {
	_, err := n.Conn.Write([]byte(data))
	return err
}

func (n *Network) Read(size int) (string, error) {
	data := make([]byte, size)
	n, err := n.Conn.Read(data)
	if err != nil {
		return data, err
	}
	return string(data[:n]), nil
}
