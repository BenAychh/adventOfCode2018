package main

import (
	"aocutils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type claim struct {
	id     string
	x      int
	y      int
	width  int
	height int
}

func main() {
	data := aocutils.LoadArrayOfStringsFromTextFile("day3.data")
	start := time.Now().UnixNano()
	part1 := overlap(data)
	end := time.Now().UnixNano()
	fmt.Println(part1)
	fmt.Println(float32(end-start) / 1000 / 1000)

	start = time.Now().UnixNano()
	part2 := noOverLap(data)
	end = time.Now().UnixNano()
	fmt.Println(part2)
	fmt.Println(float32(end-start) / 1000 / 1000)
}

func noOverLap(data []string) string {
	claims := convertToClaimsArray(data)
	overlapMap := map[string][]string{}
	for _, claim := range claims {
		claimSpaces(claim, overlapMap)
	}
	for _, claim := range claims {
		isShared := checkSharedSpaces(claim, overlapMap)
		if !isShared {
			return claim.id
		}
	}
	return ""
}

func overlap(data []string) int {
	claims := convertToClaimsArray(data)
	overlapMap := map[string][]string{}
	for _, claim := range claims {
		claimSpaces(claim, overlapMap)
	}

	multiples := 0
	for _, value := range overlapMap {
		if len(value) > 1 {
			multiples++
		}
	}

	return multiples
}

func checkSharedSpaces(claim claim, grid map[string][]string) bool {
	for x := claim.x; x < claim.x+claim.width; x++ {
		for y := claim.y; y < claim.y+claim.height; y++ {
			key := fmt.Sprintf("%d,%d", x, y)
			ids := grid[key]
			if len(ids) > 1 {
				return true
			}
		}
	}
	return false
}

func claimSpaces(claim claim, grid map[string][]string) {
	for x := claim.x; x < claim.x+claim.width; x++ {
		for y := claim.y; y < claim.y+claim.height; y++ {
			claimSpace(x, y, claim.id, grid)
		}
	}
}

func claimSpace(x, y int, id string, grid map[string][]string) {
	key := fmt.Sprintf("%d,%d", x, y)
	array, ok := grid[key]
	if ok {
		grid[key] = append(array, id)
	} else {
		grid[key] = []string{id}
	}
}

func convertToClaimsArray(data []string) (claims []claim) {
	for _, datum := range data {
		claims = append(claims, decodeClaim(datum))
	}
	return
}

func decodeClaim(line string) (claim claim) {
	pieces := strings.Split(line, " ")
	claim.id = pieces[0]
	cleanCoordsString := strings.Replace(pieces[2], ":", "", -1)
	coordsStringsArray := strings.Split(cleanCoordsString, ",")
	claim.x, _ = strconv.Atoi(coordsStringsArray[0])
	claim.y, _ = strconv.Atoi(coordsStringsArray[1])
	dimensionsStringsArray := strings.Split(pieces[3], "x")
	claim.width, _ = strconv.Atoi(dimensionsStringsArray[0])
	claim.height, _ = strconv.Atoi(dimensionsStringsArray[1])
	return
}
