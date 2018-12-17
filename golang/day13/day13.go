package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/benaychh/aocutils"
	"github.com/mgutz/ansi"
)

func main() {
	rawData := aocutils.LoadFileAsStringArray("day13.data")
	start := time.Now().UnixNano()
	tracks, carts := makeTracks(rawData)
	part1 := *getFirstCrash(tracks, carts, false)
	end := time.Now().UnixNano()
	fmt.Printf("Part 1: %v, Time: %fms\n", part1, float32(end-start)/1000.0/1000.0)

	start = time.Now().UnixNano()
	tracks, carts = makeTracks(rawData)
	part2 := *getLastCart(tracks, carts, false)
	end = time.Now().UnixNano()
	fmt.Printf("Part 2: %v, Time: %fms\n", part2, float32(end-start)/1000.0/1000.0)
}

type direction int

type cardinalDirection int

const (
	left     direction = iota
	straight direction = iota
	right    direction = iota
)

const (
	north cardinalDirection = iota
	south cardinalDirection = iota
	east  cardinalDirection = iota
	west  cardinalDirection = iota
)

type cart struct {
	cardinal      cardinalDirection
	nextTurn      direction
	completedTick int
	crashed       bool
}

var cartDirections = map[rune]cardinalDirection{
	'^': north,
	'v': south,
	'>': east,
	'<': west,
}

var directionsCart = map[cardinalDirection]rune{
	north: '^',
	south: 'v',
	east:  '>',
	west:  '<',
}

func makeTracks(lines []string) (tracks [][]rune, carts map[int]map[int]cart) {
	tracks = [][]rune{}
	carts = map[int]map[int]cart{}
	for y, line := range lines {
		row := []rune{}
		for x, r := range line {
			currentDirection, isCart := cartDirections[r]
			if isCart {
				row = append(row, getTrackPiece(currentDirection))
				cartColumn, rowExists := carts[y]
				if rowExists {
					cartColumn[x] = cart{
						cardinal: currentDirection,
					}
				} else {
					carts[y] = map[int]cart{
						x: cart{
							cardinal: currentDirection,
						},
					}
				}
				continue
			}
			row = append(row, r)
		}
		tracks = append(tracks, row)
	}
	return
}

func getTrackPiece(c cardinalDirection) rune {
	if c == north || c == south {
		return '|'
	}
	return '-'
}

func printTracks(tracks [][]rune) {
	for _, row := range tracks {
		fmt.Println(string(row))
	}
}

var cartColor = ansi.ColorFunc("white+h:2")
var crashColor = ansi.ColorFunc("white+h:red:h")

func printTracksAndCarts(tracks [][]rune, carts map[int]map[int]cart) {
	for y, row := range tracks {
		var sb strings.Builder
		for x, cell := range row {
			cartRow, hasCartRow := carts[y]
			if hasCartRow {
				cartCell, hasCartCell := cartRow[x]
				if hasCartCell {
					if cartCell.crashed {
						sb.WriteString(crashColor("X"))
					} else {
						sb.WriteString(cartColor(string(directionsCart[cartCell.cardinal])))
					}
					continue
				}
			}
			sb.WriteRune(cell)
		}
		fmt.Println(sb.String())
	}
	fmt.Println("")
}
