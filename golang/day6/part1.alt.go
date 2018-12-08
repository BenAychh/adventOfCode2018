package main

import (
	"fmt"

	"github.com/BenAychh/aocutils"
)

func largestAreaAlt(data []aocutils.Vertex) (aocutils.Vertex, int) {
	gridMaxWidth, gridMaxHeight := getBounds(data)

	gridMap := getClosestDistances(data, gridMaxWidth, gridMaxHeight)

	getEdges(gridMaxWidth, gridMaxHeight, gridMap)

	// Get any point that goes to an edge, these are disqualified from counting
	edges := getEdges(gridMaxWidth, gridMaxHeight, gridMap)

	// These are the only vertices that can produce an answer
	middles := getMiddles(data, edges)

	areas := getAreas(gridMap, middles)

	return getLargestArea(areas)
}

func getClosestDistances(data []aocutils.Vertex, gridMaxWidth, gridMaxHeight int) (everyPoint [][]*pointClaim) {
	everyPoint = [][]*pointClaim{}
	loopCount := 0
	for x := 0; x <= gridMaxWidth; x++ {
		row := []*pointClaim{}
		for y := 0; y <= gridMaxHeight; y++ {
			claim := pointClaim{}
			distance := gridMaxHeight + gridMaxWidth
			for _, center := range data {
				loopCount++
				thisDistance := calcManhattanDistance(x, y, center)
				if thisDistance < distance {
					distance = thisDistance
					claim.points = []aocutils.Vertex{center}
					claim.radius = distance
				} else if thisDistance == distance {
					claim.points = append(claim.points, center)
				}
			}
			row = append(row, &claim)
		}
		everyPoint = append(everyPoint, row)
	}
	fmt.Printf("loopCount %d\n", loopCount)
	return
}
