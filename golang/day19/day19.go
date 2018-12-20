package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"time"

	"github.com/benaychh/aocutils"
)

func main() {
	rawData := aocutils.LoadFileAsStringArray("day19.data")
	data, ipRegister := getInstructionLines(rawData)
	start := time.Now().UnixNano()
	part1 := followInstructionsUntilDone(data, [6]int{0, 0, 0, 0, 0, 0}, ipRegister)
	end := time.Now().UnixNano()
	fmt.Printf("Part 1: %v, Time: %fms\n", part1, float32(end-start)/1000.0/1000.0)

	start = time.Now().UnixNano()
	part2 := followInstructionsUntilDone(data, [6]int{1, 0, 0, 0, 0, 0}, ipRegister)
	end = time.Now().UnixNano()
	fmt.Printf("Part 1: %v, Time: %fms\n", part2, float32(end-start)/1000.0/1000.0)

}

type instructionLines []instructionLine

type instructionLine struct {
	fnName string
	set    [3]int
}

func getInstructionLines(data []string) (lines instructionLines, ip int) {
	numbers := regexp.MustCompile("[0-9]+")
	ipAsString := numbers.FindAllString(data[0], -1)
	ip, err := strconv.Atoi(ipAsString[0])
	if err != nil {
		log.Fatal("Error getting IP ", data[0])
	}
	lines = []instructionLine{}
	for i := 1; i < len(data); i++ {
		datum := data[i]
		il := instructionLine{
			fnName: datum[0:4],
		}

		ints := numbers.FindAllString(datum, -1)
		for index, value := range ints {
			i, err := strconv.Atoi(value)
			if err != nil {
				log.Fatal(err)
			}
			il.set[index] = i
		}

		lines = append(lines, il)
	}

	return
}

func (il instructionLines) hasInstruction(index int) bool {
	return index >= 0 && index < len(il)
}
