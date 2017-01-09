// AV2WHIECX 45
// BH3

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var morsemap = map[string]string{
	".-":    "A",
	"-...":  "B",
	"-.-.":  "C",
	"-..":   "D",
	".":     "E",
	"..-.":  "F",
	"--.":   "G",
	"....":  "H",
	"..":    "I",
	".---":  "J",
	"-.-":   "K",
	".-..":  "L",
	"--":    "M",
	"-.":    "N",
	"---":   "O",
	".--.":  "P",
	"--.-":  "Q",
	".-.":   "R",
	"...":   "S",
	"-":     "T",
	"..-":   "U",
	"...-":  "V",
	".--":   "W",
	"-..-":  "X",
	"-.--":  "Y",
	"--..":  "Z",
	"-----": "0",
	".----": "1",
	"..---": "2",
	"...--": "3",
	"....-": "4",
	".....": "5",
	"-....": "6",
	"--...": "7",
	"---..": "8",
	"----.": "9",
}

func solve(inp string) {
	result := ""
	newWords := strings.Split(inp, "  ")

	for i := 0; i < len(newWords); i++ {
		word := strings.Split(newWords[i], " ")
		for i := 0; i < len(word); i++ {
			result += morsemap[word[i]]
		}
		result += " "
	}
	fmt.Println(result)

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
		solve(scanner.Text())
	}
}
