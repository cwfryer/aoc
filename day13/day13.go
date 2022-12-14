package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type lol struct {
	nodeType string
	rep      string
	value    int
	children []*lol
	special  int
}

func parseList(in string) *lol {
	n := &lol{nodeType: "list", children: []*lol{}, rep: in}
	for i := 1; i < len(in); i++ {
		if in[i] == '[' {
			open := 1
			for j := i + 1; j < len(in); j++ {
				if in[j] == '[' {
					open++
				}
				if in[j] == ']' {
					open--
				}
				if open == 0 {
					n.children = append(n.children, parseList(in[i:j+1]))
					i = j
					break
				}
			}
		} else if in[i] == ',' {
			continue
		} else {
			for j := i + 1; j < len(in); j++ {
				if in[j] == ',' || in[j] == ']' {
					n.children = append(n.children, parseNumber(in[i:j]))
					i = j
					break
				}
			}
		}
	}
	return n
}

func parseNumber(in string) *lol {
	n := &lol{nodeType: "number", rep: in}
	n.value, _ = strconv.Atoi(in)
	return n
}

func cmp(x, y int) int {
	if x > y {
		return -1
	}
	if x < y {
		return 1
	}
	return 0
}
func compareLists(l1 *lol, l2 *lol) int {
	if l1.nodeType == "number" && l2.nodeType == "number" {
		return cmp(l1.value, l2.value)
	}
	if l1.nodeType == "number" {
		l1.nodeType = "list"
		l1.children = []*lol{{nodeType: "number", value: l1.value}}
	}
	if l2.nodeType == "number" {
		l2.nodeType = "list"
		l2.children = []*lol{{nodeType: "number", value: l2.value}}
	}

	mn := int(math.Min(float64(len(l1.children)), float64(len(l2.children))))
	for i := 0; i < mn; i++ {
		check := compareLists(l1.children[i], l2.children[i])
		if check == 0 {
			continue
		}
		return check
	}
	return cmp(len(l1.children), len(l2.children))
}
func main() {
	input, _ := os.ReadFile("day13/day13.input")
	split := strings.Split(strings.TrimSpace(string(input)), "\n\n")

	var pairs []string
	for _, x := range split {
		l, r, _ := strings.Cut(x, "\n")
		pairs = append(pairs, l)
		pairs = append(pairs, r)
	}
    pairs = append(pairs,"[[2]]")
    pairs = append(pairs,"[[6]]")

    var nodes []*lol
    for _,p := range pairs {
        nodes = append(nodes,parseList(p))
    }

    sorted := make([]*lol,0 ,len(nodes))
    for i,l := range nodes {
        if i == 0 {
            sorted = append(sorted,nodes[0])
            continue
        }
        
        for j,x := range sorted {
            check := compareLists(l,x)
            if check == 1 {
                if len(sorted) == 1 {
                    sorted = append([]*lol{l},sorted...)
                } else {
                    tmp := append([]*lol{},sorted[:j]...)
                    tmp = append(tmp,l)
                    sorted = append(tmp,sorted[j:]...)
                }
                break
            }
        }
        if check := compareLists(l,sorted[len(sorted)-1]); check == -1 {
            sorted = append(sorted,l)
        }

    }



    a := 1
    for i,x := range sorted {
        fmt.Println(x.rep)
        if x.rep == "[[2]]" || x.rep == "[[6]]" {
            a *= (i+1)
        }
    }
    fmt.Println(a)


}
