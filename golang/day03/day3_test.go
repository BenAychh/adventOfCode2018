package main

import (
	"testing"
)

func TestDay3Part1(t *testing.T) {
	data := []string{"#1 @ 1,3: 4x4", "#2 @ 3,1: 4x4", "#3 @ 5,5: 2x2"}
	expected := 4
	actual, _, _ := overlap(data)

	if actual != expected {
		t.Errorf("Expected %v to be %d but instead got %d", data, expected, actual)
	}
}

func TestDay3Part2(t *testing.T) {
	data := []string{"#1 @ 1,3: 4x4", "#2 @ 3,1: 4x4", "#3 @ 5,5: 2x2"}
	expected := 3
	_, grid, claims := overlap(data)
	actual := noOverLap(grid, claims)

	if actual != expected {
		t.Errorf("Expected %v to be %v but instead got %v", data, expected, actual)
	}
}
