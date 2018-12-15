package main

import (
	"fmt"
	"strconv"
)

func getPrevious5(sequence []int, print bool) int {
	recipes := []int{3, 7}
	elf1 := 0
	elf2 := 1
	lastN := make([]int, len(sequence))
	done := false
	for !done {
		if print {
			printLine(elf1, elf2, recipes)
		}
		sum := recipes[elf1] + recipes[elf2]
		sumAsString := strconv.Itoa(sum)
		for _, r := range sumAsString {
			recipes = append(recipes, int(r-'0'))
			lastN = append(lastN[1:], int(r-'0'))
			if checkEqual(lastN, sequence) {
				done = true
				break
			}
		}
		elf1 = getNewElfPosition(elf1, recipes[elf1], len(recipes))
		elf2 = getNewElfPosition(elf2, recipes[elf2], len(recipes))
	}
	return len(recipes) - len(sequence)
}

func checkEqual(a []int, b []int) bool {
	if len(a) != len(b) {
		fmt.Println("here????")
		return false
	}
	for i, num := range a {
		if b[i] != num {
			return false
		}
	}
	return true
}
