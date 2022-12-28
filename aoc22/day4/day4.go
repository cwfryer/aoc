package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type assignments struct {
    a1 []int
    a2 []int
}

func main() {
    file, err := os.Open("day4/day4.input")
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

    overlap := 0
    for _,line := range lines {
        tmp := strings.Split(line,",")
        lb1s,ub1s,_ := strings.Cut(tmp[0],"-")
        lb1,_ := strconv.Atoi(lb1s)
        ub1,_ := strconv.Atoi(ub1s)
        lb2s,ub2s,_ := strings.Cut(tmp[1],"-")
        lb2,_ := strconv.Atoi(lb2s)
        ub2,_ := strconv.Atoi(ub2s)

        if ub1 - lb1 > ub2 - lb2 {
            for i := lb2; i <= ub2; i++ {
                if i >= lb1 && i <= ub1 {
                    overlap++
                    break
                }
            }
        } else {
            for i := lb1; i <= ub1; i++ {
                if i >= lb2 && i <= ub2 {
                    overlap++
                    break
                }
            }
        }
    }

    fmt.Println(overlap)
}
