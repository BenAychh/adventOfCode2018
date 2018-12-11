package main

import (
	"aocutils"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	data := aocutils.LoadArrayOfStringsFromTextFile("day4.data")
	start := time.Now().UnixNano()
	part1Guard, part1Minute := sleepiestGuardAndSleepiestMinute(data)
	end := time.Now().UnixNano()
	fmt.Printf("Part 1: %d x %d = %d, time: %f\n", part1Guard, part1Minute, part1Guard*part1Minute, float32(end-start)/1000.0/1000.0)

	start = time.Now().UnixNano()
	part2Guard, part2Minute := sleepiestMinute(data)
	end = time.Now().UnixNano()
	fmt.Printf("Part 2: %d x %d = %d, time: %f\n", part2Guard, part2Minute, part2Guard*part2Minute, float32(end-start)/1000.0/1000.0)
}

func sleepiestGuardAndSleepiestMinute(data []string) (int, int) {
	sorted := sortData(data)
	guards := assignToGuards(sorted)
	sleepiestGuard := getSleepiestGuard(guards)
	sleepiestMinute, _ := getSleepiestMinute(guards[sleepiestGuard])
	return sleepiestGuard, sleepiestMinute
}

func sleepiestMinute(data []string) (sleepiestGuard, sleepiestMinute int) {
	maxCount := 0
	sorted := sortData(data)
	guards := assignToGuards(sorted)
	for guard, dataPoints := range guards {
		minute, count := getSleepiestMinute(dataPoints)
		if count > maxCount {
			sleepiestGuard = guard
			sleepiestMinute = minute
			maxCount = count
		}
	}
	return
}

func sortData(data []string) []string {
	sort.Strings(data)
	return data
}

type dataPoint struct {
	time     time.Time
	sleeping bool
}

func getSleepiestGuard(guards map[int][]dataPoint) (sleepiestGuard int) {
	maxTotalMinutes := 0
	maxGuard := -1
	for guard, dataPoints := range guards {
		sum := 0
		for i := 0; i < len(dataPoints); i += 2 {
			sum += int(dataPoints[i+1].time.Sub(dataPoints[i].time).Minutes())
		}
		if sum > maxTotalMinutes {
			maxTotalMinutes = sum
			maxGuard = guard
		}
	}
	return maxGuard
}

func getSleepiestMinute(dataPoints []dataPoint) (int, int) {
	minuteMap := map[int]int{}
	maxMinute := -1
	maxCount := 0
	for i := 0; i < len(dataPoints); i += 2 {
		for j := dataPoints[i].time.Minute(); j < dataPoints[i+1].time.Minute(); j++ {
			currentCount, ok := minuteMap[j]
			if ok {
				minuteMap[j] = currentCount + 1
				if minuteMap[j] > maxCount {
					maxMinute = j
					maxCount = minuteMap[j]
				}
			} else {
				minuteMap[j] = 1
			}
		}
	}
	return maxMinute, maxCount
}

func assignToGuards(data []string) (guards map[int][]dataPoint) {
	guards = map[int][]dataPoint{}
	currentGuard := -1
	for _, line := range data {
		if strings.Contains(line, "#") {
			currentGuard, _ = getGuard(line)
			_, ok := guards[currentGuard]
			if !ok {
				guards[currentGuard] = []dataPoint{}
			}
		} else {
			currentArray := guards[currentGuard]
			guards[currentGuard] = append(currentArray, getDataPointFromLine(line))
		}
	}
	return
}

func getGuard(line string) (int, error) {
	splits := strings.Split(line, " ")
	id := strings.Replace(splits[3], "#", "", -1)
	return strconv.Atoi(id)
}

func getDataPointFromLine(line string) dataPoint {
	sleeping := strings.Contains(line, "asleep")
	time, err := time.Parse("[2006-01-02 15:04", strings.Split(line, "]")[0])
	if err != nil {
		log.Fatal(err)
	}
	return dataPoint{time, sleeping}
}
