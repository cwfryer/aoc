package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main () {
    file, err := os.Open("day3/day3.input.txt")
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

    priorities := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    
    dupes := make(map[int]map[rune]int)
    for i := 0; i < len(lines)/3; i++ {
        dupes[i] = make(map[rune]int)
        l1 := lines[3*i]
        l2 := lines[3*i+1]
        l3 := lines[3*i+2]

        for _,c1 := range l1 {
            for _,c2 := range l2 {
                for _,c3 := range l3 {
                    if c1 == c2 && c1 == c3 {
                        dupes[i][c1] = 1 + strings.Index(priorities,string(c1))
                    }
                }
            }
        }
    }

    score := 0
    fmt.Println(dupes)
    for _,i := range dupes {
        for _,p := range i {
            fmt.Println(p)
            score += p
        }
    }

    fmt.Println(score)
}
