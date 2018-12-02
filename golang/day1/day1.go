package main

import (
	"aocutils"
	"fmt"
	"time"
)

func main() {
	resultAsInts := aocutils.LoadArrayOfIntsFromTextFile("day1.data")

	part1 := processor(resultAsInts)
	start := time.Now().UnixNano()
	part2 := repeatFinder(resultAsInts)
	end := time.Now().UnixNano()
	fmt.Println(part1)
	fmt.Println(part2)
	fmt.Println(float32(end-start) / 1000.0 / 1000.0)
}

func processor(data []int) (sum int) {
	sum = 0

	for _, i := range data {
		sum += i
	}
	return
}

func repeatFinder(data []int) int {
	sum := 0
	hasRepeated := false
	freqMap := make(map[int]bool)
	freqMap[0] = true

	for !hasRepeated {
		for _, i := range data {
			sum += i
			_, ok := freqMap[sum]
			if ok {
				hasRepeated = true
				return sum
			} else {
				freqMap[sum] = true
			}
		}
	}
	return 0
}
