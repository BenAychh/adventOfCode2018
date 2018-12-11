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
	part1AltVertex, part1AltArea := largestAreaAlt(data)
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
