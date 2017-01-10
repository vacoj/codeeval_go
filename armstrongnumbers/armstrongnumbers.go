// True
// True
// False

package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func solve(inp string) string {
	sum := 0.0
	comp_res, _ := strconv.Atoi(inp)
	result := "False"
	numbers := strings.Split(inp, "")
	numbers_i := []int{}
	n := len(inp)
	for i := 0; i < len(numbers); i++ {
		num_i, _ := strconv.Atoi(numbers[i])
		numbers_i = append(numbers_i, num_i)
	}
	for i := 0; i < len(numbers_i); i++ {
		sum += math.Pow(float64(numbers_i[i]), float64(n))
	}
	if comp_res == int(sum) {
		return "True"
	} else {
		return "False"
	}
	fmt.Println(sum)
	fmt.Println(numbers)
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
