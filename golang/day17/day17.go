package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"time"

	"github.com/benaychh/aocutils"
)

func main() {
	data := aocutils.LoadFileAsStringArray("day17.data")
	grid := getGrid(data)
	start := time.Now().UnixNano()
	part1 := countWater(grid, 500, 0)
	end := time.Now().UnixNano()
	fmt.Printf("Part 1: %d, Time: %fms\n", part1, float32(end-start)/1000.0/1000.0)

	start = time.Now().UnixNano()
	part2 := countSettledWater(grid)
	end = time.Now().UnixNano()
	fmt.Printf("Part 2: %d, Time: %fms\n", part2, float32(end-start)/1000.0/1000.0)

	imageGrid(grid, "result", 2)
}

func imageGrid(grid undergroundGrid, name string, pixelSize int) {
	min := -1
	width := 0
	height := 0
	for y, row := range grid {
		if y > height {
			height = y
		}
		for x := range row {
			if x > width {
				width = x
			}
			if min == -1 || x < min {
				min = x
			}
		}
	}
	width = (width - min) * pixelSize
	height = (height + 1) * pixelSize

	img := drawImage(grid, width, height, min, pixelSize)

	f, err := os.Create(name + ".png")
	if err != nil {
		log.Fatal(err)
	}
	err = png.Encode(f, img)
	if err != nil {
		log.Fatal(err)
	}
}

func drawImage(grid undergroundGrid, width, height, min, pixelSize int) *image.RGBA {
	clayColor := color.RGBA{252, 102, 53, 255}
	sandColor := color.RGBA{244, 241, 66, 255}
	stillWaterColor := color.RGBA{66, 152, 244, 255}
	flowingWaterColor := color.RGBA{8, 0, 178, 255}

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for y, row := range grid {
		for x, cell := range row {
			c := stillWaterColor
			switch cell {
			case clay:
				c = clayColor
			case sand:
				c = sandColor
			case flowingWater:
				c = flowingWaterColor
			}
			for i := (x - min) * pixelSize; i < (x-min+1)*pixelSize; i++ {
				for j := y * pixelSize; j < (y+1)*pixelSize; j++ {
					img.Set(i, j, c)
				}
			}
		}
	}
	return img
}
