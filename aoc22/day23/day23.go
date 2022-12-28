package main

import (
	"bufio"
	"fmt"
	"image"
	"os"
)

type Elf struct {
	coords   image.Point
	proposed image.Point
	consider []condition
}

type condition struct {
	check []image.Point
	move  image.Point
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	elves := []*Elf{}
	row := 0
	for sc.Scan() {
		line := sc.Text()
		for col, c := range line {
			if c == '#' {
				cons := []condition{
					{check: []image.Point{{0, -1}, {-1, -1}, {1, -1}}, move: image.Pt(0, -1)}, // north
					{check: []image.Point{{0, 1}, {-1, 1}, {1, 1}}, move: image.Pt(0, 1)},  // south
					{check: []image.Point{{-1, 0}, {-1, -1}, {-1, 1}}, move: image.Pt(-1, 0)}, // west
					{check: []image.Point{{1, 0}, {1, -1}, {1, 1}}, move: image.Pt(1, 0)},  // east
				}
				elves = append(elves, &Elf{coords: image.Pt(col, row), consider: cons})
			}
		}
		row++
	}

    var err error
    i := 1
    for {
        elves,err = round(elves)
        if err != nil {
            fmt.Println(i)
            break
        }
        i++
    }
}

func printMap(elves []*Elf) {
    var minX, maxX, minY, maxY int = 999, 0, 999, 0
    eMap := make(map[image.Point]struct{})
	for _, e := range elves {
		if e.coords.X < minX {
			minX = e.coords.X
		}
		if e.coords.Y < minY {
			minY = e.coords.Y
		}
		if e.coords.X > maxX {
			maxX = e.coords.X
		}
		if e.coords.Y > maxY {
			maxY = e.coords.Y
		}
        eMap[e.coords] = struct{}{}
	}

    printStr := ""
    for row := minY;row<=maxY;row++ {
        rowStr := ""
        for col := minX;col<=maxX;col++ {
            if _,ok := eMap[image.Pt(col,row)]; ok {
                rowStr += "#" 
            } else {
                rowStr += "."
            }
        }
        rowStr += "\n"
        printStr += rowStr
    }
    fmt.Print(printStr)

}
func part1(elves []*Elf) {
	var minX, maxX, minY, maxY int = 999, 0, 999, 0
	for _, e := range elves {
		if e.coords.X < minX {
			minX = e.coords.X
		}
		if e.coords.Y < minY {
			minY = e.coords.Y
		}
		if e.coords.X > maxX {
			maxX = e.coords.X
		}
		if e.coords.Y > maxY {
			maxY = e.coords.Y
		}
	}

	dx := maxX - minX + 1
	dy := maxY - minY + 1

	fmt.Println(dx*dy - len(elves))
}

func round(elves []*Elf) ([]*Elf,error) {
    var err error
    elves,err = propose(elves)
    if err != nil {
        return elves,err
    }
	elves = move(elves)
	elves = reorder(elves)
	return elves,nil
}

func propose(elves []*Elf) ([]*Elf,error) {
    var noMove int = 0
	for _, e := range elves {
		var possible []image.Point
	Loop:
		for _, o := range e.consider {
			for _, d := range o.check {
				ck := e.coords.Add(d)
				for _, ce := range elves {
					if ck == ce.coords {
						continue Loop
					}
				}
			}
			possible = append(possible, e.coords.Add(o.move))
		}
		if len(possible) == 4 || len(possible) == 0 {
			e.proposed = e.coords
            noMove ++

        } else {
			e.proposed = possible[0]
		}
	}
    if noMove == len(elves) {
        return elves,fmt.Errorf("stop")
    }
	return elves,nil
}

func move(elves []*Elf) []*Elf {
	pMap := make(map[image.Point]int)
	for _, e := range elves {
		pMap[e.proposed] += 1
	}

	for _, e := range elves {
		p := pMap[e.proposed]
		if p == 1 {
			e.coords = e.proposed
		}
	}

	return elves
}

func reorder(elves []*Elf) []*Elf {
	for _, e := range elves {
		opts := []condition{}
		opts = append(opts, e.consider[1:]...)
		opts = append(opts, e.consider[0])
		e.consider = opts
	}
	return elves
}
