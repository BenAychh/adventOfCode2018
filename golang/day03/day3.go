package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/benaychh/aocutils"
)

type claim struct {
	id     int
	x      int
	y      int
	width  int
	height int
}

func main() {
	data := aocutils.LoadFileAsStringArray("day3.data")
	start := time.Now().UnixNano()
	part1, overlapMap, claims := overlap(data)
	end := time.Now().UnixNano()
	fmt.Println(part1)
	fmt.Println(float32(end-start) / 1000 / 1000)

	start = time.Now().UnixNano()
	part2 := noOverLap(overlapMap, claims)
	end = time.Now().UnixNano()
	fmt.Println(part2)
	fmt.Println(float32(end-start) / 1000 / 1000)
}

func noOverLap(overlapMap map[string]int, claims []claim) int {
	for _, claim := range claims {
		isShared := checkSharedSpaces(claim, overlapMap)
		if !isShared {
			return claim.id
		}
	}
	return -1
}

func overlap(data []string) (int, map[string]int, []claim) {
	claims := convertToClaimsArray(data)
	overlapMap := map[string]int{}
	for _, claim := range claims {
		claimSpaces(claim, overlapMap)
	}

	multiples := 0
	for _, value := range overlapMap {
		if value > 1 {
			multiples++
		}
	}

	return multiples, overlapMap, claims
}

func checkSharedSpaces(claim claim, grid map[string]int) bool {
	for x := claim.x; x < claim.x+claim.width; x++ {
		for y := claim.y; y < claim.y+claim.height; y++ {
			key := fmt.Sprintf("%d,%d", x, y)
			ids := grid[key]
			if ids > 1 {
				return true
			}
		}
	}
	return false
}

func claimSpaces(claim claim, grid map[string]int) {
	for x := claim.x; x < claim.x+claim.width; x++ {
		for y := claim.y; y < claim.y+claim.height; y++ {
			claimSpace(x, y, claim.id, grid)
		}
	}
}

func claimSpace(x, y, id int, grid map[string]int) {
	key := fmt.Sprintf("%d,%d", x, y)
	value, ok := grid[key]
	if ok {
		grid[key] = value + 1
	} else {
		grid[key] = 1
	}
}

func convertToClaimsArray(data []string) (claims []claim) {
	for _, datum := range data {
		claims = append(claims, decodeClaim(datum))
	}
	return
}

func decodeClaim(line string) (claim claim) {
	pieces := strings.Split(strings.Replace(line, "#", "", -1), " ")
	claim.id, _ = strconv.Atoi(pieces[0])
	cleanCoordsString := strings.Replace(pieces[2], ":", "", -1)
	coordsStringsArray := strings.Split(cleanCoordsString, ",")
	claim.x, _ = strconv.Atoi(coordsStringsArray[0])
	claim.y, _ = strconv.Atoi(coordsStringsArray[1])
	dimensionsStringsArray := strings.Split(pieces[3], "x")
	claim.width, _ = strconv.Atoi(dimensionsStringsArray[0])
	claim.height, _ = strconv.Atoi(dimensionsStringsArray[1])
	return
}
