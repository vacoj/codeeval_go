package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func split(r rune) bool {
	return r == '+' || r == '-'
}

func numSplit(r rune) bool {
	return r == '0' || r == '1' || r == '2' || r == '3' || r == '3' || r == '4' || r == '5' || r == '6' || r == '7' || r == '8' || r == '9'
}

func calculate(inp string) int {
	result := 0
	all := strings.FieldsFunc(inp, split)
	ops := strings.FieldsFunc(inp, numSplit)
	// fmt.Println(all, ops)
	for n, i := range all {

		if n == 0 {
			ii, _ := strconv.Atoi(i)
			result = ii
		} else if ops[n-1] == "+" {
			ii, _ := strconv.Atoi(i)
			result += ii
		} else {
			ii, _ := strconv.Atoi(i)
			result -= ii
		}
	}
	return result
}

func solve(inp string) int {
	numbers := strings.Split(strings.Split(inp, " ")[0], "")
	letters := strings.Split(strings.Split(strings.Replace(strings.Replace(inp, "+", "", -1), "-", "", -1), " ")[1], "")
	mapper := map[string]string{}
	for i := 0; i < len(letters); i++ {
		mapper[letters[i]] = numbers[i]
	}

	eq := ""
	for _, i := range strings.Split(inp, " ")[1] {
		if i == '-' || i == '+' {
			eq += string(i)
		} else {
			eq += mapper[string(i)]
		}
	}
	return calculate(eq)
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
