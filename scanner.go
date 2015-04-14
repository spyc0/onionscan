package main

import (
	"fmt"
	"os"
)

func ShowUsage() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s onionlist.txt\n", os.Args[0])
		os.Exit(0)
	}
}

func main() {
	ShowUsage()
	Onions := ParseOnions(os.Args[1])

	for onion, oi := range Onions {
		fmt.Println(onion, oi["protocol"], oi["description"])
	}
}
