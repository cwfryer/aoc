package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	var encFile []int
	for sc.Scan() {
		i, _ := strconv.Atoi(sc.Text())
		encFile = append(encFile, i)
	}

	mix(encFile, 1)
}

var tmp []int
func mix(inFile []int, count int) []int {
    iMap := make(map[int]int,len(inFile))
    for i:=0;i<len(inFile);i++ {
        idx := i + inFile[i]
        if idx < 0 {
            idx += len(inFile) - 1
        }
        if idx > len(inFile) {
            idx = idx % len(inFile)
        }

        iMap[idx] = inFile[i]
    }
    fmt.Println(iMap)

    return []int{}
}

func cut(i int, xs []int) (int, []int) {
	y := xs[i]
	ys := append(xs[:i], xs[i+1:]...)
	return y, ys
}
