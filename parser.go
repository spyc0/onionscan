package main

import (
	"os"
	"bufio"
	"strings"
	"fmt"
	"regexp"
)

func ParseOnions(filename string) map[string]map[string]string {
	OnionsFile, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Failed to read list file: %s\n", err)
		os.Exit(1)
	}
	Onions := make(map[string]map[string]string)
	LineScanner := bufio.NewScanner(OnionsFile)
	re := regexp.MustCompile("[a-z2-7]{16}")
	for LineScanner.Scan() {
		if len(LineScanner.Text()) == 0 {
			continue
		}
		if !re.Match([]byte(strings.Split(LineScanner.Text(), "|")[0])) {
			continue
		}
		if len(strings.Split(LineScanner.Text(), "|")) < 3 {
			fmt.Printf("Invalid/Not onion list\n")
			os.Exit(1)
		}
		OnionInfo := strings.Split(LineScanner.Text(), "|")
		Onions[OnionInfo[0]] = map[string]string{
			"protocol": OnionInfo[2],
			"description": OnionInfo[1],
		}
	}
	if err = LineScanner.Err(); err != nil {
		fmt.Printf("Failed to parse onion list\n")
	}
	return Onions
}
