package main

import (
	"testing"

	"github.com/benaychh/aocutils"
)

var data = []aocutils.Vertex{
	aocutils.Vertex{X: 1, Y: 1},
	aocutils.Vertex{X: 1, Y: 6},
	aocutils.Vertex{X: 8, Y: 3},
	aocutils.Vertex{X: 3, Y: 4},
	aocutils.Vertex{X: 5, Y: 5},
	aocutils.Vertex{X: 8, Y: 9},
}

func TestDay6Part1(t *testing.T) {
	expected := 17
	vertex, actual := largestArea(data)

	if expected != actual {
		t.Errorf("Expected %v to be %d (%+v) but instead got %d", data, expected, vertex, actual)
	}
}

func TestDay6Part1Alt(t *testing.T) {
	expected := 17
	vertex, actual := largestAreaAlt(data)

	if expected != actual {
		t.Errorf("Expected %v to be %d (%+v) but instead got %d", data, expected, vertex, actual)
	}
}

func TestDay6Part2(t *testing.T) {
	expected := 16
	actual := underDistanceLimitFromEverything(data, 32)

	if expected != actual {
		t.Errorf("Expected %v to be %d but instead got %d", data, expected, actual)
	}
}
