// ONE ZERO, TWO ZEROS...
// CHALLENGE DESCRIPTION:

// Our agent uncovered a global criminal money-laundering network that used offshore companies to defraud international organizations of total $1,000,000,000! The agent changes his location each hour, but he manages to send us the code that we need to decipher.
// Deciphering code includes many stages, and you are taking part in one of them. Therefore, your task is the following: you have two numbers – the first one is the number of zeros in a binary code and the second one shows the range from 1 to this number, where you have to find these zeros.
// For example, for the given numbers 2 and 4, you convert all numbers from 1 to 4 inclusive into the binary system. As a result, you get 1, 10, 11, and 100. As the first given number is 2, this means that we are looking for numbers with two zeros, so only 100 suits us. Hence, the result will be 1: there is only one number with two zeros.

// The first argument is a path to a file. Each line includes a test case with two numbers: the first one is the number of zeros in a binary code that we need to find and the second one is the range from 1 to this number where you have to find these zeros.

// For example:
// INPUT SAMPLE:
// 1 8
// 2 4

// Print the total number of numerals that contain the needed amount of zeros in a binary system.

// For example:
// OUTPUT SAMPLE:
// 3
// 1
// CONSTRAINTS:

// Range can be from 5 to 1000.
// Number of zeros does not exceed the length of binary code number.
// The number of test cases is 40.

// n := int64(123)

// fmt.Println(strconv.FormatInt(n, 2)) // 1111011

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func findMatches(zs int, inp [][]string) int {
	// fmt.Println("Count we are looking for is:", zs)
	finalCount := 0

	for _, i := range inp {
		localCount := 0
		for _, j := range i {
			if j == "0" {
				localCount++
			}
		}
		if localCount == zs {
			finalCount++
		}
	}
	// fmt.Println(inp)
	return finalCount
}

func convertToSlices(zs int, rng int) int {
	bins := []int{}
	strNums := [][]string{}
	for i := 1; i <= rng; i++ {
		n := strconv.FormatInt(int64(i), 2)
		ns, _ := strconv.Atoi(n)
		bins = append(bins, ns)
	}

	for _, i := range bins {
		si := strconv.Itoa(i)
		strNums = append(strNums, strings.Split(si, ""))
	}
	return findMatches(zs, strNums)
}

func solve(zs int, rng int) {
	fmt.Println(convertToSlices(zs, rng))
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
		if err != nil {
			log.Fatal(err)
		}
		input := strings.Split(scanner.Text(), " ")
		rng, _ := strconv.Atoi(input[1])
		zs, _ := strconv.Atoi(input[0])
		solve(zs, rng)
	}
}
