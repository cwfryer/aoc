package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type cost struct {
	ore      int
	clay     int
	obsidian int
}

func (c cost) Minus(x cost) cost {
	out := cost{
		ore:      c.ore - x.ore,
		clay:     c.clay - x.clay,
		obsidian: c.obsidian - x.obsidian,
	}
	return out
}

type blueprint struct {
	oreRobotCost  cost
	clayRobotCost cost
	obsRobotCost  cost
	geoRobotCost  cost
}

type State struct {
	bp         *blueprint
	score      int
	Inventory  cost
	Production cost
	turn       int
	maxturns   int

	parents []*State
}

func (s *State) Print() {
	fmt.Println("---------------------------")
	fmt.Println("score:", s.score)
	fmt.Println("inventory:", s.Inventory)
	fmt.Println("production", s.Production)
	fmt.Println("turn:", s.turn)
	fmt.Println("remaining turns:", s.maxturns-s.turn)
}

func main() {
	bp := parseInput("day19/day19.input")
    answer := 0
	for i, b := range bp {
        fmt.Println()
        fmt.Println("###########################")
		fmt.Println("Blueprint", i+1)
		initState := &State{
			bp:         b,
			score:      0,
			Inventory:  cost{0, 0, 0},
			Production: cost{1, 0, 0},
			turn:       0,
			maxturns:   24,
			parents:    []*State{},
		}
		// initState.Print()
		// opts := initState.getOptions()
		// opts[0].Print()
		max := maxScore(initState)
		fmt.Println("score:", max)
        answer += (i+1)*max
	}
    fmt.Println(answer)
}

func parseInput(in string) []*blueprint {
	input, _ := os.ReadFile(in)
	split := strings.Split(strings.TrimSpace(string(input)), "\n")

	blueprints := make([]*blueprint, len(split))
	for _, s := range split {
		var name, oreOreCost, clayOreCost, obsidianOreCost, obsidianClayCost, geodeOreCost, geodeObsidianCost int
		fmt.Sscanf(s,
			`Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.`,
			&name, &oreOreCost, &clayOreCost, &obsidianOreCost, &obsidianClayCost, &geodeOreCost, &geodeObsidianCost)

		blueprints[name-1] = &blueprint{
			oreRobotCost:  cost{ore: oreOreCost},
			clayRobotCost: cost{ore: clayOreCost},
			obsRobotCost:  cost{ore: obsidianOreCost, clay: obsidianClayCost},
			geoRobotCost:  cost{ore: geodeOreCost, obsidian: geodeObsidianCost},
		}
	}

	return blueprints
}

func maxScore(s *State) int {
	max_score := 0
	queue := []*State{}
	queue = append(queue, s)
	return bfs(queue, max_score)
}

func bfs(queue []*State, score int) int {
	if len(queue) == 0 {
		return score
	}
    // if queue[0].score > score {
    //     fmt.Println("===========================")
    //     for _,s := range queue[0].parents {
    //         s.Print()
    //     }
    //     queue[0].Print()
    // }
	score = func(l, r int) int {
		if r > l {
			return r
		} else {
			return l
		}
	}(score, queue[0].score)

	for _, o := range queue[0].getOptions() {
		queue = append(queue, o)
	}

	return bfs(queue[1:], score)
}

