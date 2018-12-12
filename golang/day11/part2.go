package main

import (
	"github.com/BenAychh/aocutils"
)

func getHighestPowerAnySize(serial int) (aocutils.Vertex, int, int) {
	getFuelCell := getFuelCellFactory(serial)

	sizing := sizingInfo{
		previousValues: createPreviousValuesGrid(),
		rows:           createPreviousValuesGrid(),
		cols:           createPreviousValuesGrid(),
	}

	largest := 0
	size := 0
	vertex := aocutils.Vertex{}

	for s := 1; s < 300; s++ {
		for x := 0; x < 300-s; x++ {
			for y := 0; y < 300-s; y++ {
				value := getXxXFuel(x, y, s, getFuelCell, &sizing)
				if value > largest {
					largest = value
					size = s
					vertex.X = x + 1
					vertex.Y = y + 1
				}
			}
		}
	}
	return vertex, largest, size
}

type sizingInfo struct {
	previousValues [][]int
	rows           [][]int
	cols           [][]int
}

func getXxXFuel(x, y, size int, getFuelCell func(int, int) int, sizing *sizingInfo) (sum int) {
	sum = sizing.previousValues[y][x]
	extra := getFuelCell(x+size-1, y+size-1)
	column := sizing.cols[y][x+size-1] + extra
	sizing.cols[y][x+size-1] = column
	row := sizing.rows[y+size-1][x] + extra
	sizing.rows[y+size-1][x] = row
	sum += column + row - extra
	sizing.previousValues[y][x] = sum
	return
}

func createPreviousValuesGrid() (grid [][]int) {
	grid = make([][]int, 300)
	for i := range grid {
		grid[i] = make([]int, 300)
	}
	return
}
