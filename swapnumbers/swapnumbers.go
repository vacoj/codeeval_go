package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// 0Always4 8look5 9on4 2the7 8bright4 7side9 8of3 5life5
// 5Nobody5 3expects7 4the5 4Spanish6 0inquisition9

func swapnums(inp string) string {
	result := ""
	s := strings.Split(inp, "")
	s[0], s[len(s)-1] = s[len(s)-1], s[0]
	for i := 0; i < len(s); i++ {
		result += s[i]
	}
	return result
}

func solve(inp string) string {
	swapped := strings.Split(inp, " ")
	result := ""

	for _, i := range swapped {
		result += swapnums(i) + " "
	}
	return result
}

func main() {
	// file, err := os.Open(os.Args[1])
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(solve(scanner.Text()))
	}
}
