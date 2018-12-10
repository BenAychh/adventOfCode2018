package main

import (
	"container/ring"
)

func getHighestScore(playerCount, lastMarble int) int {
	r := ring.New(1)
	r.Value = 0
	players := make([]int, playerCount)
	for i := 1; i <= lastMarble; i++ {
		if i%23 == 0 {
			removedValue := r.Move(-7).Value.(int)
			r = r.Move(-8)
			r.Unlink(1)
			players[i%playerCount] += i + removedValue
			r = r.Move(1)
		} else {
			new := ring.New(1)
			new.Value = i
			r = r.Move(1)
			r.Link(new)
			r = r.Next()
		}
	}
	largest := -1
	for _, value := range players {
		if value > largest {
			largest = value
		}
	}
	return largest
}
