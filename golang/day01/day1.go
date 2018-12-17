package main

import (
	"fmt"
	"time"

	"github.com/benaychh/aocutils"
)

func main() {
	resultAsInts := aocutils.LoadFileAsIntegerArray("day1.data")

	start := time.Now().UnixNano()
	part1 := processor(resultAsInts)
	end := time.Now().UnixNano()
	fmt.Println(part1)
	fmt.Println(float32(end-start) / 1000.0 / 1000.0)

	start = time.Now().UnixNano()
	part2, count := repeatFinder(resultAsInts)
	end = time.Now().UnixNano()
	fmt.Println(part2, "Count: ", count)
	fmt.Println(float32(end-start) / 1000.0 / 1000.0)
}

func processor(data []int) (sum int) {
	sum = 0

	for _, i := range data {
		sum += i
	}
	return
}

func repeatFinder(data []int) (sum int, count int) {
	sum = 0
	count = 0
	hasRepeated := false
	freqMap := make(map[int]bool)
	freqMap[0] = true

	for !hasRepeated {
		count++
		for _, i := range data {
			sum += i
			_, ok := freqMap[sum]
			if ok {
				hasRepeated = true
				return
			} else {
				freqMap[sum] = true
			}
		}
	}
	return
}