func (s *State) getOptions() []*State {
	if s.turn > s.maxturns {
		return []*State{}
	}
	options := []*State{}
	// check if i have the production necessary to make this robot
	canMake := func(s *State, robot string) bool {
		out := true
		var reqs cost
		switch robot {
		case "ore":
			reqs = s.bp.oreRobotCost
		case "clay":
			reqs = s.bp.clayRobotCost
		case "obsidian":
			reqs = s.bp.obsRobotCost
		case "geode":
			reqs = s.bp.geoRobotCost
		}
		rq := []int{reqs.ore, reqs.clay, reqs.obsidian}
		pd := []int{s.Production.ore, s.Production.clay, s.Production.obsidian}
		for i, r := range rq {
			if r > 0 {
				if pd[i] == 0 {
					out = false
					break
				}
			}
		}
		return out
	}
	// check if i have the time necessary to make this robot
	haveTime := func(s *State, robot string) (bool, int) {
		var need cost
		switch robot {
		case "ore":
			need = s.bp.oreRobotCost.Minus(s.Inventory)
		case "clay":
			need = s.bp.clayRobotCost.Minus(s.Inventory)
		case "obsidian":
			need = s.bp.obsRobotCost.Minus(s.Inventory)
		case "geode":
			need = s.bp.geoRobotCost.Minus(s.Inventory)
		}
		needturns := 1
		turns := []int{}
		if need.ore > 0 {
			turns = append(turns, (need.ore/s.Production.ore)+(need.ore%s.Production.ore))
		}
		if need.clay > 0 {
			turns = append(turns, (need.clay/s.Production.clay)+(need.clay%s.Production.clay))
		}
		if need.obsidian > 0 {
			turns = append(turns, (need.obsidian/s.Production.obsidian)+(need.obsidian%s.Production.obsidian))
		}
		if len(turns) > 0 {
			sort.Ints(turns)
			needturns += turns[len(turns)-1]
		}
		if needturns > s.maxturns-s.turn {
			return false, needturns
		}
		return true, needturns
	}
	// check if i even need to make this robot
	needMore := func(s *State, robot string) bool {
		switch robot {
		case "ore":
			oreCosts := []int{s.bp.oreRobotCost.ore, s.bp.clayRobotCost.ore, s.bp.obsRobotCost.ore, s.bp.geoRobotCost.ore}
			sort.Ints(oreCosts)
			maxneed := oreCosts[len(oreCosts)-1]
			if rt := s.maxturns - s.turn; (s.Production.ore*rt)+s.Inventory.ore >= (maxneed * rt) {
				return false
			}
		case "clay":
			clayCosts := []int{s.bp.oreRobotCost.clay, s.bp.clayRobotCost.clay, s.bp.obsRobotCost.clay, s.bp.geoRobotCost.clay}
			sort.Ints(clayCosts)
			maxneed := clayCosts[len(clayCosts)-1]
			if rt := s.maxturns - s.turn; (s.Production.clay*rt)+s.Inventory.clay >= (maxneed * rt) {
				return false
			}
		case "obsidian":
			obsCosts := []int{s.bp.oreRobotCost.obsidian, s.bp.clayRobotCost.obsidian, s.bp.obsRobotCost.obsidian, s.bp.geoRobotCost.obsidian}
			sort.Ints(obsCosts)
			maxneed := obsCosts[len(obsCosts)-1]
			if rt := s.maxturns - s.turn; (s.Production.obsidian*rt)+s.Inventory.obsidian >= (maxneed * rt) {
				return false
			}
		}
		return true
	}

	// ORE ROBOT
	oreCM := canMake(s, "ore")
	if oreCM {
		oreHT, needturns := haveTime(s, "ore")
		oreNM := needMore(s, "ore")
		if oreHT && oreNM {
			updateInventory := cost{
				ore:      s.Inventory.ore + (s.Production.ore * needturns) - s.bp.oreRobotCost.ore,
				clay:     s.Inventory.clay + (s.Production.clay * needturns) - s.bp.oreRobotCost.clay,
				obsidian: s.Inventory.obsidian + (s.Production.obsidian * needturns) - s.bp.oreRobotCost.obsidian,
			}
			if updateInventory.ore < 0 || updateInventory.clay < 0 || updateInventory.obsidian < 0 {
				panic(fmt.Sprintf("ore fail\ninv:\t{%d,%d,%d}\nprod:\t{%d,%d,%d}\nturns:\t%d\ncost:\t{%d,%d,%d}\nupdate:\t{%d,%d,%d}",
					s.Inventory.ore, s.Inventory.clay, s.Inventory.obsidian,
					s.Production.ore, s.Production.clay, s.Production.obsidian,
					needturns,
					s.bp.clayRobotCost.ore, s.bp.clayRobotCost.clay, s.bp.clayRobotCost.obsidian,
					updateInventory.ore, updateInventory.clay, updateInventory.obsidian))
			}
			newState := &State{
				bp:         s.bp,
				score:      s.score,
				Inventory:  updateInventory,
				Production: cost{ore: s.Production.ore + 1, clay: s.Production.clay, obsidian: s.Production.obsidian},
				turn:       s.turn + needturns,
				maxturns:   s.maxturns,
				parents:    append(s.parents, s),
			}
			options = append(options, newState)
		}
	}

	// CLAY ROBOT
	clayCM := canMake(s, "clay")
	if clayCM {
		clayHT, needturns := haveTime(s, "clay")
		clayNM := needMore(s, "clay")

		if clayHT && clayNM {
			updateInventory := cost{
				ore:      s.Inventory.ore + (s.Production.ore * needturns) - s.bp.clayRobotCost.ore,
				clay:     s.Inventory.clay + (s.Production.clay * needturns) - s.bp.clayRobotCost.clay,
				obsidian: s.Inventory.obsidian + (s.Production.obsidian * needturns) - s.bp.clayRobotCost.obsidian,
			}
			if updateInventory.ore < 0 || updateInventory.clay < 0 || updateInventory.obsidian < 0 {
				panic(fmt.Sprintf("clay fail\ninv:\t{%d,%d,%d}\nprod:\t{%d,%d,%d}\nturns:\t%d\ncost:\t{%d,%d,%d}\nupdate:\t{%d,%d,%d}",
					s.Inventory.ore, s.Inventory.clay, s.Inventory.obsidian,
					s.Production.ore, s.Production.clay, s.Production.obsidian,
					needturns,
					s.bp.clayRobotCost.ore, s.bp.clayRobotCost.clay, s.bp.clayRobotCost.obsidian,
					updateInventory.ore, updateInventory.clay, updateInventory.obsidian))
			}
			newState := &State{
				bp:         s.bp,
				score:      s.score,
				Inventory:  updateInventory,
				Production: cost{ore: s.Production.ore, clay: s.Production.clay + 1, obsidian: s.Production.obsidian},
				turn:       s.turn + needturns,
				maxturns:   s.maxturns,
				parents:    append(s.parents, s),
			}
			options = append(options, newState)
		}
	}

	// OBSIDIAN ROBOT
	obsCM := canMake(s, "obsidian")
	if obsCM {
		obsHT, needturns := haveTime(s, "obsidian")
		obsNM := needMore(s, "obsidian")

		if obsHT && obsNM {
			updateInventory := cost{
				ore:      s.Inventory.ore + (s.Production.ore * needturns) - s.bp.obsRobotCost.ore,
				clay:     s.Inventory.clay + (s.Production.clay * needturns) - s.bp.obsRobotCost.clay,
				obsidian: s.Inventory.obsidian + (s.Production.obsidian * needturns) - s.bp.obsRobotCost.obsidian,
			}
			if updateInventory.ore < 0 || updateInventory.clay < 0 || updateInventory.obsidian < 0 {
				panic(fmt.Sprintf("obs fail\ninv:\t{%d,%d,%d}\nprod:\t{%d,%d,%d}\nturns:\t%d\ncost:\t{%d,%d,%d}\nupdate:\t{%d,%d,%d}",
					s.Inventory.ore, s.Inventory.clay, s.Inventory.obsidian,
					s.Production.ore, s.Production.clay, s.Production.obsidian,
					needturns,
					s.bp.clayRobotCost.ore, s.bp.clayRobotCost.clay, s.bp.clayRobotCost.obsidian,
					updateInventory.ore, updateInventory.clay, updateInventory.obsidian))
			}
			newState := &State{
				bp:         s.bp,
				score:      s.score,
				Inventory:  updateInventory,
				Production: cost{ore: s.Production.ore, clay: s.Production.clay, obsidian: s.Production.obsidian + 1},
				turn:       s.turn + needturns,
				maxturns:   s.maxturns,
				parents:    append(s.parents, s),
			}
			options = append(options, newState)
		}
	}
	// GEODE ROBOT
	geoCM := canMake(s, "geode")
	if geoCM {
		geoHT, needturns := haveTime(s, "geode")
		if geoHT {
			updateInventory := cost{
				ore:      s.Inventory.ore + (s.Production.ore * needturns) - s.bp.geoRobotCost.ore,
				clay:     s.Inventory.clay + (s.Production.clay * needturns) - s.bp.geoRobotCost.clay,
				obsidian: s.Inventory.obsidian + (s.Production.obsidian * needturns) - s.bp.geoRobotCost.obsidian,
			}
			if updateInventory.ore < 0 || updateInventory.clay < 0 || updateInventory.obsidian < 0 {
				panic(fmt.Sprintf("geo fail\ninv:\t{%d,%d,%d}\nprod:\t{%d,%d,%d}\nturns:\t%d\ncost:\t{%d,%d,%d}\nupdate:\t{%d,%d,%d}",
					s.Inventory.ore, s.Inventory.clay, s.Inventory.obsidian,
					s.Production.ore, s.Production.clay, s.Production.obsidian,
					needturns,
					s.bp.clayRobotCost.ore, s.bp.clayRobotCost.clay, s.bp.clayRobotCost.obsidian,
					updateInventory.ore, updateInventory.clay, updateInventory.obsidian))
			}
			newState := &State{
				bp:         s.bp,
				score:      s.score + (s.maxturns - (s.turn+needturns)),
				Inventory:  updateInventory,
				Production: cost{ore: s.Production.ore, clay: s.Production.clay, obsidian: s.Production.obsidian},
				turn:       s.turn + needturns,
				maxturns:   s.maxturns,
				parents:    append(s.parents, s),
			}
			options = append(options, newState)
		}
	}
	return options
}
