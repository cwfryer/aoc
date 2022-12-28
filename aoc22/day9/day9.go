package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func updateTail(pMap map[string][]int) map[string][]int {
    // H case
    if p, found := pMap["h"]; found {
        distance := int(math.Sqrt(math.Pow(float64(p[0]-pMap["1"][0]),2)+math.Pow(float64(p[1]-pMap["1"][1]),2)))
        if distance > 1 {
            // Calculate the new coordinates
            dx := math.Abs(float64(p[0] - pMap["1"][0]))
            dy := math.Abs(float64(p[1] - pMap["1"][1]))
            fmt.Printf("dx: %v, dy: %v\n",dx,dy)
            if dx > 1 && dy == 1 {
                x := pMap["1"][0] + int(math.Copysign(dx-1,float64(p[0]-pMap["1"][0])))
                pMap["1"] = []int{x,p[1]}
            } else if dy > 1 && dx == 1 {
                y := pMap["1"][1] + int(math.Copysign(dy-1,float64(p[1]-pMap["1"][1])))
                pMap["1"] = []int{p[0],y}
            } else {
                if dy > 0 {
                    y := pMap["1"][1] + int(math.Copysign(dy-1,float64(p[1]-pMap["1"][1])))
                    pMap["1"] = []int{pMap["1"][0],y}
                }
                if dx > 0 {
                    x := pMap["1"][0] + int(math.Copysign(dx-1,float64(p[0]-pMap["1"][0])))
                    pMap["1"] = []int{x,pMap["1"][1]}
                }
            }
        }
    } else {
        panic("key was missing")
    }

    if len(pMap) > 2 {
        // do the loop
        for i:=1;i<len(pMap)-1;i++{
            p1 := pMap[fmt.Sprint(i)]
            p2 := pMap[fmt.Sprint(i+1)]

            distance := int(math.Sqrt(math.Pow(float64(p1[0]-p2[0]),2)+math.Pow(float64(p1[1]-p2[1]),2)))
            if distance > 1 {
                // Calculate the new coordinates
                dx := math.Abs(float64(p1[0] - p2[0]))
                dy := math.Abs(float64(p1[1] - p2[1]))
                fmt.Printf("dx: %v, dy: %v\n",dx,dy)
                if dx > 1 && dy == 1 {
                    x := p2[0] + int(math.Copysign(dx-1,float64(p1[0]-p2[0])))
                    pMap[fmt.Sprint(i+1)] = []int{x,p1[1]}
                } else if dy > 1 && dx == 1 {
                    y := p2[1] + int(math.Copysign(dy-1,float64(p1[1]-p2[1])))
                    pMap[fmt.Sprint(i+1)] = []int{p1[0],y}
                } else {
                    mdx := math.Max(dx-1,0)
                    mdy := math.Max(dy-1,0)
                    x := p2[0] + int(math.Copysign(mdx,float64(p1[0]-p2[0])))
                    fmt.Println(x)
                    y := p2[1] + int(math.Copysign(mdy,float64(p1[1]-p2[1])))
                    pMap[fmt.Sprint(i+1)] = []int{x,y}
                }
            }
            
        }
    }

    return pMap
}

func plotMap(pMap map[string][]int, h,w int) {

    var grid [][]string
    for i:=0;i<h;i++ {
        var row []string
        for j:=0;j<w;j++ {
            row = append(row,". ")
        }
        grid = append(grid,row)
    }

    for i:=len(pMap)-1;i>0;i-- {
        x := pMap[fmt.Sprint(i)][0]
        y := pMap[fmt.Sprint(i)][1]

        grid[y][x] = fmt.Sprint(i," ")
    }

    x := pMap["h"][0]
    y := pMap["h"][1]
    grid[y][x] = fmt.Sprint("H ")

    for row:=len(grid)-1;row>=0;row-- {
        for _,el := range grid[row] {
            fmt.Print(el)
        }
        fmt.Println()
    }
    fmt.Println()
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

    pMap := make(map[string][]int)
	pMap["h"] = []int{0, 0}

    ropeLength := 10
    for i:=1;i<ropeLength;i++ { // length of rope
        pMap[fmt.Sprint(i)] = []int{0, 0}
    }

    var visitedList [][]int
    
    visitedList = append(visitedList,[]int{0, 0})
	
    for sc.Scan() {
		step := sc.Text()

        fmt.Println("==",step,"==")
		d, ms, _ := strings.Cut(step, " ")
		m, _ := strconv.Atoi(ms)
		switch d {
		case "U":
            for y:=1;y<=m;y++ {
                pMap["h"] = []int{pMap["h"][0],pMap["h"][1] + 1}
                pMap = updateTail(pMap)
                visitedList = append(visitedList,pMap[fmt.Sprint(ropeLength-1)])
            }
		case "D":
            for y:=m;y>0;y--{
                pMap["h"] = []int{pMap["h"][0],pMap["h"][1] - 1}
                pMap = updateTail(pMap)
                visitedList = append(visitedList,pMap[fmt.Sprint(ropeLength-1)])
            }
		case "L":
            for x:=m;x>0;x--{
                pMap["h"] = []int{pMap["h"][0] - 1,pMap["h"][1]}
                pMap = updateTail(pMap)
                visitedList = append(visitedList,pMap[fmt.Sprint(ropeLength-1)])
            }
		case "R":
            for x:=1;x<=m;x++{
                pMap["h"] = []int{pMap["h"][0] + 1,pMap["h"][1]}
                pMap = updateTail(pMap)
                visitedList = append(visitedList,pMap[fmt.Sprint(ropeLength-1)])
            }
		}
	}

    unique := make(map[string]struct{})
    for _,p := range visitedList {
        unique[fmt.Sprint(p[0],p[1])] = struct{}{}
    }

	fmt.Println(len(unique))
}
