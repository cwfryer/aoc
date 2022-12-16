package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"regexp"
	"strings"
)

type valve struct {
	idx     int
	name    string
	rate    int
	on      bool
	tunnels []*valve
	dists   []distance
}

type distance struct {
	v        *valve
	distance int
}

func parseInput(fn string) map[string]*valve {
	file, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	tunnelMap := make(map[string]*valve)
	for i, line := range lines {
		var name string
		var rate int
		fmt.Sscanf(line, "Valve %s has flow rate=%d", &name, &rate)
		tunnelMap[name] = &valve{idx: i, name: name, rate: rate, on: false}
	}
	for _, line := range lines {
		var tunnels []*valve

		var name, str string
		r := regexp.MustCompile(`Valve (?P<name>[A-Z][A-Z]) h.*valves? (?P<str>.*)`)
		res := r.FindStringSubmatch(line)
		name = res[1]
		str = res[2]
		tuns := strings.Split(str, ", ")
		for _, t := range tuns {
			v := tunnelMap[t]
			tunnels = append(tunnels, v)
		}
		v_ := tunnelMap[name]
		v_.tunnels = tunnels
		tunnelMap[name] = v_
	}

	return tunnelMap
}

var itermax = 10000000

func main() {
	vm := parseInput("day16/day16.input")
    max_pressure := 0
	for i := 0; i < itermax; i++ {
		if i%(itermax/100) == 0 {
			fmt.Println(i, "loops")
		}
		previous := []string{"AA", "AA"}
		current := []string{"AA", "AA"}
		opened := make(map[string]struct{})
		minutes := 26
		pressure := 0

		for m := minutes; m >= 0; m-- {
			for move := range []string{"me","elephant"} {
				valve := vm[current[move]]
				if valve.rate != 0 && !in(current[move], opened) && rand.Float64() > 0.15 {
					opened[valve.name] = struct{}{}
					pressure += int(math.Max(0, float64(m-1))) * valve.rate
				} else {
                    choices := make(map[string]struct{})
                    for _,t := range valve.tunnels {
                        choices[t.name] = struct{}{}
                    }
                    if rand.Float64() > 0.05 && in(previous[move],choices) && len(choices) > 1 {
                        delete(choices,previous[move])
                    }
                    if rand.Float64() > 0.20 && move == 1 && in(current[0],choices) && len(choices) > 1 {
                        delete(choices,previous[move])
                    }
                    previous[move] = current[move]
                    var cm string
                    r := rand.Intn(len(choices))
                    for k := range choices {
                        if r == 0 {
                            cm = k
                        }
                        r--
                    }
                    current[move] = cm
                }

			}
		}
        if pressure >= max_pressure {
            max_pressure = pressure
            fmt.Println(i, max_pressure)
        }
	}
}

func in(s string, m map[string]struct{}) bool {
	for e := range m {
		if s == e {
			return true
		}
	}
	return false
}
