package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func walkToEdges(a [][]string, x int, y int) bool {

	n := func() bool {
		v, _ := strconv.Atoi(a[y][x])
		for step := y - 1; step >= 0; step-- {
			height := a[step][x]
			k, _ := strconv.Atoi(height)
			if k >= v {
				return false
			}
		}
		return true
	}

	s := func() bool {
		v, _ := strconv.Atoi(a[y][x])
		for step := y + 1; step < len(a); step++ {
			height := a[step][x]
			k, _ := strconv.Atoi(height)
			if k >= v {
				return false
			}
		}
		return true
	}

	e := func() bool {
		v, _ := strconv.Atoi(a[y][x])
		for step := x + 1; step < len(a[0]); step++ {
			height := a[y][step]
			k, _ := strconv.Atoi(height)
			if k >= v {
				return false
			}
		}
		return true
	}

	w := func() bool {
		v, _ := strconv.Atoi(a[y][x])
		for step := x - 1; step >= 0; step-- {
			height := a[y][step]
			k, _ := strconv.Atoi(height)
			if k >= v {
				return false
			}
		}
		return true
	}

	visible := n() || s() || e() || w()
	return visible
}

func scenicScore(a [][]string, x int, y int) int {
	n := func() int {
		v, _ := strconv.Atoi(a[y][x])
		for step := y - 1; step >= 0; step-- {
			height := a[step][x]
			k, _ := strconv.Atoi(height)
			if k >= v {
				return (y-step)
			}
		}
		return y
	}

	s := func() int {
		v, _ := strconv.Atoi(a[y][x])
		for step := y + 1; step < len(a); step++ {
			height := a[step][x]
			k, _ := strconv.Atoi(height)
			if k >= v {
				return (step-y)
			}
		}
		return len(a)-1-y
	}

	e := func() int {
		v, _ := strconv.Atoi(a[y][x])
		for step := x + 1; step < len(a[0]); step++ {
			height := a[y][step]
			k, _ := strconv.Atoi(height)
			if k >= v {
				return (step-x)
			}
		}
		return len(a[0])-1-x
	}

	w := func() int {
		v, _ := strconv.Atoi(a[y][x])
		for step := x - 1; step >= 0; step-- {
			height := a[y][step]
			k, _ := strconv.Atoi(height)
			if k >= v {
				return (x-step)
			}
		}
		return x
	}

	score := n() * s() * e() * w()
	return score
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	var treeArray [][]string
	for sc.Scan() {
		line := sc.Text()
		array := strings.Split(line, "")
		treeArray = append(treeArray, array)
	}

	// ------------------------------
	// Part 1
	// ------------------------------

	// visible := 0
	// for y:=0;y<len(treeArray);y++ {
	//     for x:=0;x<len(treeArray[0]);x++ {
	//         seen := walkToEdges(treeArray,x,y)
	//         if seen {
	//             visible ++
	//         }
	//     }
	// }

	// fmt.Println(visible)

	// ------------------------------
	// Part 2
	// ------------------------------
    max := 0
	for y:=0;y<len(treeArray);y++ {
	    for x:=0;x<len(treeArray[0]);x++ {
	        score := scenicScore(treeArray,x,y)
	        if score > max {
                max = score
	        }
	    }
    }
	fmt.Println(max)

}
