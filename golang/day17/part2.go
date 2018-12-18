package main

func countSettledWater(grid undergroundGrid) (waterCount int) {
	sortedRows := grid.sortedKeys()
	hasClay := false
	for _, y := range sortedRows[0 : len(sortedRows)-1] {
		row := grid[y]
		localCount := 0
		for _, cell := range row {
			if cell == settledWater {
				localCount++
			} else if cell == clay {
				hasClay = true
			}
		}
		if hasClay {
			waterCount += localCount
		}
	}
	return
}
