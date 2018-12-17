package main

import (
	"fmt"

	"github.com/benaychh/aocutils"
)

func largestArea(data []aocutils.Vertex) (aocutils.Vertex, int) {
	gridMaxWidth, gridMaxHeight := getBounds(data)
	gridMap := [][]*pointClaim{}
	for i := 0; i <= gridMaxWidth; i++ {
		row := []*pointClaim{}
		for j := 0; j <= gridMaxHeight; j++ {
			row = append(row, nil)
		}
		gridMap = append(gridMap, row)
	}

	// Plot the data points
	for _, center := range data {
		gridMap[center.X][center.Y] = &pointClaim{points: []aocutils.Vertex{center}, radius: 0}
	}

	// Loop through the points making bigger and bigger manhattan circles
	keepGoingMap := map[aocutils.Vertex]bool{}
	keepGoing := true
	radius := 1
	loopCount := 0
	for keepGoing {
		keepGoing = false
		for _, center := range data {
			shouldKeepGoing, valueExists := keepGoingMap[center]
			if shouldKeepGoing || !valueExists {
				tempKeepGoing, tempLoopCount := claimManhattanCircle(center, radius, gridMaxWidth, gridMaxHeight, gridMap)
				loopCount += tempLoopCount
				keepGoingMap[center] = tempKeepGoing
				keepGoing = tempKeepGoing || keepGoing
			}
		}
		radius++
	}

	fmt.Printf("radius %d, loopCount %d\n", radius, loopCount)

	// Get any point that goes to an edge, these are disqualified from counting
	edges := getEdges(gridMaxWidth, gridMaxHeight, gridMap)

	// These are the only vertices that can produce an answer
	middles := getMiddles(data, edges)

	areas := getAreas(gridMap, middles)

	return getLargestArea(areas)
}

func getAreas(gridMap [][]*pointClaim, middles map[aocutils.Vertex]bool) (areas map[aocutils.Vertex]int) {
	areas = map[aocutils.Vertex]int{}
	for _, row := range gridMap {
		for _, pointClaim := range row {
			if len(pointClaim.points) == 1 {
				vertex := pointClaim.points[0]
				_, isMiddle := middles[pointClaim.points[0]]
				if isMiddle {
					count, hasCount := areas[vertex]
					if hasCount {
						areas[vertex] = count + 1
					} else {
						areas[vertex] = 1
					}
				}
			}
		}
	}
	return
}

func getMiddles(data []aocutils.Vertex, edges map[aocutils.Vertex]bool) (middles map[aocutils.Vertex]bool) {
	middles = map[aocutils.Vertex]bool{}
	for _, center := range data {
		_, ok := edges[center]
		if !ok {
			middles[center] = true
		}
	}
	return
}

func getLargestArea(areas map[aocutils.Vertex]int) (largestAreaVertex aocutils.Vertex, largestArea int) {
	for vertex, area := range areas {
		if area > largestArea {
			largestAreaVertex = vertex
			largestArea = area
		}
	}
	return
}

func getEdges(gridMaxWidth, gridMaxHeight int, gridMap [][]*pointClaim) map[aocutils.Vertex]bool {
	edges := map[aocutils.Vertex]bool{}
	// Top and bottom of the grid
	for i := 0; i <= gridMaxWidth; i++ {
		topSpace := gridMap[i][0]
		// Only count as an edge if only one center owns this space
		if len(topSpace.points) == 1 {
			edges[topSpace.points[0]] = true
		}
		bottomSpace := gridMap[i][gridMaxHeight]
		// Only count as an edge if only one center owns this space
		if len(bottomSpace.points) == 1 {
			edges[bottomSpace.points[0]] = true
		}
	}

	// Left and right of the grid
	for i := 1; i < gridMaxHeight; i++ {
		leftSpace := gridMap[0][i]
		// Only count as an edge if only one center owns this space
		if len(leftSpace.points) == 1 {
			edges[leftSpace.points[0]] = true
		}
		rightSpace := gridMap[gridMaxWidth][i]
		// Only count as an edge if only one center owns this space
		if len(rightSpace.points) == 1 {
			edges[rightSpace.points[0]] = true
		}
	}
	return edges
}

func claimManhattanCircle(center aocutils.Vertex, radius, gridMaxWidth, gridMaxHeight int, grid [][]*pointClaim) (areAnySpacesClaimable bool, loopCount int) {
	// Start directly above vertext
	x := center.X
	y := center.Y - radius
	areAnySpacesClaimable = false
	// Start on the graph
	if y < 0 {
		x = x + (0 - y)
		y = 0
	}
	// Center Top to Right Center ↘
	for x < center.X+radius {
		loopCount++
		space := aocutils.Vertex{X: x, Y: y}
		// Set to true if we can claim at least one space - don't ever reset to false
		areAnySpacesClaimable = tryToClaimSpace(space, center, radius, gridMaxWidth, gridMaxHeight, grid) || areAnySpacesClaimable
		x = x + 1
		y = y + 1
		// Outside of the graph, skip to the end
		if x > gridMaxWidth {
			break
		}
	}
	// Start on the graph
	if x > gridMaxWidth {
		y = y + 2*(center.Y-y)
	}
	// Right Center to Center Bottom ↙
	for x > center.X {
		loopCount++
		space := aocutils.Vertex{X: x, Y: y}
		// Set to true if we can claim at least one space - don't ever reset to false
		areAnySpacesClaimable = tryToClaimSpace(space, center, radius, gridMaxWidth, gridMaxHeight, grid) || areAnySpacesClaimable
		x = x - 1
		y = y + 1
		// Outside of the graph, skip to the end
		if y > gridMaxHeight {
			break
		}
	}
	// Start on the graph
	if y > gridMaxHeight {
		x = x - 2*(x-center.X)
	}
	// Center Bottom to Left Center ↖
	for x > center.X-radius {
		loopCount++
		space := aocutils.Vertex{X: x, Y: y}
		// Set to true if we can claim at least one space - don't ever reset to false
		areAnySpacesClaimable = tryToClaimSpace(space, center, radius, gridMaxWidth, gridMaxHeight, grid) || areAnySpacesClaimable
		x = x - 1
		y = y - 1
		if x < 0 {
			break
		}
	}
	// Start on the graph
	if x < 0 {
		y = y - 2*(y-center.Y)
	}
	// Left Center to Center Top ↗
	for x < center.X {
		loopCount++
		space := aocutils.Vertex{X: x, Y: y}
		// Set to true if we can claim at least one space - don't ever reset to false
		areAnySpacesClaimable = tryToClaimSpace(space, center, radius, gridMaxWidth, gridMaxHeight, grid) || areAnySpacesClaimable
		x = x + 1
		y = y - 1
		if y < 0 {
			break
		}
	}
	return
}

func tryToClaimSpace(space aocutils.Vertex, center aocutils.Vertex, radius, gridMaxWidth, gridMaxHeight int, grid [][]*pointClaim) bool {
	// Check to make sure we are inside the grid
	if space.X <= gridMaxWidth && space.Y <= gridMaxHeight && space.X >= 0 && space.Y >= 0 {
		claim := grid[space.X][space.Y]
		// This space has not been claimed before
		if claim == nil {
			newClaim := &pointClaim{points: []aocutils.Vertex{center}, radius: radius}
			grid[space.X][space.Y] = newClaim
			return true
		} else if radius == claim.radius {
			// This is a shared space
			claim.points = append(claim.points, center)
			return false
		}
	}
	return false
}
