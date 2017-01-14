package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func solve(inp string) int {
	result := 0
	if len(strings.Split(inp, "Label")) > 1 {
		for h, i := range strings.Split(inp, "Label ") {
			if h > 0 {
				for k, j := range strings.Split(i, "\"}") {
					if k == 0 {

						n, _ := strconv.Atoi(j)
						// fmt.Println(n, j)
						result += n
						break
					}
				}
			}
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
