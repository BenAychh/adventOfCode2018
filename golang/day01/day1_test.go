package main

import (
	"testing"
)

func TestDay1Part1(t *testing.T) {
	t.Run("+1, +1, +1 results in  3", day1Part1Func([]int{1, 1, 1}, 3))
	t.Run("+1, +1, -2 results in  0", day1Part1Func([]int{1, 1, -2}, 0))
	t.Run("-1, -2, -3 results in -6", day1Part1Func([]int{-1, -2, -3}, -6))
}

func TestDay1Part2(t *testing.T) {
	t.Run("+1, -1 first reaches 0 twice", day1Part2Func([]int{1, -1}, 0))
	t.Run("+3, +3, +4, -2, -4 first reaches 10 twice", day1Part2Func([]int{3, 3, 4, -2, -4}, 10))
	t.Run("-6, +3, +8, +5, -6 first reaches 5 twice.", day1Part2Func([]int{-6, 3, 8, 5, -6}, 5))
	t.Run("+7, +7, -2, -7, -4 first reaches 14 twice.", day1Part2Func([]int{7, 7, -2, -7, -4}, 14))
}

func day1Part1Func(data []int, expected int) func(*testing.T) {
	return func(t *testing.T) {
		actual := processor(data)

		if actual != expected {
			t.Errorf("Expected %v to be %d but instead got %d", data, expected, actual)
		}
	}
}

func day1Part2Func(data []int, expected int) func(*testing.T) {
	return func(t *testing.T) {
		actual, _ := repeatFinder(data)

		if actual != expected {
			t.Errorf("Expected %v to be %d but instead got %d", data, expected, actual)
		}
	}
}
