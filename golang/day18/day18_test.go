package main

import (
	"testing"

	"github.com/benaychh/aocutils"
)

func TestResourceValue(t *testing.T) {
	fileName := "day18.test.data"
	data := aocutils.LoadFileAsStringArray(fileName)
	grid := getGrid(data)
	expectedTrees, expectedLumberyards := 37, 31
	actualTrees, actualLumberyards := getResourceValueAfter(grid, 10)

	if actualTrees != expectedTrees || actualLumberyards != actualLumberyards {
		t.Errorf("Expected %v to be %dx%d (after 10 minutes) but instead got %dx%d", fileName, expectedTrees, expectedLumberyards, actualTrees, actualLumberyards)
		printGrid(grid)
	}
}
