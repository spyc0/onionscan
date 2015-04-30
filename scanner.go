package main

import (
	"fmt"
	"os"
	"sync"
)

var (
	isUp bool
	wg sync.WaitGroup
)

func ShowUsage() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s onionlist.txt\n", os.Args[0])
		os.Exit(0)
	}
}

func Scanner(onion string, oi map[string]string, w *sync.WaitGroup) {
	switch oi["protocol"] {
	case "HTTP":
		isUp = ScanHTTP(onion)
		if isUp {
			fmt.Printf("%s\t[IS ALIVE]\t%s\n", onion, oi["description"])
		} else {
			fmt.Printf("%s\t[IS DEAD]\t%s\n", onion, oi["description"])
		}
	default:
		fmt.Printf("%s\t%s\tUnsupported protocol: %s\n", onion, oi["description"], oi["protocol"])
	}
	wg.Done()
}

func main() {
	ShowUsage()
	Onions := ParseOnions(os.Args[1])

	for onion, oi := range Onions {
		wg.Add(1)
		go Scanner(onion, oi, &wg)
	}
	wg.Wait()
}
