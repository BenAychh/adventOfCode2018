package main

import (
	"testing"
)

func TestDay5Part1(t *testing.T) {
	data := "dabAcCaCBAcCcaDA"
	expected := "dabCBAcaDA"
	actual := collapsePolymer(data)

	if expected != actual {
		t.Errorf("Expected %v to be %v but instead got %v", string(data), expected, actual)
	}
}

func TestDay5Part1_extra(t *testing.T) {
	data := "JjDdoODgGdNnIiTtJjFfnsSN"
	expected := ""
	actual := collapsePolymer(data)

	if expected != actual {
		t.Errorf("Expected %v to be %v but instead got %v", string(data), expected, actual)
	}
}

func TestDay5Part2(t *testing.T) {
	data := "dabAcCaCBAcCcaDA"
	expected := 4
	letter, actual := getShortestPolymerAfterCleaning(data)

	if expected != actual {
		t.Errorf("Expected %v to be %d (%v) but instead got %d", string(data), expected, string(letter), actual)
	}
}

func TestDay5Part2_extra(t *testing.T) {
	data := "JjDdoODgGdNnIiTtJjFfnsSNsVUaAugGpPLAalvVHRrhvhoOZzQqHLhjJSsOoJjSaAsrqkBbKQRtTHlSqQeEioOICcUuHmMhPVvJjpJjDdhHUuSL"
	expected := 1
	letter, actual := getShortestPolymerAfterCleaning(data)

	if expected != actual {
		t.Errorf("Expected %v to be %d (%v) but instead got %d", string(data), expected, string(letter), actual)
	}
}
