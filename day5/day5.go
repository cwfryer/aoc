package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
    file, err := os.Open("day5/day5.test")
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
}
