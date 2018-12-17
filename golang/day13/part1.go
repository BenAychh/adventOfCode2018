package main

import (
	"fmt"
	"log"
	"sort"

	"github.com/benaychh/aocutils"
)

func getFirstCrash(tracks [][]rune, carts map[int]map[int]cart, print bool) (crashSite *aocutils.Vertex) {
	i := 0
	for ; crashSite == nil; i++ {
		crashSite = tick(tracks, carts, i)
		if print {
			printTracksAndCarts(tracks, carts)
		}
	}
	fmt.Println("totalTicks: ", i)
	return
}

func tick(tracks [][]rune, carts map[int]map[int]cart, time int) *aocutils.Vertex {
	for _, y := range getSortedYKeys(carts) {
		row := carts[y]
		for _, x := range getSortedXKeys(row) {
			currentCart := row[x]
			if currentCart.completedTick < time {
				vertex, newCart := getNewCartPosition(x, y, currentCart, tracks, time)
				newRow, newRowExists := carts[vertex.Y]
				if newRowExists {
					_, cellExists := newRow[vertex.X]
					if cellExists {
						newCart.crashed = true
						newRow[vertex.X] = newCart
						delete(row, x)
						return &vertex
					}
					newRow[vertex.X] = newCart
				} else {
					carts[vertex.Y] = map[int]cart{
						vertex.X: newCart,
					}
				}
				delete(row, x)
			}
		}
	}
	return nil
}

func getSortedYKeys(m map[int]map[int]cart) (sortedKeys []int) {
	sortedKeys = []int{}
	for key := range m {
		sortedKeys = append(sortedKeys, key)
	}
	sort.Ints(sortedKeys)
	return sortedKeys
}

func getSortedXKeys(m map[int]cart) (sortedKeys []int) {
	sortedKeys = []int{}
	for key := range m {
		sortedKeys = append(sortedKeys, key)
	}
	sort.Ints(sortedKeys)
	return sortedKeys
}

func getNewCartPosition(currentX, currentY int, cart cart, tracks [][]rune, time int) (aocutils.Vertex, cart) {
	vertex := aocutils.Vertex{
		X: currentX,
		Y: currentY,
	}
	cart.completedTick = time
	switch cart.cardinal {
	case north:
		return handleNorth(vertex, cart, tracks)
	case east:
		return handleEast(vertex, cart, tracks)
	case south:
		return handleSouth(vertex, cart, tracks)
	case west:
		return handleWest(vertex, cart, tracks)
	}
	log.Fatalf("Invalid Cart Direction: (%d, %d) %+v", currentX, currentY, cart)
	return vertex, cart
}

func handleNorth(vertex aocutils.Vertex, cart cart, tracks [][]rune) (aocutils.Vertex, cart) {
	vertex.Y = vertex.Y - 1
	switch tracks[vertex.Y][vertex.X] {
	case '\\':
		cart.cardinal = west
	case '/':
		cart.cardinal = east
	case '+':
		switch cart.nextTurn {
		case left:
			cart.cardinal = west
			cart.nextTurn = straight
		case straight:
			cart.nextTurn = right
		case right:
			cart.cardinal = east
			cart.nextTurn = left
		}
	case ' ':
		log.Fatal(vertex, cart)
	}
	return vertex, cart
}

func handleEast(vertex aocutils.Vertex, cart cart, tracks [][]rune) (aocutils.Vertex, cart) {
	vertex.X = vertex.X + 1
	switch tracks[vertex.Y][vertex.X] {
	case '\\':
		cart.cardinal = south
	case '/':
		cart.cardinal = north
	case '+':
		switch cart.nextTurn {
		case left:
			cart.cardinal = north
			cart.nextTurn = straight
		case straight:
			cart.nextTurn = right
		case right:
			cart.cardinal = south
			cart.nextTurn = left
		}
	}
	return vertex, cart
}

func handleSouth(vertex aocutils.Vertex, cart cart, tracks [][]rune) (aocutils.Vertex, cart) {
	vertex.Y = vertex.Y + 1
	switch tracks[vertex.Y][vertex.X] {
	case '\\':
		cart.cardinal = east
	case '/':
		cart.cardinal = west
	case '+':
		switch cart.nextTurn {
		case left:
			cart.cardinal = east
			cart.nextTurn = straight
		case straight:
			cart.nextTurn = right
		case right:
			cart.cardinal = west
			cart.nextTurn = left
		}
	}
	return vertex, cart
}

func handleWest(vertex aocutils.Vertex, cart cart, tracks [][]rune) (aocutils.Vertex, cart) {
	vertex.X = vertex.X - 1
	switch tracks[vertex.Y][vertex.X] {
	case '\\':
		cart.cardinal = north
	case '/':
		cart.cardinal = south
	case '+':
		switch cart.nextTurn {
		case left:
			cart.cardinal = south
			cart.nextTurn = straight
		case straight:
			cart.nextTurn = right
		case right:
			cart.cardinal = north
			cart.nextTurn = left
		}
	}
	return vertex, cart
}
