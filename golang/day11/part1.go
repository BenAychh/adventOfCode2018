package main

import (
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/BenAychh/aocutils"
)

func getHighestPower(serial int) (aocutils.Vertex, int) {
	getFuelCell := getFuelCellFactory(serial)
	largest := 0
	vertex := aocutils.Vertex{}
	for x := 0; x < 298; x++ {
		for y := 0; y < 298; y++ {
			value := get3x3Fuel(x, y, getFuelCell)
			if value > largest {
				largest = value
				vertex.X = x + 1
				vertex.Y = y + 1
			}
		}
	}
	return vertex, largest
}

func get3x3Fuel(x, y int, getFuelCell func(int, int) int) (sum int) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			sum += getFuelCell(x+i, y+j)
		}
	}
	return sum
}

func getFuelCellFactory(serial int) func(x, y int) int {
	memoGrid := createMemoGrid()
	return func(x, y int) int {
		if y < len(memoGrid) && x < len(memoGrid[y]) {
			if memoGrid[y][x] == math.MaxInt64 {
				memoGrid[y][x] = getFuelInCell(x+1, y+1, serial)
			}
			return memoGrid[y][x]
		}
		fmt.Println(x, y)
		fmt.Println(memoGrid[y][x])
		return 0
	}
}

func createMemoGrid() (grid [][]int) {
	grid = [][]int{}
	for y := 0; y <= 300; y++ {
		row := []int{}
		for x := 0; x <= 300; x++ {
			row = append(row, math.MaxInt64)
		}
		grid = append(grid, row)
	}
	return grid
}

func getFuelInCell(x, y, serial int) int {
	rackID := x + 10
	start := rackID * y
	increased := start + serial
	power := increased * rackID
	hundreds := getHundredsDigit(power)
	return hundreds - 5
}

func getHundredsDigit(num int) int {
	if num < 100 {
		return 0
	}
	numberAsString := strconv.Itoa(num)
	digit := numberAsString[len(numberAsString)-3]
	asInt, err := strconv.Atoi(string(digit))
	if err != nil {
		log.Fatal(err)
	}
	return asInt
}
