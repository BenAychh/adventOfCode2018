package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getNext10(after int, print bool) string {
	recipes := []int{3, 7}
	elf1 := 0
	elf2 := 1
	for i := 0; len(recipes) < after+10; i++ {
		if print {
			printLine(elf1, elf2, recipes)
		}
		sum := recipes[elf1] + recipes[elf2]
		sumAsString := strconv.Itoa(sum)
		for _, r := range sumAsString {
			recipes = append(recipes, int(r-'0'))
		}
		elf1 = getNewElfPosition(elf1, recipes[elf1], len(recipes))
		elf2 = getNewElfPosition(elf2, recipes[elf2], len(recipes))
	}
	var sb strings.Builder
	for _, i := range recipes[after : after+10] {
		sb.WriteString(strconv.Itoa(i))
	}
	return sb.String()
}

func getNewElfPosition(currentPos, currentValue, recipesLength int) int {
	return (currentPos + currentValue + 1) % recipesLength
}

func printLine(elf1, elf2 int, recipes []int) {
	var sb strings.Builder
	for i, score := range recipes {
		if i == elf1 {
			str := fmt.Sprintf("(%d)", score)
			sb.WriteString(str)
		} else if i == elf2 {
			str := fmt.Sprintf("[%d]", score)
			sb.WriteString(str)
		} else {
			str := fmt.Sprintf(" %d ", score)
			sb.WriteString(str)
		}
	}
	fmt.Println(sb.String())
}
