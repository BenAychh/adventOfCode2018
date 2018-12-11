package main

import (
	"fmt"
	"time"

	"github.com/BenAychh/aocutils"
)

func main() {
	data := aocutils.LoadFileAsStringArray("day7.data")

	start := time.Now().UnixNano()
	order := getOrder(data)
	end := time.Now().UnixNano()
	fmt.Printf("Part 1: %v, Time: %fms\n", order, float32(end-start)/1000.0/1000.0)

	start = time.Now().UnixNano()
	totalTime := getTotalTime(data, 5, 60)
	end = time.Now().UnixNano()
	fmt.Printf("Part 2: %d, Time: %fms\n", totalTime, float32(end-start)/1000.0/1000.0)
}
