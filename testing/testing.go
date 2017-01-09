// TESTING
// CHALLENGE DESCRIPTION:

// In many teams, there is a person who tests a project, finds bugs and errors,
// and prioritizes them. Now, you have the unique opportunity to try yourself
// as a tester and test a product. Here, you have two strings - the first one
// is provided by developers, and the second one is mentioned in design. You
// have to find and count the number of bugs, and also prioritize them for
// fixing using the following statuses: 'Low', 'Medium', 'High', 'Critical' or 'Done'.

// INPUT SAMPLE:

// The first argument is a path to a file. Each line includes a test case with
// two strings separated by a pipeline '|'. The first string is the one the
// developers provided to you for testing, and the second one is from design.

// Heelo Codevval | Hello Codeeval
// hELLO cODEEVAL | Hello Codeeval
// Hello Codeeval | Hello Codeeval
// OUTPUT SAMPLE:

// Write a program that counts the number of bugs and prioritizes them for fixing
// using the following statuses:
// 'Low' - 2 or fewer bugs;
// 'Medium' - 4 or fewer bugs;
// 'High' - 6 or fewer bugs;
// 'Critical' - more than 6 bugs;
// 'Done' - all is done;

// Low
// Critical
// Done
// CONSTRAINTS:

// Strings are of the same length from 5 to 40 characters.
// Upper and lower case matters.
// The number of test cases is 40.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func weighResults(n int) string {
	if n == 0 {
		return "Done"
	} else if n <= 2 {
		return "Low"
	} else if n <= 4 {
		return "Medium"
	} else if n <= 6 {
		return "High"
	} else {
		return "Critical"
	}
}

func solve(inp string) string {
	results := 0
	inpData := strings.Split(inp, " | ")
	testCase := inpData[0]
	goodState := inpData[1]
	for i := 0; i < len(goodState); i++ {
		if goodState[i] != testCase[i] {
			results += 1
		}
	}
	return weighResults(results)
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
