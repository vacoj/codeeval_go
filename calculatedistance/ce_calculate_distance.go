// You have coordinates of 2 points and need to find the distance between them.

// INPUT SAMPLE:

// Your program should accept as its first argument a path to a filename. Input example is the following

// (25, 4) (1, -6)
// (47, 43) (-25, -11)
// All numbers in input are integers between -100 and 100.

// OUTPUT SAMPLE:

// Print results in the following way.

// 26
// 90
// You don't need to round the results you receive. They must be integer numbers.

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

func solve(s []string) float64 {
	d := []int{}
	for i := 0; i < len(s); i++ {
		v, _ := strconv.Atoi(s[i])
		d = append(d, v)
	}
	x1, y1, x2, y2 := d[0], d[1], d[2], d[3]
	return math.Sqrt(float64((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1)))
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
		s := scanner.Text()
		s = strings.Replace(s, "(", "", -1)
		s = strings.Replace(s, ")", "", -1)
		s = strings.Replace(s, ",", "", -1)
		data_i := strings.Split(s, " ")
		fmt.Println(solve(data_i))
	}
}
