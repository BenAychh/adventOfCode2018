package main

import (
	"testing"
)

var data = []int{2, 3, 0, 3, 10, 11, 12, 1, 1, 0, 1, 99, 2, 1, 1, 2}

func TestPart1(t *testing.T) {
	expected := 138
	_, actual := metadataSum(data)

	if actual != expected {
		t.Errorf("Expected %v to be %d but instead got %d", data, expected, actual)
	}
}

func TestPart2(t *testing.T) {
	rootNode, _ := metadataSum(data)
	expected := 66
	actual := complicatedSum(rootNode)

	if actual != expected {
		t.Errorf("Expected %v to be %d but instead got %d", data, expected, actual)
	}
}
