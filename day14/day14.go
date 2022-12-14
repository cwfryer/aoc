package main

import (
	"bufio"
	"fmt"
	"image"
	"math"
	"os"
	"strconv"
	"strings"
)

func getMin(coords []image.Point) image.Point { // return top left corner
	var min image.Point
	for i, p := range coords {

		if i == 0 {
			min = image.Pt(p.X, 0)
		}
		for _, c := range coords {
			if c.X < min.X {
				min.X = c.X
			}
			if c.Y < min.Y {
				min.Y = c.Y
			}
		}
	}
	return min
}

func getMax(coords []image.Point) image.Point { // return top left corner
	var max image.Point
	for i, p := range coords {
		if i == 0 {
			max = p
		}
		for _, c := range coords {
			if c.X > max.X {
				max.X = c.X
			}
			if c.Y > max.Y {
				max.Y = c.Y
			}
		}
	}

	return max
}

func checkSteps(p image.Point, ob []image.Point, b image.Rectangle) (image.Point, error) {
	if p.Y > b.Max.Y+2 {
		return image.Point{}, fmt.Errorf("out of bounds")
	}
	check := p.Add(image.Pt(0, 1))
	// if i can go do, do it
	obstacle := false
	for _, o := range ob {
		if o == check {
			obstacle = true
			break
		}
	}
    if check.Y == b.Max.Y + 2 {
        obstacle = true
    }
	if !obstacle {
		check, err := checkSteps(check, ob, b)
		return check, err
	}

	// if i cant go down, try going down-left
	check = p.Add(image.Pt(-1, 1))
	// passby := p.Add(image.Pt(-1, 0)) // you have to be able to go left first
	obstacle = false
	for _, o := range ob {
		if o == check {
			obstacle = true
			break
		}
		// if o == passby {
		//     obstacle = true
		//     break
		// }
	}
    if check.Y == b.Max.Y + 2 {
        obstacle = true
    }
	if !obstacle {
		check, err := checkSteps(check, ob, b)
		return check, err
	}

	// if i can't go down-left, try going down-right
	check = p.Add(image.Pt(1, 1))
	// passby = p.Add(image.Pt(1, 0)) // you have to be able to go left first
	obstacle = false
	for _, o := range ob {
		if o == check {
			obstacle = true
			break
		}
		// if o == passby {
		//     obstacle = true
		//     break
		// }
	}
    if check.Y == b.Max.Y + 2 {
        obstacle = true
    }
	if !obstacle {
		return checkSteps(check, ob, b)
	}

	return p, nil

}

func dropSand(obstacles []image.Point, bounds image.Rectangle) image.Point { //

	newSand := image.Pt(500, 0)
	var err error
	newSand, err = checkSteps(newSand, obstacles, bounds)
	if err != nil {
		return image.Pt(-10000000, -10000000)
	}

	return newSand
}

func defineObstacles(lines [][]image.Point) []image.Point {
	obstacles := []image.Point{}
	for _, line := range lines {
		for i := 1; i < len(line); i++ {
			start := line[i-1]
			end := line[i]
			dX := start.X - end.X
			dY := start.Y - end.Y

			if dX != 0 {
				mx := int(math.Max(float64(start.X), float64(end.X)))
				mn := int(math.Min(float64(start.X), float64(end.X)))
				for i := mn; i <= mx; i++ {
					obstacles = append(obstacles, image.Pt(i, start.Y))
				}
			}
			if dY != 0 {
				mx := int(math.Max(float64(start.Y), float64(end.Y)))
				mn := int(math.Min(float64(start.Y), float64(end.Y)))
				for i := mn; i <= mx; i++ {
					obstacles = append(obstacles, image.Pt(start.X, i))
				}
			}
		}
	}
	return obstacles
}

func addFloor(obstacles []image.Point, bounds image.Rectangle) []image.Point {
	floor := []image.Point{}
	for i := -100000; i < 100000; i++ {
		floor = append(floor, image.Pt(i+500, bounds.Max.Y+2))
	}
	return append(obstacles, floor...)
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	var coordList [][]image.Point

	for sc.Scan() {
		line := sc.Text()
		tmp := strings.Split(line, " -> ")
		coords := []image.Point{}
		for _, c := range tmp {
			l, r, _ := strings.Cut(c, ",")
			x, _ := strconv.Atoi(l)
			y, _ := strconv.Atoi(r)
			coords = append(coords, image.Pt(x, y))
		}
		coordList = append(coordList, coords)
	}

	obstacles := defineObstacles(coordList)
	tl := getMin(obstacles)
	br := getMax(obstacles)
	bounds := image.Rect(tl.X, tl.Y, br.X, br.Y)
	fmt.Println(bounds)
	// obstacles = addFloor(obstacles,bounds)
	// tl = getMin(obstacles)
	// br = getMax(obstacles)
	// bounds = image.Rect(tl.X, tl.Y, br.X, br.Y)
	// fmt.Println(bounds)

	sand := []image.Point{}
	for {
		// start := len(sand)
		s := dropSand(obstacles, bounds)
        if s.Y <= bounds.Max.Y + 2 {
			obstacles = append(obstacles, s)
			sand = append(sand, s)
        }
		// if s.In(bounds) {
		// 	obstacles = append(obstacles, s)
		// 	sand = append(sand, s)
		// }
		if s == image.Pt(500, 0) {
			break
		}

		// fmt.Println(sand)
		// end := len(sand)
		// if start == end {
		// 	break
		// }
	}

	fmt.Println(len(sand))

}
