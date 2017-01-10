// However, it was not implemented until 1998 and 2000
// The first programming language
// The Manchester Mark 1 ran programs written in Autocode from 1952

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func solve(inp string) string {
	mapper := map[int]string{}
	result := ""
	setup := strings.Split(inp, ";")
	scrambled, indexes := strings.Split(setup[0], " "), strings.Split(setup[1], " ")

	indexes_i := []int{}

	for i := 0; i < len(indexes); i++ {
		ii, _ := strconv.Atoi(indexes[i])
		indexes_i = append(indexes_i, ii)
	}

	mappedWords := []string{}

	for i := 0; i < len(indexes_i); i++ {
		// bug is here?
		mapper[indexes_i[i]] = scrambled[i]
		mappedWords = append(mappedWords, scrambled[i])
	}

	sortedIndexes := indexes_i
	sort.Sort(sort.IntSlice(sortedIndexes))

	missingWord := ""
	for i := 0; i < len(scrambled); i++ {
		if !stringInSlice(scrambled[i], mappedWords) {
			missingWord = scrambled[i]
			break
		}
	}

	missingIndex := -1
	for i := 1; i <= len(sortedIndexes); i++ {
		if i != sortedIndexes[i-1] {
			missingIndex = i
			break
		}
	}
	if missingIndex > -1 {
		mapper[missingIndex] = missingWord
	}
	for k, v := range mapper {
		scrambled[k-1] = v

	}

	for i := 0; i < len(scrambled); i++ {
		result += scrambled[i]
		if i < len(scrambled) {
			result += " "
		}
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
