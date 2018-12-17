package main

import (
	"fmt"

	"github.com/benaychh/aocutils"
)

func getLastCart(tracks [][]rune, carts map[int]map[int]cart, print bool) (lastCartPosition *aocutils.Vertex) {
	i := 0
	for ; lastCartPosition == nil; i++ {
		tickAndClearTracks(tracks, carts, i)
		lastCartPosition = getCartIfLastOne(carts)
		if print {
			printTracksAndCarts(tracks, carts)
		}
	}
	fmt.Println("totalTicks: ", i)
	return
}

func getCartIfLastOne(carts map[int]map[int]cart) *aocutils.Vertex {
	if len(carts) > 1 {
		return nil
	}
	for y, row := range carts {
		if len(row) > 1 {
			return nil
		}
		for x := range row {
			return &aocutils.Vertex{
				X: x,
				Y: y,
			}
		}
	}
	return nil
}

func tickAndClearTracks(tracks [][]rune, carts map[int]map[int]cart, time int) *aocutils.Vertex {
	for _, y := range getSortedYKeys(carts) {
		for _, x := range getSortedXKeys(carts[y]) {
			currentCart := carts[y][x]
			if currentCart.completedTick < time {
				vertex, newCart := getNewCartPosition(x, y, currentCart, tracks, time)
				newRow, newRowExists := carts[vertex.Y]
				if newRowExists {
					_, cellExists := newRow[vertex.X]
					if cellExists {
						if len(carts[vertex.Y]) == 1 {
							delete(carts, vertex.Y)
						} else {
							delete(carts[vertex.Y], vertex.X)
						}
						if len(carts[y]) == 1 {
							delete(carts, y)
							break
						} else {
							delete(carts[y], x)
						}
						continue
					} else {
						newRow[vertex.X] = newCart
					}
				} else {
					carts[vertex.Y] = map[int]cart{
						vertex.X: newCart,
					}
				}
				if len(carts[y]) == 1 {
					delete(carts, y)
				} else {
					delete(carts[y], x)
				}
			}
		}
	}
	return nil
}
