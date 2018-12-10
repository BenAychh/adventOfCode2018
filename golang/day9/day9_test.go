package main

import (
	"testing"
)

func TestPart1Example1(t *testing.T) {
	players := 9
	lastMarble := 25
	expected := 32
	actual := getHighestScore(players, lastMarble)

	if actual != expected {
		t.Errorf("Expected %d players and %d marbles to be %d but instead got %d", players, lastMarble, expected, actual)
	}
}

func TestPart1Example2(t *testing.T) {
	players := 10
	lastMarble := 1618
	expected := 8317
	actual := getHighestScore(players, lastMarble)

	if actual != expected {
		t.Errorf("Expected %d players and %d marbles to be %d but instead got %d", players, lastMarble, expected, actual)
	}
}

func TestPart1Example3(t *testing.T) {
	players := 13
	lastMarble := 7999
	expected := 146373
	actual := getHighestScore(players, lastMarble)

	if actual != expected {
		t.Errorf("Expected %d players and %d marbles to be %d but instead got %d", players, lastMarble, expected, actual)
	}
}

func TestPart1Example4(t *testing.T) {
	players := 17
	lastMarble := 1104
	expected := 2764
	actual := getHighestScore(players, lastMarble)

	if actual != expected {
		t.Errorf("Expected %d players and %d marbles to be %d but instead got %d", players, lastMarble, expected, actual)
	}
}

func TestPart1Example5(t *testing.T) {
	players := 21
	lastMarble := 6111
	expected := 54718
	actual := getHighestScore(players, lastMarble)

	if actual != expected {
		t.Errorf("Expected %d players and %d marbles to be %d but instead got %d", players, lastMarble, expected, actual)
	}
}

func TestPart1Example6(t *testing.T) {
	players := 30
	lastMarble := 5807
	expected := 37305
	actual := getHighestScore(players, lastMarble)

	if actual != expected {
		t.Errorf("Expected %d players and %d marbles to be %d but instead got %d", players, lastMarble, expected, actual)
	}
}
