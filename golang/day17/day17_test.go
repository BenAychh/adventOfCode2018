package main

import (
	"testing"

	"github.com/benaychh/aocutils"
)

func TestWaterCount(t *testing.T) {
	fileName := "day17.test.data"
	data := aocutils.LoadFileAsStringArray(fileName)
	grid := getGrid(data)
	expected := 57
	actual := countWater(grid, 500, 0)

	if actual != expected {
		t.Errorf("Expected %v to be %d but instead got %d", fileName, expected, actual)
		printGrid(grid)
	}
}

func TestSettledWaterCount(t *testing.T) {
	fileName := "day17.test.data"
	data := aocutils.LoadFileAsStringArray(fileName)
	grid := getGrid(data)
	countWater(grid, 500, 0)
	expected := 29
	actual := countSettledWater(grid)

	if actual != expected {
		t.Errorf("Expected %v to be %d but instead got %d", fileName, expected, actual)
		printGrid(grid)
	}
}
