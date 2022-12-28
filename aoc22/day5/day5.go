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

func main() {
	file, err := os.Open("day5/day5.input")
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

	// Get the crate lines
	var start []string
	for _, line := range lines {
		if line == "" {
			break
		}
		start = append(start, line)
	}
	// Reverse the lines that have crate info
	for i, j := 0, len(start)-1; i < j; i, j = i+1, j-1 {
		start[i], start[j] = start[j], start[i]
	}

	crates := make(map[int][]string)
	for i, line := range start {
		if i == 0 { // dont include the column index line
			continue
		}
		cols := math.Ceil(float64(len(line)) / 4)
		for i := 0; i < int(cols); i++ {
			if string(line[1+(4*i)]) == " " {
				continue
			} else {
				crates[i+1] = append(crates[i+1], string(line[1+(4*i)]))
			}
		}
	}

	// Get the instruction lines
	for _, line := range lines {
		if line == "" {
			continue
		}
		// if it starts with m its an instruction
		if string(line[0]) == "m" {
			tmp := strings.TrimPrefix(line, "move ")
			count, _ := strconv.Atoi(strings.SplitN(tmp, " ", 2)[0])
			_, tmp, _ = strings.Cut(line, "from ")
			start, _ := strconv.Atoi(string(tmp[0]))          //starting pile
			end, _ := strconv.Atoi(string(line[len(line)-1])) //ending pile

			k := len(crates[start]) - 1 // index of the top crate
			if count > 1 {
				grab := crates[start][k-(count-1) : k+1] // grab the count of crates to top
				crates[start] = append(crates[start][:k-(count-1)], crates[start][k+1:]...) // remake the start pile minus the taken
				crates[end] = append(crates[end], grab...) // append the taken to the end pile
			} else {
				grab := crates[start][k]
				crates[start] = append(crates[start][:k], crates[start][k+1:]...)
				crates[end] = append(crates[end], grab)
			}
		}
	}

	var final string
	for i := 1; i <= len(crates); i++ { // loop over the piles
		k := len(crates[i]) - 1 // index of the top crate
		final += crates[i][k] // read out the top crate and append to string
	}

	fmt.Println(final)
}
