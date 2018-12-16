package main

import (
	"aocutils"
	"testing"
)

func TestExamplePart1_1(t *testing.T) {
	fileName := "test.data.1"
	data := aocutils.LoadArrayOfStringsFromTextFile(fileName)
	grid, entities := processMap(data)
	expectedRounds, expectedHitpoints := 47, 590
	actualRounds, actualHitpoints := getRoundsAndRemainingHitpoints(grid, entities, false)

	if actualRounds != expectedRounds || actualHitpoints != expectedHitpoints {
		t.Errorf("Expected %v to be %d,%d but instead got %d,%d", fileName, expectedRounds, expectedHitpoints, actualRounds, actualHitpoints)
	}
}

func TestExamplePart1_2(t *testing.T) {
	fileName := "test.data.2"
	data := aocutils.LoadArrayOfStringsFromTextFile(fileName)
	grid, entities := processMap(data)
	expectedRounds, expectedHitpoints := 37, 982
	actualRounds, actualHitpoints := getRoundsAndRemainingHitpoints(grid, entities, false)

	if actualRounds != expectedRounds || actualHitpoints != expectedHitpoints {
		t.Errorf("Expected %v to be %d,%d but instead got %d,%d", fileName, expectedRounds, expectedHitpoints, actualRounds, actualHitpoints)
	}
}

func TestExamplePart1_3(t *testing.T) {
	fileName := "test.data.3"
	data := aocutils.LoadArrayOfStringsFromTextFile(fileName)
	grid, entities := processMap(data)
	expectedRounds, expectedHitpoints := 46, 859
	actualRounds, actualHitpoints := getRoundsAndRemainingHitpoints(grid, entities, false)

	if actualRounds != expectedRounds || actualHitpoints != expectedHitpoints {
		t.Errorf("Expected %v to be %d,%d but instead got %d,%d", fileName, expectedRounds, expectedHitpoints, actualRounds, actualHitpoints)
	}
}

func TestExamplePart1_4(t *testing.T) {
	fileName := "test.data.4"
	data := aocutils.LoadArrayOfStringsFromTextFile(fileName)
	grid, entities := processMap(data)
	expectedRounds, expectedHitpoints := 35, 793
	actualRounds, actualHitpoints := getRoundsAndRemainingHitpoints(grid, entities, false)

	if actualRounds != expectedRounds || actualHitpoints != expectedHitpoints {
		t.Errorf("Expected %v to be %d,%d but instead got %d,%d", fileName, expectedRounds, expectedHitpoints, actualRounds, actualHitpoints)
	}
}

func TestExamplePart1_5(t *testing.T) {
	fileName := "test.data.5"
	data := aocutils.LoadArrayOfStringsFromTextFile(fileName)
	grid, entities := processMap(data)
	expectedRounds, expectedHitpoints := 54, 536
	actualRounds, actualHitpoints := getRoundsAndRemainingHitpoints(grid, entities, false)

	if actualRounds != expectedRounds || actualHitpoints != expectedHitpoints {
		t.Errorf("Expected %v to be %d,%d but instead got %d,%d", fileName, expectedRounds, expectedHitpoints, actualRounds, actualHitpoints)
	}
}

func TestExamplePart1_6(t *testing.T) {
	fileName := "test.data.6"
	data := aocutils.LoadArrayOfStringsFromTextFile(fileName)
	grid, entities := processMap(data)
	expectedRounds, expectedHitpoints := 20, 937
	actualRounds, actualHitpoints := getRoundsAndRemainingHitpoints(grid, entities, false)

	if actualRounds != expectedRounds || actualHitpoints != expectedHitpoints {
		t.Errorf("Expected %v to be %d,%d but instead got %d,%d", fileName, expectedRounds, expectedHitpoints, actualRounds, actualHitpoints)
	}
}

func TestExamplePart2_1(t *testing.T) {
	fileName := "test.data.1"
	data := aocutils.LoadArrayOfStringsFromTextFile(fileName)
	grid, entities := processMap(data)
	expectedRounds, expectedHitpoints := 29, 172
	actualRounds, actualHitpoints := findLowestPowerNeededToKeepElvesAlive(grid, entities, false)

	if actualRounds != expectedRounds || actualHitpoints != expectedHitpoints {
		t.Errorf("Expected %v to be %d,%d but instead got %d,%d", fileName, expectedRounds, expectedHitpoints, actualRounds, actualHitpoints)
	}
}

func TestExamplePart2_2(t *testing.T) {
	fileName := "test.data.3"
	data := aocutils.LoadArrayOfStringsFromTextFile(fileName)
	grid, entities := processMap(data)
	expectedRounds, expectedHitpoints := 33, 948
	actualRounds, actualHitpoints := findLowestPowerNeededToKeepElvesAlive(grid, entities, false)

	if actualRounds != expectedRounds || actualHitpoints != expectedHitpoints {
		t.Errorf("Expected %v to be %d,%d but instead got %d,%d", fileName, expectedRounds, expectedHitpoints, actualRounds, actualHitpoints)
	}
}

func TestExamplePart2_3(t *testing.T) {
	fileName := "test.data.4"
	data := aocutils.LoadArrayOfStringsFromTextFile(fileName)
	grid, entities := processMap(data)
	expectedRounds, expectedHitpoints := 37, 94
	actualRounds, actualHitpoints := findLowestPowerNeededToKeepElvesAlive(grid, entities, false)

	if actualRounds != expectedRounds || actualHitpoints != expectedHitpoints {
		t.Errorf("Expected %v to be %d,%d but instead got %d,%d", fileName, expectedRounds, expectedHitpoints, actualRounds, actualHitpoints)
	}
}

func TestExamplePart2_4(t *testing.T) {
	fileName := "test.data.5"
	data := aocutils.LoadArrayOfStringsFromTextFile(fileName)
	grid, entities := processMap(data)
	expectedRounds, expectedHitpoints := 39, 166
	actualRounds, actualHitpoints := findLowestPowerNeededToKeepElvesAlive(grid, entities, false)

	if actualRounds != expectedRounds || actualHitpoints != expectedHitpoints {
		t.Errorf("Expected %v to be %d,%d but instead got %d,%d", fileName, expectedRounds, expectedHitpoints, actualRounds, actualHitpoints)
	}
}

func TestExamplePart2_5(t *testing.T) {
	fileName := "test.data.6"
	data := aocutils.LoadArrayOfStringsFromTextFile(fileName)
	grid, entities := processMap(data)
	expectedRounds, expectedHitpoints := 30, 38
	actualRounds, actualHitpoints := findLowestPowerNeededToKeepElvesAlive(grid, entities, false)

	if actualRounds != expectedRounds || actualHitpoints != expectedHitpoints {
		t.Errorf("Expected %v to be %d,%d but instead got %d,%d", fileName, expectedRounds, expectedHitpoints, actualRounds, actualHitpoints)
	}
}
