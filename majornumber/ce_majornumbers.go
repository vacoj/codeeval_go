package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// The major element in a sequence with the length of L is the element
// which appears in a sequence more than L/2 times. The challenge is to
// find that element in a sequence.

var input = [][]int{
	{92, 19, 19, 76, 19, 21, 19, 85, 19, 19, 19, 94, 19, 19, 22, 67, 83, 19, 19, 54, 59, 1, 19, 19},
	{92, 11, 30, 92, 1, 11, 92, 38, 92, 92, 43, 92, 92, 51, 92, 36, 97, 92, 92, 92, 43, 22, 84, 92, 92},
	{4, 79, 89, 98, 48, 42, 39, 79, 55, 70, 21, 39, 98, 16, 96, 2, 10, 24, 14, 47, 0, 50, 95, 20, 95, 48, 50, 12, 42}}

func find_major(inp []int) int {
	bigL := len(inp) / 2
	var result int
	var dict = map[int]int{}
	for i := 0; i < len(inp); i++ {
		dict[inp[i]] = 0
	}
	result = 0
	for i := 0; i < len(inp); i++ {
		dict[inp[i]] += 1
	}
	max_dictkey := 0
	max_dictval := 0
	for k, v := range dict {
		if v > max_dictval {
			max_dictkey = k
			max_dictval = v
		} else {
			result = max_dictkey
		}
	}
	if dict[max_dictkey] > bigL {
		return result
	} else {
		return -987654323
	}
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
		data_s := strings.Split(scanner.Text(), ",")
		data_i := make([]int, 0)
		for i := 0; i < len(data_s); i++ {
			v, err := strconv.Atoi(data_s[i])
			if err != nil {
				log.Fatal(err)
			}
			data_i = append(data_i, v)
		}
		result := find_major(data_i)
		if result != -987654323 {
			fmt.Println(result)
		} else {
			fmt.Println("None")
		}
	}
}
