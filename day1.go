package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
    file, err := os.Open("day1-input.txt")
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

    var calsums []int
    var cal int 
    cal = 0
    for _,line := range lines {
        if line == "" {
            calsums = append(calsums,cal)
            cal = 0
        } else {
            i,_ := strconv.Atoi(line)
            cal += i
        }
    }

    var m1,m2,m3, t1,t2,t3 int
    for _,c := range calsums {
        if c > t1{
            t1 = c
            m1 = t1
        }
    }
    for _,c := range calsums {
        if c > t2 && c < m1{
            t2 = c
            m2 = t2
        }
    }
    for _,c := range calsums {
        if c > t3 && c < m2{
            t3 = c
            m3 = t3
        }
    }

    fmt.Println(m1+m2+m3)
}
