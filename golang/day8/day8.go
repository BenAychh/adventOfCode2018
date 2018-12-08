package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/BenAychh/aocutils"
)

func main() {
	line := aocutils.LoadFileAsStringArray("day8.data")[0]
	data := []int{}
	for _, intAsString := range strings.Split(line, " ") {
		asInt, conversionError := strconv.Atoi(intAsString)
		if conversionError != nil {
			log.Fatal(conversionError)
		}
		data = append(data, asInt)
	}

	start := time.Now().UnixNano()
	rootNode, sum := metadataSum(data)
	end := time.Now().UnixNano()
	fmt.Printf("Part 1: %v, Time: %fms\n", sum, float32(end-start)/1000.0/1000.0)

	start = time.Now().UnixNano()
	cSum := complicatedSum(rootNode)
	end = time.Now().UnixNano()
	fmt.Printf("Part 2: %v, Time: %fms\n", cSum, float32(end-start)/1000.0/1000.0)

}
