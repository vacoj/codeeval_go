// There is a game where each player picks a number from 1 to 9, writes it
// on a paper and gives to a guide. A player wins if his number is the lowest
// unique. We may have 10-20 players in our game.

// INPUT SAMPLE:

// Your program should accept as its first argument a path to a filename.

// You're a guide and you're given a set of numbers from players for
// the round of game. E.g. 2 rounds of the game look this way:

// 3 3 9 1 6 5 8 1 5 3
// 9 2 9 9 1 8 8 8 2 1 1
// OUTPUT SAMPLE:

// Print a winner's position or 0 in case there is no winner.
// In the first line of input sample the lowest unique number is 6. So player 5 wins.

// 5
// 0

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func solve(inp []int) int {
	uql := find_uniques(inp)
	wv := find_lowest(uql)
	answer := 0
	for i := 0; i < len(inp); i++ {
		if wv == inp[i] {
			answer = i + 1
		}
	}
	return answer
}

func find_uniques(inp []int) []int {
	var dict = map[int]int{}
	result := []int{}
	for i := 0; i < len(inp); i++ {
		dict[inp[i]] += 1
	}
	for k, v := range dict {
		if v == 1 {
			result = append(result, k)
		}
	}
	return result
}

func find_lowest(inp []int) int {
	var l int
	set := false
	for i := 0; i < len(inp); i++ {
		if inp[i] < l || !set {
			l = inp[i]
			set = true
		}
	}
	return l
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
		data_s := strings.Split(scanner.Text(), " ")
		data_i := []int{}
		for i := 0; i < len(data_s); i++ {
			n, _ := strconv.Atoi(data_s[i])
			data_i = append(data_i, n)
		}

		// test, err := strconv.Atoi(scanner.Text())
		// if err != nil {
		// 	log.Fatal(err)
		// }
		//'scanner.Text()' represents the test case, do something with it
		fmt.Println(solve(data_i))
	}
}
