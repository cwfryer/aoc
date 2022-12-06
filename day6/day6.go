package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func checklastfour(in []rune) bool {
    // fmt.Printf("checking %c\n", input)
    check := in[len(in)-4:]

    duplicates := 0
    // var remainder []rune
    for _,c := range check {
        for _,r := range check {
            if c == r { duplicates++ }
        }
    }

    if duplicates > 4 {
        return false
    } else {
        return true
    }
}

func checklastfourteen(in []rune) bool {
    // fmt.Printf("checking %c\n", input)
    check := in[len(in)-14:]

    duplicates := 0
    // var remainder []rune
    for _,c := range check {
        for _,r := range check {
            if c == r { duplicates++ }
        }
    }

    if duplicates > 14 {
        return false
    } else {
        return true
    }
}

func main() {
	file, err := os.Open("day6/day6.input")
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

    line := lines[0]

    count := 0
    var input []rune
    for _,c := range line {
        count++
        input = append(input,c)
        if len(input) >=14 {
            f := checklastfourteen(input)
            if f {
                fmt.Println(count)
                break
            }
        }
    }

}
