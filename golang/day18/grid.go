package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/mgutz/ansi"
)

type landType int
type landRow map[int]landType
type landGrid map[int]landRow

const (
	open       landType = iota
	trees      landType = iota
	lumberyard landType = iota
)

func getGrid(lines []string) (grid landGrid) {
	grid = landGrid{}

	for y, line := range lines {
		for x, r := range line {
			switch r {
			case '.':
				grid.assign(x, y, open)
			case '|':
				grid.assign(x, y, trees)
			case '#':
				grid.assign(x, y, lumberyard)
			}
		}
	}

	return
}

func printGrid(grid landGrid) {
	openColor := ansi.ColorFunc("white:yellow")
	treeColor := ansi.ColorFunc("white:green")
	lumberyardColor := ansi.ColorFunc("white:brown")
	var sb strings.Builder
	for _, y := range grid.sortedKeys() {
		row := grid[y]
		for _, x := range row.sortedKeys() {
			cell := row[x]
			switch cell {
			case open:
				sb.WriteString(openColor("."))
			case trees:
				sb.WriteString(treeColor("|"))
			case lumberyard:
				sb.WriteString(lumberyardColor("#"))
			}
		}
		sb.WriteString("\n")
	}
	fmt.Println(sb.String())
}

func (l landGrid) assign(x, y int, gt landType) {
	row, ok := l[y]
	if ok {
		row[x] = gt
	} else {
		l[y] = map[int]landType{
			x: gt,
		}
	}
}

func (l landGrid) getType(x, y int) landType {
	row, ok := l[y]
	if ok {
		cell, ok := row[x]
		if ok {
			return cell
		}
	}
	return open
}

func (u landGrid) sortedKeys() (keys []int) {
	keys = []int{}
	for key := range u {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	return
}

func (c landRow) sortedKeys() (keys []int) {
	keys = []int{}
	for key := range c {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	return
}

func imageGrid(grid landGrid, name string, pixelSize int) {
	height := 0
	width := 0
	for y, row := range grid {
		if y > height {
			height = y
		}
		for x := range row {
			if x > width {
				width = x
			}
		}
	}
	width = width * pixelSize
	height = height * pixelSize

	img := drawImage(grid, width, height, pixelSize)

	f, err := os.Create(name + ".png")
	if err != nil {
		log.Fatal(err)
	}
	err = png.Encode(f, img)
	if err != nil {
		log.Fatal(err)
	}
}

func drawImage(grid landGrid, width, height, pixelSize int) *image.RGBA {
	openColor := color.RGBA{255, 255, 204, 255}
	treeColor := color.RGBA{0, 204, 0, 255}
	lumberyardColor := color.RGBA{153, 51, 0, 255}

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for y, row := range grid {
		for x, cell := range row {
			c := openColor
			switch cell {
			case trees:
				c = treeColor
			case lumberyard:
				c = lumberyardColor
			}
			for i := x * pixelSize; i < (x+1)*pixelSize; i++ {
				for j := y * pixelSize; j < (y+1)*pixelSize; j++ {
					img.Set(i, j, c)
				}
			}
		}
	}
	return img
}
