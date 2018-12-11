package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"time"

	"github.com/BenAychh/aocutils"
)

func main() {
	lines := aocutils.LoadFileAsStringArray("day10.data")
	data := []movingPoint{}
	for _, line := range lines {
		data = append(data, processLine(line))
	}

	start := time.Now().UnixNano()
	part1 := getMessage(data)
	end := time.Now().UnixNano()
	fmt.Printf("Part 1: %v, Time: %fms\n", part1, float32(end-start)/1000.0/1000.0)
}

type movingPoint struct {
	x  int
	y  int
	dx int
	dy int
}

func processLine(line string) movingPoint {
	re := regexp.MustCompile("-?[0-9]+")
	data := []int{}
	for _, value := range re.FindAllString(line, -1) {
		i, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, i)
	}

	return movingPoint{
		x:  data[0],
		y:  data[1],
		dx: data[2],
		dy: data[3],
	}
}
