package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	name      string
	mode      string
	number    int
	op        func(int, int) int
	inverseOp func(int, int) int
	opString  string
	children  []string
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	monkeys := make(map[string]Monkey)
	for sc.Scan() {
		line := sc.Text()
		parts := strings.Fields(line)
		m := Monkey{name: parts[0][:4]}
		if len(parts) > 2 {
			m.mode = "op"
			m.op = map[string]func(int, int) int{
				"+": func(m1, m2 int) int { return m1 + m2 },
				"-": func(m1, m2 int) int { return m1 - m2 },
				"*": func(m1, m2 int) int { return m1 * m2 },
				"/": func(m1, m2 int) int { return m1 / m2 },
			}[parts[2]]
            m.inverseOp = map[string]func(int, int) int{
				"+": func(m1, m2 int) int { return m2 - m1 },
				"-": func(m1, m2 int) int { return m1 + m2 },
				"*": func(m1, m2 int) int { return m2 / m1 },
				"/": func(m1, m2 int) int { return m1 * m2 },
			}[parts[2]]
			m.opString = parts[2]
			m.children = []string{parts[1], parts[3]}
		} else {
			m.mode = "int"
			m.number, _ = strconv.Atoi(parts[1])
		}
		monkeys[m.name] = m
	}

    var target int
    for _,c := range monkeys["root"].children {
        if !findHuman(monkeys[c],monkeys) {
            target = getValue(monkeys[c],monkeys)
        }
    }
    for _,c := range monkeys["root"].children {
        if findHuman(monkeys[c],monkeys) {
            fmt.Println(inverse(target,monkeys[c],monkeys))
        }
    }

}

func getValue(m Monkey, monkeyMap map[string]Monkey) int {
	if m.mode == "int" {
		return m.number
	} else {
		return m.op(
            getValue(monkeyMap[m.children[0]], monkeyMap),
            getValue(monkeyMap[m.children[1]], monkeyMap),
        )
    }
}

func findHuman(m Monkey, monkeyMap map[string]Monkey) bool {
    if m.name == "humn" {
        return true
    }
    if len(m.children) == 0 {
        return false
    }
    m1 := monkeyMap[m.children[0]]
    m2 := monkeyMap[m.children[1]]
    return findHuman(m1,monkeyMap) || findHuman(m2,monkeyMap)
}

func inverse(target int, m Monkey, monkeyMap map[string]Monkey) int {
    if m.name == "humn" {
        return target
    }
    if len(m.children) == 0 {
        return target - m.number
    }
    m1 := monkeyMap[m.children[0]]
    m2 := monkeyMap[m.children[1]]
    switch m.opString {
    case "/":
        if findHuman(m1,monkeyMap) {
            target *= getValue(m2,monkeyMap)
            return inverse(target,m1,monkeyMap)
        }
        if findHuman(m2,monkeyMap) {
            target = getValue(m1,monkeyMap) / target
            return inverse(target,m2,monkeyMap)
        }
        panic("neither side has human, this shouldn't be possible")
    case "*":
        if findHuman(m1,monkeyMap){
            target /= getValue(m2,monkeyMap)
            return inverse(target,m1,monkeyMap)
        }
        if findHuman(m2,monkeyMap){
            target /= getValue(m1,monkeyMap)
            return inverse(target,m2,monkeyMap)
        }
        panic("neither side has human, this shouldn't be possible")
    case "-":
        if findHuman(m1,monkeyMap) {
            target += getValue(m2,monkeyMap)
            return inverse(target,m1,monkeyMap)
        }
        if findHuman(m2,monkeyMap) {
            target = getValue(m1,monkeyMap) - target
            return inverse(target,m2,monkeyMap)
        }
        panic("neither side has human, this shouldn't be possible")
    default:
        if findHuman(m1,monkeyMap) {
            target -= getValue(m2,monkeyMap)
            return inverse(target,m1,monkeyMap)
        }
        if findHuman(m2,monkeyMap) {
            target -= getValue(m1,monkeyMap)
            return inverse(target,m2,monkeyMap)
        }
        panic("neither side has human, this shouldn't be possible")
    }
    // Should be unreachable
    panic("this should never be reachable")
}

