package main

import (
	"testing"
)

func TestMessage(t *testing.T) {
	initialState, notes := getPuzzleInput("day12.test.data")
	expected := 325
	actual := getGeneration(initialState, notes, 20)
	if actual != expected {
		t.Errorf("Expected test data to be %d but instead got %d", expected, actual)
	}
}
