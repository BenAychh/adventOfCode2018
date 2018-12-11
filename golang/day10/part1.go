package main

import (
	"fmt"
	"math"

	"github.com/BenAychh/aocutils"
)

func getMessage(points []movingPoint) (seconds int) {
	lowestLineCountSecond := 0
	randomNonZeroNum := 100
	lastFiveLineCounts := []int{randomNonZeroNum, randomNonZeroNum, randomNonZeroNum, randomNonZeroNum, randomNonZeroNum}
	wasWordFound := false
	for i := 1; !wasWordFound && /* loop safety */ i < 100000; i++ {
		currentVertices := calculatePointsLocation(points, i)
		lineCount := calculateRowCounts(currentVertices)
		lastFiveLineCounts = append(lastFiveLineCounts[1:], lineCount)
		wasWordFound = checkIfValley(lastFiveLineCounts)
		lowestLineCountSecond = i - 2
	}
	printGraph(calculatePointsLocation(points, lowestLineCountSecond))
	return lowestLineCountSecond
}

func checkIfValley(values []int) bool {
	temp := make([]int, len(values))
	copy(temp, values)
	halfwayIndex := int(math.Floor(float64(len(temp)) / 2.0))
	middleValue := temp[halfwayIndex]
	remainder := append(temp[:halfwayIndex], temp[halfwayIndex+1:]...)
	for _, remainder := range remainder {
		if middleValue >= remainder {
			return false
		}
	}
	return true
}

func calculatePointsLocation(bases []movingPoint, seconds int) (vertices []aocutils.Vertex) {
	for _, vertex := range bases {
		vertices = append(vertices, calculatePointLocation(vertex, seconds))
	}
	return
}

func calculatePointLocation(base movingPoint, seconds int) aocutils.Vertex {
	return aocutils.Vertex{
		X: base.x + base.dx*seconds,
		Y: base.y + base.dy*seconds,
	}
}

func calculateRowCounts(vertices []aocutils.Vertex) (rowCount int) {
	min := vertices[0].Y
	max := vertices[0].Y

	for _, vertex := range vertices {
		if vertex.Y > max {
			max = vertex.Y
		} else if vertex.Y < min {
			min = vertex.Y
		}
	}
	return max - min
}

func printGraph(vertices []aocutils.Vertex) {
	minX, maxX, minY, maxY := getBounds(vertices)
	maxWidth := maxX - minX
	maxHeight := maxY - minY
	picture := [][]rune{}
	for i := 0; i <= maxHeight; i++ {
		row := []rune{}
		for j := 0; j <= (maxWidth); j++ {
			row = append(row, '.')
		}
		picture = append(picture, row)
	}

	for _, vertex := range vertices {
		picture[vertex.Y-minY][vertex.X-minX] = '#'
	}

	for _, line := range picture {
		fmt.Println(string(line))
	}
}

func getBounds(data []aocutils.Vertex) (minX, maxX, minY, maxY int) {
	minX = data[0].X
	maxX = data[0].X
	minY = data[0].Y
	maxY = data[0].Y

	for _, vertex := range data {
		if vertex.X < minX {
			minX = vertex.X
		}
		if vertex.X > maxX {
			maxX = vertex.X
		}
		if vertex.Y < minY {
			minY = vertex.Y
		}
		if vertex.Y > maxY {
			maxY = vertex.Y
		}
	}
	return
}
