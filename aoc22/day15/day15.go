package main

import (
	"bufio"
	"fmt"
	"image"
	"math"
	"os"
)

type scanner struct {
	coords   image.Point
	beacon   image.Point
	distance int
	covers   []image.Point
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	var scanners []scanner
	for sc.Scan() {
		var sx, sy, bx, by int
		fmt.Sscanf(sc.Text(),
			`Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d`,
			&sx, &sy, &bx, &by)
		d := int(math.Abs(float64(sx-bx)) + math.Abs(float64(sy-by)))
		scanners = append(scanners,
			scanner{
				coords:   image.Pt(sx, sy),
				beacon:   image.Pt(bx, by),
				distance: d,
			})
	}

	// // Part 1
	// targetRow := 2000000
	//
	// // coverageMap := make(map[image.Point]struct{})
	// beaconMap := make(map[image.Point]struct{})
	// for _, s := range scanners {
	// 	if s.beacon.Y == targetRow {
	// 		beaconMap[s.beacon] = struct{}{}
	// 	}
	// 	s.covers = getCoveredPoints(s)
	// 	// for _,c := range s.covers {
	// 	//     coverageMap[c] = struct{}{}
	// 	// }
	// }
	//
	// // tl := getMin(allCovered)
	// // br := getMax(allCovered)
	// // bounds := image.Rect(tl.X,tl.Y,br.X,br.Y)
	//
	// targetRowPoints := make(map[image.Point]struct{})
	// for _, s := range scanners {
	// 	if hitsTargetRow(s, targetRow) {
	// 		tmp := getTargetPoints(s, targetRow)
	// 		for _, p := range tmp {
	// 			targetRowPoints[p] = struct{}{}
	// 		}
	// 	}
	// }
	//
	// fmt.Println(len(targetRowPoints) - len(beaconMap))

    // Part 2 (this crashes my pc...)
    // maxGrid := 4000000
    // coverageMap := make(map[image.Point]struct{})
    // for _,s := range scanners {
    //     s.covers = getCoveredPoints(s,maxGrid)
    //     for _,c := range s.covers {
    //         coverageMap[c] = struct{}{}
    //     }
    // }
    //
    // for x:=0;x<=maxGrid;x++ {
    //     for y:=0;y<=maxGrid;y++ {
    //         if _,ok := coverageMap[image.Pt(x,y)]; !ok {
    //             fmt.Println(x*4000000+y)
    //         }
    //     }
    // }

    // Part 2 attempt 2
    maxGrid := 4000000
    for _,s := range scanners {
        oor := getOutOfReachPoints(s,maxGrid)
        for _,p := range oor {
            if p.X >=0 && p.X <= maxGrid && p.Y >= 0 && p.Y <= maxGrid {
                isValid := true
                for _,s2 := range scanners {
                    fmt.Println(manhattan(p,s2.coords))
                    fmt.Println(s2.distance)
                    fmt.Println(manhattan(p,s2.coords),"is less than",s2.distance)
                    if manhattan(p,s2.coords) <= s2.distance {
                        isValid = false
                        break
                    }
                    
                }
                if isValid {
                    fmt.Println()
                    fmt.Println(p)
                    fmt.Println(p.X*4000000+ p.Y)
                    return
                }
            }
            
        }

    }


}

func manhattan(p1,p2 image.Point) int {
    return int(math.Abs(float64(p1.X-p2.X))+math.Abs(float64(p1.Y-p2.Y)))
}

func hitsTargetRow(s scanner, target int) bool {
	if target < s.coords.Y+s.distance || target > s.coords.Y-s.distance {
		return true
	}
	return false
}

func getTargetPoints(s scanner, target int) []image.Point {
	// p := s.coords

	var points []image.Point
	dY := int(math.Abs(float64(target - s.coords.Y)))
	dX := s.distance - dY

	for x := -dX; x <= dX; x++ {
		points = append(points, image.Pt(s.coords.X+x, target))
	}

	return points
}

func getCoveredPoints(s scanner, dist, max int) []image.Point {

	var points []image.Point

    cx := s.coords.X
    cy := s.coords.Y
    cd := dist

    var minX,maxX int
    var minY,maxY int
    minY = int(math.Max(float64(cy-cd),0)) // can't go below 0
    maxY = int(math.Min(float64(cy+cd),float64(max))) // can't go above max

    for y:=minY;y<=maxY;y++ {
        dX := cd - int(math.Abs(float64(y-cy)))
        minX = int(math.Max(float64(-dX+cx),0)) // can't go below 0
        maxX = int(math.Min(float64(dX+cx),float64(max))) // can't go above max
        for x:=minX;x<=maxX;x++ {
            points = append(points,image.Pt(x,y))
        }
    }

	return points
}

func getOutOfReachPoints(s scanner, max int) []image.Point {

	var points []image.Point

    cx := s.coords.X
    cy := s.coords.Y
    cd := s.distance + 1

    for y := cy-cd; y <= cy+cd; y++ {
        dX := cd - int(math.Abs(float64(y-cy)))
        points = append(points,image.Pt(cx+dX,y))
        if dX > 0 {
            points = append(points,image.Pt(cx-dX,y))
        }
    }

	return points
}

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
