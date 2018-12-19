package main

import (
	"fmt"
	"math"
)

func getResourceValueAfterBigNumber(grid landGrid, minutes int) (int, int, landGrid) {
	counts := map[string]int{}
	lastRepeat := 0
	lastDifference := 0
	var ac areaCounts
	for i := 0; i < minutes; i++ {
		ac = areaCounts{}
		oldGrid := grid
		grid = landGrid{}
		for y, row := range oldGrid {
			for x, cell := range row {
				handleCell(grid, oldGrid, &ac, cell, x, y)
			}
		}
		key := fmt.Sprintf("%d,%d,%d", ac.open, ac.trees, ac.lumberyards)
		count, ok := counts[key]
		if ok {
			if i-count == lastDifference && i == lastRepeat+1 {
				i = calculateHyperSpeedI(minutes, lastDifference, i)
			} else {
				lastDifference = i - count
				lastRepeat = i
			}
		} else {
			counts[key] = i
		}
	}
	return ac.trees, ac.lumberyards, grid
}

func handleCell(grid landGrid, oldGrid landGrid, ac *areaCounts, cell landType, x, y int) {
	var lt landType
	switch cell {
	case open:
		lt = handleOpenArea(oldGrid, x, y)
	case trees:
		lt = handleWoodedArea(oldGrid, x, y)
	case lumberyard:
		lt = handleLumberyards(oldGrid, x, y)
	}
	switch lt {
	case open:
		ac.open++
	case trees:
		ac.trees++
	case lumberyard:
		ac.lumberyards++
	}
	grid.assign(x, y, lt)
}

func calculateHyperSpeedI(target, slope, current int) int {
	steps := int(math.Floor(float64(target-1-current) / float64(slope)))
	return current + slope*steps
}
