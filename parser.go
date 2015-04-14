package main

import (
	"os"
	"bufio"
	"strings"
	"log"
)

func ParseOnions(filename string) map[string]map[string]string {
	OnionsFile, err := os.Open(filename)
	if err != nil {
		log.Fatal("Failed to read list file")
	}
	Onions := make(map[string]map[string]string)
	LineScanner := bufio.NewScanner(OnionsFile)
	for LineScanner.Scan() {
		OnionInfo := strings.Split(LineScanner.Text(), "|")
		Onions[OnionInfo[0]] = map[string]string{
			"protocol": OnionInfo[2],
			"description": OnionInfo[1],
		}
	}
	if err = LineScanner.Err(); err != nil {
		log.Fatal("Failed to parse onion list")
	}
	return Onions
}
