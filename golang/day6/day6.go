package main

import (
	"fmt"
	"time"

	"github.com/BenAychh/aocutils"
)

func main() {
	data := aocutils.LoadFileAsVertices("day6.data")

	start := time.Now().UnixNano()
	part1Vertex, part1Area := largestArea(data)
	end := time.Now().UnixNano()
	fmt.Printf("Part 1: %d (%v), Time: %fms\n", part1Area, part1Vertex, float32(end-start)/1000.0/1000.0)

	start = time.Now().UnixNano()
	part1AltVertex, part1AltArea := largestArea(data)
	end = time.Now().UnixNano()
	fmt.Printf("Part 1 Alt: %d (%v), Time: %fms\n", part1AltArea, part1AltVertex, float32(end-start)/1000.0/1000.0)

	start = time.Now().UnixNano()
	part2 := underDistanceLimitFromEverything(data, 10000)
	end = time.Now().UnixNano()
	fmt.Printf("Part 2: %d, Time: %fms\n", part2, float32(end-start)/1000.0/1000.0)
}

type pointClaim struct {
	points []aocutils.Vertex
	radius int
}

func underDistanceLimitFromEverything(data []aocutils.Vertex, distanceLimit int) int {
	gridMaxWidth, gridMaxHeight := getBounds(data)
	underDistanceLimitCount := 0
	for x := 0; x <= gridMaxWidth; x++ {
		for y := 0; y <= gridMaxHeight; y++ {
			sum := 0
			for _, center := range data {
				sum += calcManhattanDistance(x, y, center)
				if sum > distanceLimit {
					break
				}
			}
			if sum < distanceLimit {
				underDistanceLimitCount++
			}
		}
	}
	return underDistanceLimitCount
}

