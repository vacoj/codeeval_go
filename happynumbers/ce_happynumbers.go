package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//var input_data = [3]int{1, 7, 22, 567}

func check_exists(num int, arr []int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == num {
			return true
		}
	}
	return false
}

func check_happy(num int) int {
	result := num
	counter := 0
	used_vals := make([]int, 0)
	// A happy number is defined by the following process.
	// Starting with any positive integer, replace the number
	// by the sum of the squares of its digits, and repeat the process
	// until the number equals 1 (where it will stay), or it loops endlessly
	// in a cycle which does not include 1. Those numbers for which this
	// process ends in 1 are happy numbers,
	// while those that do not end in 1 are unhappy numbers.
	// For the curious, here's why 7 is a happy number: 7->49->97->130->10->1.
	for result != 0 {
		if result == 1 {
			// return result
			return result
		} else {
			str_res := strconv.Itoa(result)
			temp_res := strings.Split(str_res, "")

			vals_list := make([]int, 0)

			for i := 0; i < len(used_vals); i++ {
				if result == used_vals[i] {
					return 0
				}
			}
			// add value to a bucket to prevent evaluation of it again.
			used_vals = append(used_vals, result)

			for i := 0; i < len(temp_res); i++ {
				t_conv, err := strconv.Atoi(temp_res[i])
				if err != nil {
					fmt.Println(err)
				}
				vals_list = append(vals_list, t_conv)
			}

			_divnum := 0
			for i := 0; i < len(vals_list); i++ {
				_divnum += vals_list[i] * vals_list[i]
			}
			if _divnum == 1 {
				return 1
			} else {
				result = _divnum
			}
			counter += 1
		}
	}
	return counter
}

func main() {
	file, err := os.Open(os.Args[1])
	// file, err := os.Open("ce_happynumbers.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		test, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		//'scanner.Text()' represents the test case, do something with it
		fmt.Println(check_happy(test))
	}
}

// for i := 0; i < len(input_data); i++ {
//  fmt.Println(input_data[i]))
// }
