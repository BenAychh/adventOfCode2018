package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now().UnixNano()
	part1 := getNext10(170641, false)
	end := time.Now().UnixNano()
	fmt.Printf("Part 1: %v, Time: %fms\n", part1, float32(end-start)/1000.0/1000.0)

	start = time.Now().UnixNano()
	part2 := getPrevious5([]int{1, 7, 0, 6, 4, 1}, false)
	end = time.Now().UnixNano()
	fmt.Printf("Part 2: %v, Time: %fms\n", part2, float32(end-start)/1000.0/1000.0)
}
