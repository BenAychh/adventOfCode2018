package main

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/mgutz/ansi"
)

type lowHighs struct {
	x lowHigh
	y lowHigh
}

type lowHigh struct {
	low  int
	high int
}

type groundType int
type undergroundColumn map[int]groundType
type undergroundGrid map[int]undergroundColumn

const (
	sand         groundType = iota
	clay         groundType = iota
	settledWater groundType = iota
	flowingWater groundType = iota
	spring       groundType = iota
)

func getGrid(lines []string) (grid undergroundGrid) {
	bounds := lowHighs{}
	bounds.x.low = 500
	bounds.x.high = 500
	grid = undergroundGrid{}
	grid.assign(500, 0, spring)

	numbers := regexp.MustCompile("[0-9]+")
	for _, line := range lines {
		lhs := getXandYRange(line, numbers)
		for x := lhs.x.low; x <= lhs.x.high; x++ {
			if x < bounds.x.low {
				bounds.x.low = x
			}
			if x > bounds.x.high {
				bounds.x.high = x
			}
			for y := lhs.y.low; y <= lhs.y.high; y++ {
				if y > bounds.y.high {
					bounds.y.high = y
				}
				grid.assign(x, y, clay)
			}
		}
	}
	for x := bounds.x.low - 2; x <= bounds.x.high+2; x++ {
		for y := 0; y <= bounds.y.high+1; y++ {
			if !grid.hasType(x, y) {
				grid.assign(x, y, sand)
			}
		}
	}
	return
}

func getXandYRange(line string, r *regexp.Regexp) (lh lowHighs) {
	split := strings.Split(line, ", ")
	if split[0][0] == 'x' {
		lh.x = getLowAndHighAsInt(split[0], r)
		lh.y = getLowAndHighAsInt(split[1], r)
	} else {
		lh.y = getLowAndHighAsInt(split[0], r)
		lh.x = getLowAndHighAsInt(split[1], r)
	}
	return
}

func getLowAndHighAsInt(input string, r *regexp.Regexp) (lh lowHigh) {
	split := r.FindAllString(input, -1)
	low, _ := strconv.Atoi(split[0])
	high := low
	if len(split) == 2 {
		high, _ = strconv.Atoi(split[1])
	}
	return lowHigh{
		low:  low,
		high: high,
	}
}

func printGrid(grid undergroundGrid) {
	blue := ansi.ColorFunc("white:blue")
	red := ansi.ColorFunc("white:red")
	yellow := ansi.ColorFunc("black:yellow")
	var sb strings.Builder
	for _, y := range grid.sortedKeys() {
		row := grid[y]
		for _, x := range row.sortedKeys() {
			cell := row[x]
			switch cell {
			case sand:
				sb.WriteString(yellow("."))
			case clay:
				sb.WriteString(red("#"))
			case spring:
				sb.WriteString(blue("+"))
			case settledWater:
				sb.WriteString(blue("~"))
			case flowingWater:
				sb.WriteString(blue("|"))
			}
		}
		sb.WriteString("\n")
	}
	fmt.Println(sb.String())
}

func (u undergroundGrid) assign(x, y int, gt groundType) {
	row, ok := u[y]
	if ok {
		cell, ok := row[x]
		if ok {
			if (gt != clay && cell == clay) || cell == spring {
				imageGrid(u, "error", 2)
				log.Fatal(x, y, " cannot overwrite clay or spring ", cell)
			} else {
				row[x] = gt
			}
		} else {
			row[x] = gt
		}
	} else {
		u[y] = map[int]groundType{
			x: gt,
		}
	}
}

func (u undergroundGrid) hasType(x, y int) bool {
	row, ok := u[y]
	if ok {
		_, ok := row[x]
		return ok
	}
	return false
}

func (u undergroundGrid) getType(x, y int) groundType {
	if u.hasType(x, y) {
		return u[y][x]
	}
	imageGrid(u, "error", 2)
	log.Fatal("type not assigned at ", x, y)
	return sand
}

func (u undergroundGrid) sortedKeys() (keys []int) {
	keys = []int{}
	for key := range u {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	return
}

func (u undergroundColumn) sortedKeys() (keys []int) {
	keys = []int{}
	for key := range u {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	return
}
