package main

func getResourceValueAfter(grid landGrid, minutes int) (int, int, landGrid) {
	for i := 0; i < minutes; i++ {
		oldGrid := grid
		grid = landGrid{}
		for y, row := range oldGrid {
			for x, cell := range row {
				switch cell {
				case open:
					grid.assign(x, y, handleOpenArea(oldGrid, x, y))
				case trees:
					grid.assign(x, y, handleWoodedArea(oldGrid, x, y))
				case lumberyard:
					grid.assign(x, y, handleLumberyards(oldGrid, x, y))
				}
			}
		}
		if i%1000 == 0 {
			imageGrid(grid, "result", 10)
		}
	}
	counts := getTotalCounts(grid)
	return counts.trees, counts.lumberyards, grid
}

func handleOpenArea(grid landGrid, x, y int) landType {
	nCounts := getNeighborCounts(grid, x, y)
	if nCounts.trees >= 3 {
		return trees
	}
	return open
}

func handleWoodedArea(grid landGrid, x, y int) landType {
	nCounts := getNeighborCounts(grid, x, y)
	if nCounts.lumberyards >= 3 {
		return lumberyard
	}
	return trees
}

func handleLumberyards(grid landGrid, x, y int) landType {
	nCounts := getNeighborCounts(grid, x, y)
	if nCounts.lumberyards >= 1 && nCounts.trees >= 1 {
		return lumberyard
	}
	return open
}

type areaCounts struct {
	open        int
	trees       int
	lumberyards int
}

func getNeighborCounts(grid landGrid, x, y int) (nCounts areaCounts) {
	deltas := [3]int{-1, 0, 1}

	for _, dX := range deltas {
		for _, dY := range deltas {
			if dX != 0 || dY != 0 {
				switch grid.getType(x+dX, y+dY) {
				case open:
					nCounts.open++
				case trees:
					nCounts.trees++
				case lumberyard:
					nCounts.lumberyards++
				}
			}
		}
	}
	return
}

func getTotalCounts(grid landGrid) (ac areaCounts) {
	for _, row := range grid {
		for _, cell := range row {
			switch cell {
			case open:
				ac.open++
			case trees:
				ac.trees++
			case lumberyard:
				ac.lumberyards++
			}
		}
	}
	return
}
