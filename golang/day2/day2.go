package main

import (
	"aocutils"
	"fmt"
	"time"
)

func main() {
	data := aocutils.LoadArrayOfStringsFromTextFile("day2.data")

	start := time.Now().UnixNano()
	part1 := checksum(data)
	end := time.Now().UnixNano()
	fmt.Println(part1)
	fmt.Println(float32(end-start) / 1000.0 / 1000.0)

	start = time.Now().UnixNano()
	part2 := commonLetters(data)
	end = time.Now().UnixNano()
	fmt.Println(part2)
	fmt.Println(float32(end-start) / 1000.0 / 1000.0)
}

func checksum(data []string) int {
	twos := 0
	threes := 0

	for _, str := range data {
		counts := getCounts(str)
		two, three := get2sAnd3s(counts)
		twos += two
		threes += three
	}
	return twos * threes
}

func commonLetters(data []string) string {
	arrayOfRunesArrays := convertToRunesArrays(data)
	id1, _, diffencesIndexes := getSimilarIds(arrayOfRunesArrays)
	index := diffencesIndexes[0]
	common := append(id1[:index], id1[index+1:]...)
	return string(common)
}

func getSimilarIds(arrayOfRunesArrays [][]rune) (id1 []rune, id2 []rune, differencesIndexes []int) {
	head := arrayOfRunesArrays[0]
	tail := arrayOfRunesArrays[1:]
	for _, runeArray := range tail {
		differencesIndexes = getDifferencesLimitTo2(head, runeArray)
		if len(differencesIndexes) == 1 {
			return head, runeArray, differencesIndexes
		}
	}
	return getSimilarIds(tail)
}

func convertToRunesArrays(data []string) (arrayOfRunesArrays [][]rune) {
	for _, str := range data {
		runeArray := []rune{}
		for _, r := range str {
			runeArray = append(runeArray, r)
		}
		arrayOfRunesArrays = append(arrayOfRunesArrays, runeArray)
	}
	return
}

func getDifferencesLimitTo2(str1, str2 []rune) (differencesIndexes []int) {
	for index, r := range str1 {
		if r != str2[index] {
			differencesIndexes = append(differencesIndexes, index)
		}
		if len(differencesIndexes) > 1 {
			return
		}
	}
	return
}

func getCounts(str string) (counts map[rune]int) {
	counts = make(map[rune]int)
	for _, char := range str {
		count, ok := counts[char]
		if ok {
			counts[char] = count + 1
		} else {
			counts[char] = 1
		}
	}
	return
}

func get2sAnd3s(counts map[rune]int) (twos, threes int) {
	for _, v := range counts {
		if v == 2 {
			twos = 1
		}
		if v == 3 {
			threes = 1
		}
		if twos == 1 && threes == 1 {
			return
		}
	}
	return
}
