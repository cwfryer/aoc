package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type blueprint struct {
	id         int
	oreRobot   oreRobot
	clayRobot  clayRobot
	obsRobot   obsRobot
	geodeRobot geodeRobot
}

type oreRobot struct {
	oreCost int
}

type clayRobot struct {
	oreCost int
}

type obsRobot struct {
	oreCost, clayCost int
}

type geodeRobot struct {
	oreCost, obsCost int
}

var globalBest = 0

func MaxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func main() {
	file, err := os.Open("day19.input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	blueprints := parseInput(file)
	fmt.Println("Part 1:", solvePart1(blueprints))
	fmt.Println("Part 2:", solvePart2(blueprints))
}

func parseInput(r io.Reader) []blueprint {
	scanner := bufio.NewScanner(r)
	blueprints := []blueprint{}

	for scanner.Scan() {
		var id int
		oreRobot := oreRobot{}
		clayRobot := clayRobot{}
		obsRobot := obsRobot{}
		geodeRobot := geodeRobot{}

		fmt.Sscanf(scanner.Text(),
			"Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.",
			&id, &oreRobot.oreCost, &clayRobot.oreCost, &obsRobot.oreCost, &obsRobot.clayCost, &geodeRobot.oreCost, &geodeRobot.obsCost)

		bp := blueprint{
			id,
			oreRobot,
			clayRobot,
			obsRobot,
			geodeRobot,
		}

		blueprints = append(blueprints, bp)
	}

	return blueprints
}

func solvePart1(blueprints []blueprint) int {
	result := 0
	for _, bp := range blueprints {

		result += bp.id * search(bp, 0, 0, 0, 24, 1, 0, 0, 0, 0)
		globalBest = 0
	}

	return result
}

func search(bp blueprint, ore, clay, obs, time, oreRobots, clayRobots, obsRobots, geodeRobots, geodes int) int {
	if time == 0 || globalBest >= geodes+rangeSum(geodeRobots, geodeRobots+time-1) {
		return 0
	}
	if oreRobots >= bp.geodeRobot.oreCost && obsRobots >= bp.geodeRobot.obsCost {
		return rangeSum(geodeRobots, geodeRobots+time-1)
	}

	oreLimitHit := oreRobots >= MaxInt(bp.geodeRobot.oreCost, MaxInt(bp.clayRobot.oreCost, bp.obsRobot.oreCost))
	clayLimitHit := clayRobots >= bp.obsRobot.clayCost
	obsLimitHit := obsRobots >= bp.geodeRobot.obsCost
	best := 0

	if !oreLimitHit {
		best = MaxInt(
			best,
			geodeRobots+search(
				bp, ore+oreRobots, clay+clayRobots, obs+obsRobots,
				time-1, oreRobots, clayRobots, obsRobots, geodeRobots, geodes+geodeRobots))
	}
	if ore >= bp.oreRobot.oreCost && !oreLimitHit {
		best = MaxInt(
			best,
			geodeRobots+search(
				bp, ore-bp.oreRobot.oreCost+oreRobots, clay+clayRobots, obs+obsRobots,
				time-1, oreRobots+1, clayRobots, obsRobots, geodeRobots, geodes+geodeRobots))
	}
	if ore >= bp.clayRobot.oreCost && !clayLimitHit {
		best = MaxInt(
			best, geodeRobots+search(
				bp, ore-bp.clayRobot.oreCost+oreRobots, clay+clayRobots, obs+obsRobots,
				time-1, oreRobots, clayRobots+1, obsRobots, geodeRobots, geodes+geodeRobots))
	}
	if ore >= bp.obsRobot.oreCost && clay >= bp.obsRobot.clayCost && !obsLimitHit {
		best = MaxInt(
			best, geodeRobots+search(
				bp, ore-bp.obsRobot.oreCost+oreRobots, clay-bp.obsRobot.clayCost+clayRobots, obs+obsRobots,
				time-1, oreRobots, clayRobots, obsRobots+1, geodeRobots, geodes+geodeRobots))
	}
	if ore >= bp.geodeRobot.oreCost && obs >= bp.geodeRobot.obsCost {
		best = MaxInt(
			best, geodeRobots+search(
				bp, ore-bp.geodeRobot.oreCost+oreRobots, clay+clayRobots, obs-bp.geodeRobot.obsCost+obsRobots,
				time-1, oreRobots, clayRobots, obsRobots, geodeRobots+1, geodes+geodeRobots))
	}

	globalBest = MaxInt(best, globalBest)
	return best
}

func rangeSum(first, last int) int {
	return last*(last+1)/2 - ((first - 1) * first / 2)
}

func solvePart2(blueprints []blueprint) int {
	if len(blueprints) < 3 {
		return -1
	}
	result := 1
	for i := 0; i < 3; i++ {
		result *= search(blueprints[i], 0, 0, 0, 32, 1, 0, 0, 0, 0)
		globalBest = 0
	}

	return result
}
