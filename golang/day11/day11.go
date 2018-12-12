package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now().UnixNano()
	part1, power := getHighestPower(9445)
	end := time.Now().UnixNano()
	fmt.Printf("Part 1: %v (power: %d), Time: %fms\n", part1, power, float32(end-start)/1000.0/1000.0)

	start = time.Now().UnixNano()
	part2, power, size := getHighestPowerAnySize(9445)
	end = time.Now().UnixNano()
	fmt.Printf("Part 2: %v,%d (power: %d), Time: %fms\n", part2, size, power, float32(end-start)/1000.0/1000.0)
}
