package main

import (
	"fmt"
	"os"
	"strings"
)

type cost struct {
	ore      int
	clay     int
	obsidian int
}
type blueprint struct {
	oreRobotCost  cost
	clayRobotCost cost
	obsRobotCost  cost
	geoRobotCost  cost
}

func (b blueprint) estimateNeeds(turns int) {

}

func main() {
	input, _ := os.ReadFile("day19/day19.test")
	split := strings.Split(strings.TrimSpace(string(input)), "\n")

	blueprints := make([]blueprint, len(split))
	for _, s := range split {
		var name, oreOreCost, clayOreCost, obsidianOreCost, obsidianClayCost, geodeOreCost, geodeObsidianCost int
		fmt.Sscanf(s,
			`Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.`,
			&name, &oreOreCost, &clayOreCost, &obsidianOreCost, &obsidianClayCost, &geodeOreCost, &geodeObsidianCost)

		blueprints[name-1] = blueprint{
			oreRobotCost:  cost{ore: oreOreCost},
			clayRobotCost: cost{ore: clayOreCost},
			obsRobotCost:  cost{ore: obsidianOreCost, clay: obsidianClayCost},
			geoRobotCost:  cost{ore: geodeOreCost, obsidian: geodeObsidianCost},
		}
	}

	var MaxProducer int = 0
	currentMax := 0
	for i, b := range blueprints {
		max := runSimulation(24, b)
		fmt.Println(max)
		if max > currentMax {
			MaxProducer = i + 1
		}
	}
	fmt.Println(MaxProducer)
}

type state struct {
	blueprint  blueprint
	production cost
	inventory  cost
	oreRobots  int
	clayRobots int
	obsRobots  int
	score      int
}

func runSimulation(turns int, b blueprint) int {
	fmt.Println(b)
	oreProduction := 1
	clayProduction := 0
	obsProduction := 0
	geoProduction := 0
	oreFactory, clayFactory, obsFactory, geoFactory := 0, 0, 0, 0
	var oreInventory, clayInventory, obsInventory, geoInventory int
	for t := 1; t <= turns; t++ {
		oreInventory += oreProduction
		clayInventory += clayProduction
		obsInventory += obsProduction
		geoInventory += geoProduction

		if oreFactory > 0 {
			oreProduction++
			oreFactory--
		}
		if clayFactory > 0 {
			clayProduction++
			clayFactory--
		}
		if obsFactory > 0 {
			obsProduction++
			obsFactory--
		}
		if geoFactory > 0 {
			geoProduction++
			geoFactory--
		}

	}

	return geoInventory
}
