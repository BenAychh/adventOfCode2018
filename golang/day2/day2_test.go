package main

import (
	"testing"
)

func TestDay2Part1(t *testing.T) {
	data := []string{"abcdef", "bababc", "abbcde", "abcccd", "aabcdd", "abcdee", "ababab"}
	expected := 12
	actual := checksum(data)

	if actual != expected {
		t.Errorf("Expected %v to be %d but instead got %d", data, expected, actual)
	}
}

func TestDay2Part2(t *testing.T) {
	data := []string{"abcde", "fghij", "klmno", "pqrst", "fguij", "axcye", "wvxyz"}
	expected := "fgij"
	actual := commonLetters(data)

	if actual != expected {
		t.Errorf("Expected %v to be %s but instead got %s", data, expected, actual)
	}
}
