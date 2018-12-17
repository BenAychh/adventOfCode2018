package main

import (
	"testing"

	"github.com/benaychh/aocutils"
)

func TestCellCalculations1(t *testing.T) {
	expected := 4
	actual := getFuelInCell(3, 5, 8)
	if actual != expected {
		t.Errorf("Expected (3, 5, 8) to be %d but instead got %d", expected, actual)
	}
}

func TestCellCalculations2(t *testing.T) {
	expected := -5
	actual := getFuelInCell(122, 79, 57)
	if actual != expected {
		t.Errorf("Expected (3, 5, 8) to be %d but instead got %d", expected, actual)
	}
}

func TestCellCalculations3(t *testing.T) {
	expected := 0
	actual := getFuelInCell(217, 196, 39)
	if actual != expected {
		t.Errorf("Expected (3, 5, 8) to be %d but instead got %d", expected, actual)
	}
}

func TestCellCalculations4(t *testing.T) {
	expected := 4
	actual := getFuelInCell(101, 153, 71)
	if actual != expected {
		t.Errorf("Expected (3, 5, 8) to be %d but instead got %d", expected, actual)
	}
}

func TestLargest3x3Grid1(t *testing.T) {
	data := 18
	expected := aocutils.Vertex{X: 33, Y: 45}
	actual, power := getHighestPower(data)
	if actual != expected {
		t.Errorf("Expected %d to be %d but instead got %d (with power %d)", data, expected, actual, power)
	}
}

func TestLargest3x3Grid2(t *testing.T) {
	data := 42
	expected := aocutils.Vertex{X: 21, Y: 61}
	actual, power := getHighestPower(data)
	if actual != expected {
		t.Errorf("Expected %d to be %d but instead got %d (with power %d)", data, expected, actual, power)
	}
}

func TestLargestXxXGrid1(t *testing.T) {
	data := 18
	expected := aocutils.Vertex{X: 90, Y: 269}
	actual, power, size := getHighestPowerAnySize(data)
	if actual != expected {
		t.Errorf("Expected %d to be %d but instead got %d (with power %d and size %d)", data, expected, actual, power, size)
	}
}

func TestLargestXxXGrid2(t *testing.T) {
	data := 42
	expected := aocutils.Vertex{X: 232, Y: 251}
	actual, power, size := getHighestPowerAnySize(data)
	if actual != expected {
		t.Errorf("Expected %d to be %d but instead got %d (with power %d and size %d)", data, expected, actual, power, size)
	}
}
