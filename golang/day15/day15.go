package main

import (
	"aocutils"
	"fmt"
	"strings"
	"time"
)

func main() {
	data := aocutils.LoadArrayOfStringsFromTextFile("day15.data")
	grid, entities := processMap(data)
	start := time.Now().UnixNano()
	part1Rounds, part1Hitpoints := getRoundsAndRemainingHitpoints(grid, entities, true)
	end := time.Now().UnixNano()
	fmt.Printf("Part 1: %d x %d = %d, Time: %fms\n", part1Rounds, part1Hitpoints, part1Rounds*part1Hitpoints, float32(end-start)/1000.0/1000.0)
}

func processMap(lines []string) (grid [][]bool, entities entityLocations) {
	entities = entityLocations{}
	grid = [][]bool{}
	for y, line := range lines {
		row := []bool{}
		for x, r := range line {
			switch r {
			case '.':
				row = append(row, true)
			case 'G':
				row = append(row, true)
				entities.placeAt(x, y, &entity{
					attackPower: 3,
					hitPoints:   200,
					species:     goblin,
					enemy:       elf,
				})
			case 'E':
				row = append(row, true)
				entities.placeAt(x, y, &entity{
					attackPower: 3,
					hitPoints:   200,
					species:     elf,
					enemy:       goblin,
				})
			default:
				row = append(row, false)
			}
		}
		grid = append(grid, row)
	}
	return
}

func getPrintableString(grid [][]bool, entities entityLocations) string {
	var sb strings.Builder
	for y, row := range grid {
		for x, cell := range row {
			if entities.speciesAtLocation(x, y, elf) {
				sb.WriteRune('E')
			} else if entities.speciesAtLocation(x, y, goblin) {
				sb.WriteRune('G')
			} else if cell {
				sb.WriteRune('.')
			} else {
				sb.WriteRune('#')
			}
		}
		rowOfEntities, ok := entities[y]
		if ok {
			sb.WriteString("   ")
			for _, eX := range rowOfEntities.sortedKeys() {
				e := rowOfEntities[eX]
				initial := "E"
				if e.species == goblin {
					initial = "G"
				}
				sb.WriteString(fmt.Sprintf("%v(%d), ", initial, e.hitPoints))
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}
