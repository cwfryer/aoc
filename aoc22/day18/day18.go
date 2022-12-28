package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cube struct {
	x, y, z int
}

func (c cube) adjacent() []cube {
	var adjacent []cube
	adjacent = append(adjacent, cube{x: c.x + 1, y: c.y, z: c.z})
	adjacent = append(adjacent, cube{x: c.x - 1, y: c.y, z: c.z})
	adjacent = append(adjacent, cube{x: c.x, y: c.y + 1, z: c.z})
	adjacent = append(adjacent, cube{x: c.x, y: c.y - 1, z: c.z})
	adjacent = append(adjacent, cube{x: c.x, y: c.y, z: c.z + 1})
	adjacent = append(adjacent, cube{x: c.x, y: c.y, z: c.z - 1})

	return adjacent
}

func (c cube) surfaceArea(allcubes []cube) int {
	sa := 6
	for _, a := range c.adjacent() {
		for _, t := range allcubes {
			if a == t {
				sa--
			}
		}
	}
	return sa
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	var allcubes []cube
	for sc.Scan() {
		line := sc.Text()
		coords := strings.Split(line, ",")
		j, _ := strconv.Atoi(coords[0])
		k, _ := strconv.Atoi(coords[1])
		l, _ := strconv.Atoi(coords[2])
		c := cube{x: j, y: k, z: l}
		allcubes = append(allcubes, c)
	}
	// Part 1
	surfaceArea := 0
	for _, c := range allcubes {
		surfaceArea += c.surfaceArea(allcubes)
	}

	fmt.Println(surfaceArea)

    // Part 2
	var maxX, maxY, maxZ int = 0, 0, 0
	var minX, minY, minZ int = 1000, 1000, 1000
	for _, c := range allcubes {
		if c.x > maxX {
			maxX = c.x
		}
		if c.x < minX {
			minX = c.x
		}
		if c.y > maxY {
			maxY = c.y
		}
		if c.y < minY {
			minY = c.y
		}
		if c.z > maxZ {
			maxZ = c.z
		}
		if c.z < minZ {
			minZ = c.z
		}
	}
	cubeSize := cube{x: maxX, y: maxY, z: maxZ}
	cubeSpace := make(map[cube]bool)
	for _, c := range allcubes {
		cubeSpace[c] = true
	}

	exterior := make(map[cube]bool)
	count := markExterior(cube{x: 0, y: 0, z: 0}, cubeSpace, exterior, cubeSize)
    fmt.Println(count)

}

func markExterior(c cube, cubeSpace map[cube]bool, exterior map[cube]bool, cubeSize cube) int {
    if exterior[c] {
        return 0
    }
    if c.x < -1 || c.x > cubeSize.x+1 || c.y < -1 || c.y > cubeSize.y+1 || c.z < -1 || c.z > cubeSize.y+1 {
        return 0
    }
    if cubeSpace[c] {
        return 1
    }
    exterior[c] = true
    var count int
    for _,n := range c.adjacent() {
        count += markExterior(n, cubeSpace, exterior, cubeSize)
    }
    return count
}
