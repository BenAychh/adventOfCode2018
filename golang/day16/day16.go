package main

import (
	"aocutils"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"time"
)

func main() {
	rawData := aocutils.LoadArrayOfStringsFromTextFile("part1.data")
	data := processData(rawData)
	start := time.Now().UnixNano()
	part1 := findAllInterpretationsWithOver2Matches(data)
	end := time.Now().UnixNano()
	fmt.Printf("Part 1: %d, Time: %fms\n", part1, float32(end-start)/1000.0/1000.0)

	testProgram := aocutils.LoadArrayOfStringsFromTextFile("part2.data")
	start = time.Now().UnixNano()
	part2 := processTestProgram(data, testProgram)
	end = time.Now().UnixNano()
	fmt.Printf("Part 2: %d, Time: %fms\n", part2, float32(end-start)/1000.0/1000.0)
}

type interpretation struct {
	before   [4]int
	register [4]int
	after    [4]int
}

func processData(lines []string) (interpretations []interpretation) {
	interpretations = []interpretation{}
	re := regexp.MustCompile("-?[0-9]+")
	for i := 0; i < len(lines); i += 4 {
		interp := interpretation{}
		interp.before = processLine(lines[i], re)
		interp.register = processLine(lines[i+1], re)
		interp.after = processLine(lines[i+2], re)
		interpretations = append(interpretations, interp)
	}
	return
}

func processLine(line string, regex *regexp.Regexp) (arr [4]int) {
	ints := regex.FindAllString(line, -1)
	if len(ints) != 4 {
		log.Fatalf("too many numbers %v", ints)
	}
	for index, value := range ints {
		i, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal(err)
		}
		arr[index] = i
	}
	return
}
