package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func check_beauty(inp []string) int {
	max_val := 26
	var dict = map[string]int{}
	for i := 0; i < len(inp); i++ {
		dict[inp[i]] = 0
	}
	for i := 0; i < len(inp); i++ {
		c := strings.ToLower(inp[i])

		if c == "a" || c == "A" {
			dict[c] += 1
		} else if c == "b" {
			dict[c] += 1
		} else if c == "c" {
			dict[c] += 1
		} else if c == "d" {
			dict[c] += 1
		} else if c == "e" {
			dict[c] += 1
		} else if c == "f" {
			dict[c] += 1
		} else if c == "g" {
			dict[c] += 1
		} else if c == "h" {
			dict[c] += 1
		} else if c == "i" {
			dict[c] += 1
		} else if c == "j" {
			dict[c] += 1
		} else if c == "k" {
			dict[c] += 1
		} else if c == "l" {
			dict[c] += 1
		} else if c == "m" {
			dict[c] += 1
		} else if c == "n" {
			dict[c] += 1
		} else if c == "o" {
			dict[c] += 1
		} else if c == "p" {
			dict[c] += 1
		} else if c == "q" {
			dict[c] += 1
		} else if c == "r" {
			dict[c] += 1
		} else if c == "s" {
			dict[c] += 1
		} else if c == "t" {
			dict[c] += 1
		} else if c == "u" {
			dict[c] += 1
		} else if c == "v" {
			dict[c] += 1
		} else if c == "w" {
			dict[c] += 1
		} else if c == "x" {
			dict[c] += 1
		} else if c == "y" {
			dict[c] += 1
		} else if c == "z" {
			dict[c] += 1
		}
	}

	valscount := []int{}
	for _, v := range dict {
		if v > 0 {
			valscount = append(valscount, v)
			// fmt.Println(k)
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(valscount)))
	// fmt.Println(valscount)

	fin := 0
	for i := 0; i < len(valscount); i++ {
		fin += valscount[i] * max_val
		max_val -= 1
	}
	return fin
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
		test := strings.Split(scanner.Text(), "")
		if err != nil {
			log.Fatal(err)
		}
		//'scanner.Text()' represents the test case, do something with it
		fmt.Println(check_beauty(test))
	}
}

// find total counts per char in a map
// iterate map, sort counts Large to Small into slice
// iterate slice, multiply counts per loop, reducing max val after mult
