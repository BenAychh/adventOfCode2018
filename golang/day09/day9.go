package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now().UnixNano()
	part1 := getHighestScore(403, 71920)
	end := time.Now().UnixNano()
	fmt.Printf("Part 1: %v, Time: %fms\n", part1, float32(end-start)/1000.0/1000.0)

	start = time.Now().UnixNano()
	part2 := getHighestScore(403, 71920*100)
	end = time.Now().UnixNano()
	fmt.Printf("Part 1: %v, Time: %fms\n", part2, float32(end-start)/1000.0/1000.0)
}
