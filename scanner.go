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

	var IsUp bool
	for onion, oi := range Onions {
		switch oi["protocol"] {
		case "HTTP":
			IsUp = ScanHTTP(onion)
			if IsUp {
				fmt.Printf("%s\t[IS ALIVE]\t%s\n", onion, oi["description"])
			} else {
				fmt.Printf("%s\t[IS DEAD]\t%s\n", onion, oi["description"])
			}
		default:
			fmt.Printf("%s\t%s\tUnsupported protocol: %s\n", onion, oi["description"], oi["protocol"])
		}
	}
}
