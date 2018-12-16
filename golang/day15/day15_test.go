package main

import (
	"aocutils"
	"testing"
)

func TestExample1(t *testing.T) {
	fileName := "test.data.1"
	data := aocutils.LoadArrayOfStringsFromTextFile(fileName)
	grid, entities := processMap(data)
	expectedRounds, expectedHitpoints := 47, 590
	actualRounds, actualHitpoints := getRoundsAndRemainingHitpoints(grid, entities, false)

	if actualRounds != expectedRounds || actualHitpoints != expectedHitpoints {
		t.Errorf("Expected %v to be %d,%d but instead got %d,%d", fileName, expectedRounds, expectedHitpoints, actualRounds, actualHitpoints)
	}
}

func TestExample2(t *testing.T) {
	fileName := "test.data.2"
	data := aocutils.LoadArrayOfStringsFromTextFile(fileName)
	grid, entities := processMap(data)
	expectedRounds, expectedHitpoints := 37, 982
	actualRounds, actualHitpoints := getRoundsAndRemainingHitpoints(grid, entities, false)

	if actualRounds != expectedRounds || actualHitpoints != expectedHitpoints {
		t.Errorf("Expected %v to be %d,%d but instead got %d,%d", fileName, expectedRounds, expectedHitpoints, actualRounds, actualHitpoints)
	}
}

func TestExample3(t *testing.T) {
	fileName := "test.data.3"
	data := aocutils.LoadArrayOfStringsFromTextFile(fileName)
	grid, entities := processMap(data)
	expectedRounds, expectedHitpoints := 46, 859
	actualRounds, actualHitpoints := getRoundsAndRemainingHitpoints(grid, entities, false)

	if actualRounds != expectedRounds || actualHitpoints != expectedHitpoints {
		t.Errorf("Expected %v to be %d,%d but instead got %d,%d", fileName, expectedRounds, expectedHitpoints, actualRounds, actualHitpoints)
	}
}

func TestExample4(t *testing.T) {
	fileName := "test.data.4"
	data := aocutils.LoadArrayOfStringsFromTextFile(fileName)
	grid, entities := processMap(data)
	expectedRounds, expectedHitpoints := 35, 793
	actualRounds, actualHitpoints := getRoundsAndRemainingHitpoints(grid, entities, false)

	if actualRounds != expectedRounds || actualHitpoints != expectedHitpoints {
		t.Errorf("Expected %v to be %d,%d but instead got %d,%d", fileName, expectedRounds, expectedHitpoints, actualRounds, actualHitpoints)
	}
}

func TestExample5(t *testing.T) {
	fileName := "test.data.5"
	data := aocutils.LoadArrayOfStringsFromTextFile(fileName)
	grid, entities := processMap(data)
	expectedRounds, expectedHitpoints := 54, 536
	actualRounds, actualHitpoints := getRoundsAndRemainingHitpoints(grid, entities, false)

	if actualRounds != expectedRounds || actualHitpoints != expectedHitpoints {
		t.Errorf("Expected %v to be %d,%d but instead got %d,%d", fileName, expectedRounds, expectedHitpoints, actualRounds, actualHitpoints)
	}
}

func TestExample6(t *testing.T) {
	fileName := "test.data.6"
	data := aocutils.LoadArrayOfStringsFromTextFile(fileName)
	grid, entities := processMap(data)
	expectedRounds, expectedHitpoints := 20, 937
	actualRounds, actualHitpoints := getRoundsAndRemainingHitpoints(grid, entities, false)

	if actualRounds != expectedRounds || actualHitpoints != expectedHitpoints {
		t.Errorf("Expected %v to be %d,%d but instead got %d,%d", fileName, expectedRounds, expectedHitpoints, actualRounds, actualHitpoints)
	}
}
