package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
)

func solve(inp []string) string {
    parsedinput := []string{}
    for _, k := range inp {
        parsedinput = append(parsedinput, strings.ToLower(k))
    }
    result := ""
    used := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
    for _, j := range used {
        for vi, i := range parsedinput {
            if j == i {
                break
            }
            if vi == len(parsedinput)-1 {
                result += j
            }
        }
    }
    if result == "" {
        result = "NULL"
    }
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
        fmt.Println(solve(strings.Split(scanner.Text(), "")))
    }
}
