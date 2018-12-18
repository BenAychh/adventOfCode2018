package main

func waterMovesDown(grid undergroundGrid, x, y int, loopCount *int) {
	for grid.getType(x, y+1) == sand {
		*loopCount++
		j := y + 1
		for ; grid.hasType(x, j) && (grid.getType(x, j) == sand); j++ {
		}
		if grid.hasType(x, j) {
			if grid.getType(x, j) == flowingWater {
				if grid.getType(x, j-1) != flowingWater && grid.getType(x, j-1) != spring {
					grid.assign(x, j-1, flowingWater)
				}
			} else {
				waterSpreadsOut(grid, x, j-1, loopCount)
			}
		} else {
			grid.assign(x, j-1, flowingWater)
		}
	}
}

func waterSpreadsOut(grid undergroundGrid, x, y int, loopCount *int) {
	leftBound := waterSpreadsLeft(grid, x, y, loopCount)
	rightBound := waterSpreadsRight(grid, x, y, loopCount)
	waterType := settledWater
	if grid.getType(leftBound, y+1) == flowingWater || grid.getType(rightBound, y+1) == flowingWater {
		waterType = flowingWater
	}
	for i := leftBound; i <= rightBound; i++ {
		if grid.hasType(i, y) && grid.getType(i, y) != clay {
			grid.assign(i, y, waterType)
		}
	}
}

func waterSpreadsLeft(grid undergroundGrid, x, y int, loopCount *int) (bound int) {
	bound = x
	for ; (grid.getType(bound, y) == sand) && grid.hasType(bound, y+1) && grid.getType(bound, y+1) != flowingWater; bound-- {
		if grid.hasType(bound, y+1) && grid.getType(bound, y+1) == sand {
			waterMovesDown(grid, bound, y, loopCount)
			bound++
		}
	}
	return
}

func waterSpreadsRight(grid undergroundGrid, x, y int, loopCount *int) (bound int) {
	bound = x
	for ; (grid.getType(bound, y) == sand) && grid.hasType(bound, y+1) && grid.getType(bound, y+1) != flowingWater; bound++ {
		if grid.hasType(bound, y+1) && grid.getType(bound, y+1) == sand {
			waterMovesDown(grid, bound, y, loopCount)
			bound--
		}
	}
	return
}