func calcManhattanDistance(x, y int, center aocutils.Vertex) int {
	return abs(x-center.X) + abs(y-center.Y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func largestArea(data []aocutils.Vertex) (aocutils.Vertex, int) {
	gridMaxWidth, gridMaxHeight := getBounds(data)
	gridMap := map[string]pointClaim{}

	// Plot the data points
	for _, center := range data {
		gridMap[getVertexKey(center)] = pointClaim{points: []aocutils.Vertex{center}, radius: 0}
	}

	// Loop through the points making bigger and bigger manhattan circles
	keepGoing := true
	radius := 1
	for keepGoing {
		keepGoing = false
		for _, center := range data {
			// Only keep going if claimManhattanCircle was able to claim at least one space previously - don't even return to false
			keepGoing = claimManhattanCircle(center, radius, gridMaxWidth, gridMaxHeight, gridMap) || keepGoing
		}
		radius++
	}

	// Get any point that goes to an edge, these are disqualified from counting
	edges := getEdges(gridMaxWidth, gridMaxHeight, gridMap)

	// These are the only vertices that can produce an answer
	middles := map[aocutils.Vertex]bool{}
	for _, center := range data {
		_, ok := edges[center]
		if !ok {
			middles[center] = true
		}
	}

	areas := map[aocutils.Vertex]int{}

	for _, pointClaim := range gridMap {
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

	largestArea := 0
	var largestAreaVertex aocutils.Vertex

	for vertex, area := range areas {
		if area > largestArea {
			largestAreaVertex = vertex
			largestArea = area
		}
	}

	return largestAreaVertex, largestArea
}

func getBounds(data []aocutils.Vertex) (maxX, maxY int) {
	maxX = 0
	maxY = 0
	for _, vertex := range data {
		if vertex.X > maxX {
			maxX = vertex.X
		}
		if vertex.Y > maxY {
			maxY = vertex.Y
		}
	}
	return
}

func getEdges(gridMaxWidth, gridMaxHeight int, gridMap map[string]pointClaim) map[aocutils.Vertex]bool {
	edges := map[aocutils.Vertex]bool{}
	// Top and bottom of the grid
	for i := 0; i <= gridMaxWidth; i++ {
		topKey := fmt.Sprintf("%d,%d", i, 0)
		bottomKey := fmt.Sprintf("%d,%d", i, gridMaxHeight)
		topSpace := gridMap[topKey]
		// Only count as an edge if only one center owns this space
		if len(topSpace.points) == 1 {
			edges[topSpace.points[0]] = true
		}
		bottomSpace := gridMap[bottomKey]
		// Only count as an edge if only one center owns this space
		if len(bottomSpace.points) == 1 {
			edges[bottomSpace.points[0]] = true
		}
	}

	// Left and right of the grid
	for i := 1; i < gridMaxHeight; i++ {
		leftKey := fmt.Sprintf("%d,%d", 0, i)
		rightKey := fmt.Sprintf("%d,%d", gridMaxWidth, i)
		leftSpace := gridMap[leftKey]
		// Only count as an edge if only one center owns this space
		if len(leftSpace.points) == 1 {
			edges[leftSpace.points[0]] = true
		}
		rightSpace := gridMap[rightKey]
		// Only count as an edge if only one center owns this space
		if len(rightSpace.points) == 1 {
			edges[rightSpace.points[0]] = true
		}
	}
	return edges
}

func claimManhattanCircle(center aocutils.Vertex, radius, gridMaxWidth, gridMaxHeight int, grid map[string]pointClaim) (areAnySpacesClaimable bool) {
	// Start directly above vertext
	x := center.X
	y := center.Y - radius
	areAnySpacesClaimable = false
	// Center Top to Right Center ↘
	for x < center.X+radius {
		space := aocutils.Vertex{X: x, Y: y}
		// Set to true if we can claim at least one space - don't ever reset to false
		areAnySpacesClaimable = tryToClaimSpace(space, center, radius, gridMaxWidth, gridMaxHeight, grid) || areAnySpacesClaimable
		x = x + 1
		y = y + 1
		// Outside of the graph, skip to the end
	}
	// Right Center to Center Bottom ↙
	for x > center.X {
		space := aocutils.Vertex{X: x, Y: y}
		// Set to true if we can claim at least one space - don't ever reset to false
		areAnySpacesClaimable = tryToClaimSpace(space, center, radius, gridMaxWidth, gridMaxHeight, grid) || areAnySpacesClaimable
		x = x - 1
		y = y + 1
		// Outside of the graph, skip to the end
	}
	// Center Bottom to Left Center ↖
	for x > center.X-radius {
		space := aocutils.Vertex{X: x, Y: y}
		// Set to true if we can claim at least one space - don't ever reset to false
		areAnySpacesClaimable = tryToClaimSpace(space, center, radius, gridMaxWidth, gridMaxHeight, grid) || areAnySpacesClaimable
		x = x - 1
		y = y - 1
	}
	// Left Center to Center Top ↗
	for x < center.X {
		space := aocutils.Vertex{X: x, Y: y}
		// Set to true if we can claim at least one space - don't ever reset to false
		areAnySpacesClaimable = tryToClaimSpace(space, center, radius, gridMaxWidth, gridMaxHeight, grid) || areAnySpacesClaimable
		x = x + 1
		y = y - 1
	}
	return
}

func tryToClaimSpace(space aocutils.Vertex, center aocutils.Vertex, radius, gridMaxWidth, gridMaxHeight int, grid map[string]pointClaim) bool {
	// Check to make sure we are inside the grid
	if space.X <= gridMaxWidth && space.Y <= gridMaxHeight && space.X >= 0 && space.Y >= 0 {
		spaceKey := getVertexKey(space)
		claim, ok := grid[spaceKey]
		// This space has not been claimed before
		if !ok {
			newClaim := pointClaim{points: []aocutils.Vertex{center}, radius: radius}
			grid[spaceKey] = newClaim
			return true
		} else if radius == claim.radius {
			// This is a shared space
			claim.points = append(claim.points, center)
			grid[spaceKey] = claim
			return false
		}
	}
	return false
}

func getVertexKey(vertex aocutils.Vertex) string {
	return getKey(vertex.X, vertex.Y)
}

func getKey(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}
