// Write a program which finds the first non-repeated character in a string.

// INPUT SAMPLE:

// The first argument is a path to a file. The file contains strings.

// For example:
// yellow
// tooth
// OUTPUT SAMPLE:

// Print to stdout the first non-repeated character, one per line.

// For example:
// y
// h

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func solve(inp string) string {
	inp_slc := strings.Split(inp, "")
	items_count := make(map[string]int)
	result := ""

	for i := 0; i < len(inp_slc); i++ {
		items_count[inp_slc[i]] += 1
	}
	for i := 0; i < len(inp_slc); i++ {
		if result == "" && items_count[inp_slc[i]] == 1 {
			result = inp_slc[i]
		}
	}
	return result
}

func main() {
	file, err := os.Open(os.Args[1])
	// file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(solve(scanner.Text()))
	}
}
