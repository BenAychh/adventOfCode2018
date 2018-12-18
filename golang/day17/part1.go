package main

func countWater(grid undergroundGrid, x, y int) (waterCount int) {
	count := 0
	waterMovesDown(grid, 500, 0, &count)
	sortedRows := grid.sortedKeys()
	hasClay := false
	for _, y := range sortedRows[0 : len(sortedRows)-1] {
		row := grid[y]
		localCount := 0
		for _, cell := range row {
			if cell == flowingWater || cell == settledWater {
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
