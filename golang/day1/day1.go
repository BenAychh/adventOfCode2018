package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	jsonFile, err := os.Open("day1.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result []string
	_ = json.Unmarshal(byteValue, &result)

	var resultAsInts = convertToIntegers(result)

	part1 := processor(resultAsInts)
	start := time.Now().UnixNano()
	part2 := repeatFinder(resultAsInts)
	end := time.Now().UnixNano()
	fmt.Println(part1)
	fmt.Println(part2)
	fmt.Println((end - start) / 1000 / 1000)
}

func convertToIntegers(data []string) (ints []int) {
	ints = []int{}
	for _, str := range data {
		int, err := strconv.Atoi(strings.Replace(str, "+", "", -1))
		if err != nil {
			log.Fatal(err)
		}
		ints = append(ints, int)
	}
	return ints
}

func processor(data []int) (sum int) {
	sum = 0

	for _, i := range data {
		sum += i
	}
	return
}

func repeatFinder(data []int) int {
	sum := 0
	hasRepeated := false
	freqMap := make(map[int]bool)
	freqMap[0] = true

	for !hasRepeated {
		for _, i := range data {
			sum += i
			_, ok := freqMap[sum]
			if ok {
				hasRepeated = true
				return sum
			} else {
				freqMap[sum] = true
			}
		}
	}
	return 0
}
