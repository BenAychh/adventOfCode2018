package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func getGeneration(initialState string, notes map[string]rune, targetGeneration int) (sum int) {
	currentGeneration := map[int]rune{}
	for i, r := range initialState {
		currentGeneration[i] = r
	}
	for i := 0; i < targetGeneration; i++ {
		printPadded(currentGeneration)
		currentGeneration, _ = getNextGeneration(currentGeneration, notes)
	}

	for key, value := range currentGeneration {
		if value == '#' {
			sum += key
		}
	}

	return
}

func printPadded(generation map[int]rune) {
	var sb strings.Builder
	keys := []int{}
	for key := range generation {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	for i := -3; i < keys[0]; i++ {
		sb.WriteString(".")
	}
	for _, key := range keys {
		sb.WriteRune(generation[key])
	}
	for i := keys[len(keys)-1]; i < 35; i++ {
		sb.WriteString(".")
	}
	fmt.Println(sb.String())
}

func get5Pots(start int, pots map[int]rune) string {
	var sb strings.Builder
	for i := start; i < start+5; i++ {
		plant, ok := pots[i]
		if ok {
			sb.WriteRune(plant)
		} else {
			sb.WriteString(".")
		}
	}
	return sb.String()
}

func trimPots(pots map[int]rune) {
	keys := getSortedKeys(pots)
	firstPlant := math.MinInt64
	lastPlant := 0
	for _, key := range keys {
		if pots[key] == '#' {
			if firstPlant == math.MinInt64 {
				firstPlant = key
			}
			lastPlant = key
		}
	}

	for i := keys[0]; i < firstPlant; i++ {
		delete(pots, i)
	}
	for i := lastPlant + 1; i <= keys[len(keys)-1]; i++ {
		delete(pots, i)
	}
}

func getSortedKeys(pots map[int]rune) []int {
	keys := []int{}
	for key := range pots {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	return keys
}

func getNextGeneration(generation map[int]rune, notes map[string]rune) (map[int]rune, int) {
	newGeneration := map[int]rune{}
	sortedKeys := getSortedKeys(generation)
	min := sortedKeys[0] - 4
	max := sortedKeys[len(sortedKeys)-1] + 4
	for i := min; i < max; i++ {
		fivePots := get5Pots(i, generation)
		plant, hasNote := notes[fivePots]
		if hasNote {
			newGeneration[i+2] = plant
		} else {
			newGeneration[i+2] = '.'
		}
	}
	trimPots(newGeneration)

	return newGeneration, sortedKeys[0]
}
