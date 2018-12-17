package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/benaychh/aocutils"
)

func main() {
	initialState, notes := getPuzzleInput("day12.data")
	start := time.Now().UnixNano()
	part1 := getGeneration(initialState, notes, 20)
	end := time.Now().UnixNano()
	fmt.Printf("Part 1: %v, Time: %fms\n", part1, float32(end-start)/1000.0/1000.0)

	start = time.Now().UnixNano()
	startGeneration := 49999999998
	diffrenceBetweenGenerationAndStart := 35
	stableString := "#.##.##.##.##.##.##.##.##.##.##.##.##.##.##.##.##.##.##.##.##.##.##.##.##.##.##.##.##.##.##.##.##.##.##.##.##.##.##.##.##.##.#...#.#"
	part2 := getGenerationOptimized(stableString, startGeneration-diffrenceBetweenGenerationAndStart, startGeneration, notes, 50e9)
	end = time.Now().UnixNano()
	fmt.Printf("Part 2: %v, Time: %fms\n", part2, float32(end-start)/1000.0/1000.0)
}

func getPuzzleInput(filename string) (initialState string, notes map[string]rune) {
	rawData := aocutils.LoadFileAsStringArray(filename)
	notes = map[string]rune{}
	initialState = strings.Split(rawData[0], ": ")[1]

	for i := 2; i < len(rawData); i++ {
		split := strings.Split(rawData[i], " => ")
		notes[split[0]] = rune(split[1][0])
	}
	return initialState, notes
}
