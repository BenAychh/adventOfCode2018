package main

import "github.com/BenAychh/aocutils"

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
