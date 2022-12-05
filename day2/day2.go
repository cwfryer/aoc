package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main () {
    file, err := os.Open("day2/day2-input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer func() {
        if err = file.Close(); err != nil {
            log.Fatal(err)
        }
    }()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    m := make(map[string]int)
    m["A"] = 1
    m["B"] = 2
    m["C"] = 3

    o := make(map[string]string)
    o["A X"] = "C"
    o["B X"] = "A"
    o["C X"] = "B"
    o["A Y"] = "A"
    o["B Y"] = "B"
    o["C Y"] = "C"
    o["A Z"] = "B"
    o["B Z"] = "C"
    o["C Z"] = "A"
    var score int = 0
    for _,line := range lines {
        p2 := strings.Split(line," ")[1]

        if p2 == "X" {
            score += m[o[line]]
        }
        if p2 == "Y" {
            score += m[o[line]]
            score += 3
        }
        if p2 == "Z" {
            score += m[o[line]]
            score += 6
        }
    }

    fmt.Println(score)
}
