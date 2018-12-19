package main

import (
	"fmt"
	"time"

	"github.com/benaychh/aocutils"
)

func main() {
	data := aocutils.LoadFileAsStringArray("day18.data")
	grid := getGrid(data)
	start := time.Now().UnixNano()
	part1Trees, part1Lumberyards, part1Grid := getResourceValueAfter(grid, 10)
	end := time.Now().UnixNano()
	imageGrid(part1Grid, "part1", 10)
	fmt.Printf("Part 1: %dx%d=%d, Time: %fms\n", part1Trees, part1Lumberyards, part1Trees*part1Lumberyards, float32(end-start)/1000.0/1000.0)

	start = time.Now().UnixNano()
	part2Trees, part2Lumberyards, part2Grid := getResourceValueAfterBigNumber(grid, 1000000000)
	end = time.Now().UnixNano()
	imageGrid(part2Grid, "part2", 10)
	fmt.Printf("Part 2: %dx%d=%d, Time: %fms\n", part2Trees, part2Lumberyards, part2Trees*part2Lumberyards, float32(end-start)/1000.0/1000.0)
}
