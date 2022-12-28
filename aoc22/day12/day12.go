package main

import (
	"bufio"
	"fmt"
	"image"
	"math"
	"os"
	"sort"

	"github.com/fzipp/astar"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	array := graph{}
	for sc.Scan() {
		array = append(array, sc.Text())
	}

	var sx, sy, ex, ey int
	for i, row := range array {
		for j, c := range row {
			if string(c) == "E" {
				ex = j
				ey = i
			}

		}
	}
    var distances []int
	dest := image.Pt(ex, ey)
    for i, row := range array {
		for j, c := range row {
			if string(c) == "S" || string(c) == "a" {
				sx = j
				sy = i
                start := image.Pt(sx, sy)
                path := astar.FindPath[image.Point](array, start, dest, nodeDist, nodeDist)
                distances = append(distances,len(path)-1)
			}
		}
	}
    sort.Sort((sort.IntSlice(distances)))
	fmt.Println(distances)

}

func nodeDist(p, q image.Point) float64 {
	d := q.Sub(p)
	return math.Sqrt(float64(d.X*d.X + d.Y*d.Y))
}

type graph []string

func (g graph) Neighbours(p image.Point) []image.Point {
	offsets := []image.Point{
		image.Pt(0, -1), //North
		image.Pt(1, 0),  //East
		image.Pt(0, 1),  //South
		image.Pt(-1, 0), //West
	}

	res := make([]image.Point, 0, 4)
	for _, off := range offsets {
		q := p.Add(off)
		if g.isFreeAt(p, q) {
			res = append(res, q)
		}
	}

	return res
}

func (g graph) isFreeAt(p, q image.Point) bool {
	var start, dest byte
	if g[p.Y][p.X] == 'S' {
		start = 'a'
	} else {
		start = g[p.Y][p.X]
	}
	if g.isInBounds(q) {
		if g[q.Y][q.X] == 'E' {
			dest = 'z'
		} else {
			dest = g[q.Y][q.X]
		}
	}

	//     height = 'z'
	// } else { height = g[p.Y][p.X]}

	return g.isInBounds(q) && dest <= start+1
}

func (g graph) isInBounds(p image.Point) bool {
	return p.Y >= 0 && p.X >= 0 && p.Y < len(g) && p.X < len(g[p.Y])
}
