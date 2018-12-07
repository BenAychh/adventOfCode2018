package main

import (
	"github.com/BenAychh/aocutils"
)

func largestAreaAlt(data []aocutils.Vertex) (aocutils.Vertex, int) {
	gridMaxWidth, gridMaxHeight := getBounds(data)

	gridMap := getClosestDistances(data, gridMaxWidth, gridMaxHeight)

	getEdges(gridMaxWidth, gridMaxHeight, gridMap)

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

func getClosestDistances(data []aocutils.Vertex, gridMaxWidth, gridMaxHeight int) (everyPoint map[string]pointClaim) {
	everyPoint = map[string]pointClaim{}
	for x := 0; x <= gridMaxWidth; x++ {
		for y := 0; y <= gridMaxHeight; y++ {
			claim := pointClaim{}
			distance := gridMaxHeight + gridMaxWidth
			for _, center := range data {
				thisDistance := calcManhattanDistance(x, y, center)
				if thisDistance < distance {
					distance = thisDistance
					claim.points = []aocutils.Vertex{center}
					claim.radius = distance
				} else if thisDistance == distance {
					claim.points = append(claim.points, center)
				}
			}
			everyPoint[getKey(x, y)] = claim
		}
	}
	return
}
