package main

import (
	"fmt"
	"strings"
	"time"
	"unicode"

	"github.com/benaychh/aocutils"
)

func main() {
	data := aocutils.LoadFileAsStringArray("day5.data")[0]

	start := time.Now().UnixNano()
	polymer := collapsePolymer(data)
	part1 := len(polymer)
	end := time.Now().UnixNano()
	fmt.Printf("Part 1: %d, Time: %fms\n", part1, float32(end-start)/1000.0/1000.0)

	start = time.Now().UnixNano()
	unit, part2 := getShortestPolymerAfterCleaning(polymer)
	end = time.Now().UnixNano()
	fmt.Printf("Part 1: %d, (Unit: %v) Time: %fms\n", part2, string(unit), float32(end-start)/1000.0/1000.0)
}

func getShortestPolymerAfterCleaning(data string) (unit rune, count int) {
	lowerLetterMap := map[rune]int{}
	for _, letter := range data {
		_, ok := lowerLetterMap[letter]
		if unicode.IsLower(letter) && !ok {
			result := len(removeUnitAndCollapse(letter, data))
			lowerLetterMap[letter] = result
		}
	}

	count = len(data)

	for letter, letterCount := range lowerLetterMap {
		if letterCount < count {
			count = letterCount
			unit = letter
		}
	}
	return
}

func removeUnitAndCollapse(lowerLetter rune, data string) string {
	upperLetter := unicode.ToUpper(lowerLetter)
	cleanedPolymer := strings.Replace(strings.Replace(data, string(lowerLetter), "", -1), string(upperLetter), "", -1)
	return collapsePolymer(cleanedPolymer)
}

func collapsePolymer(data string) string {
	polymerAsRunes := aocutils.ConvertStringToRuneArray(data)
	i := 0
	for i < len(polymerAsRunes) {
		current := polymerAsRunes[i]
		if i+1 == len(polymerAsRunes) {
			return string(polymerAsRunes)
		}
		next := polymerAsRunes[i+1]

		if current == getOpposite(next) {
			polymerAsRunes = append(polymerAsRunes[:i], polymerAsRunes[i+2:]...)
			i = max(i-1, 0)
		} else {
			i++
		}
	}
	return string(polymerAsRunes)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getOpposite(r rune) rune {
	capital := unicode.ToUpper(r)
	if r == capital {
		return unicode.ToLower(r)
	}
	return capital
}
