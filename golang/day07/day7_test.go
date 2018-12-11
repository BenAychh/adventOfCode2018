package main

import (
	"testing"
)

var data = []string{
	"Step C must be finished before step A can begin.",
	"Step C must be finished before step F can begin.",
	"Step A must be finished before step B can begin.",
	"Step A must be finished before step D can begin.",
	"Step B must be finished before step E can begin.",
	"Step D must be finished before step E can begin.",
	"Step F must be finished before step E can begin.",
}

func TestDay7Part1(t *testing.T) {
	expected := "CABDFE"
	actual := getOrder(data)

	if actual != expected {
		t.Errorf("Expected %v to be %v but instead got %v", data, expected, actual)
	}
}

func TestDay7Part2(t *testing.T) {
	expected := 15
	actual := getTotalTime(data, 2, 0)

	if actual != expected {
		t.Errorf("Expected %v to be %v but instead got %v", data, expected, actual)
	}
}
