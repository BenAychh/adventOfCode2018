package main

import (
	"sort"
)

type pathPiece struct {
	x        int
	y        int
	distance int
	cameFrom *pathPiece
}

func getAttackableEnemies(entities entityMap, enemy magicalCreature, x, y int) (e *entity, aX, aY int) {
	yDeltas := [3]int{-1, 0, 1}
	lowestHealth := 300
	for _, yDelta := range yDeltas {
		if yDelta != 0 {
			if entities.speciesAtLocation(x, y+yDelta, enemy) {
				temp := entities.getAt(x, y+yDelta)
				if temp.hitPoints < lowestHealth {
					e = temp
					lowestHealth = e.hitPoints
					aX = x
					aY = y + yDelta
				}
			}
		} else {
			xDeltas := [2]int{-1, 1}
			for _, xDelta := range xDeltas {
				if entities.speciesAtLocation(x+xDelta, y, enemy) {
					temp := entities.getAt(x+xDelta, y)
					if temp.hitPoints < lowestHealth {
						e = temp
						lowestHealth = e.hitPoints
						aX = x + xDelta
						aY = y
					}
				}
			}
		}
	}
	return e, aX, aY
}

func getNextMove(grid [][]bool, entities entityMap, entity *entity, currentX, currentY int) *pathPiece {
	me := &pathPiece{
		x:        currentX,
		y:        currentY,
		distance: 0,
		cameFrom: nil,
	}
	yDeltas := [3]int{-1, 0, 1}
	xDeltas := [3]int{-1, 1}
	processedSquares := map[int]map[int]bool{}
	queue := []*pathPiece{me}
	possibleTargets := []*pathPiece{}
	for len(queue) != 0 {
		current := queue[0]
		if len(possibleTargets) > 0 && current.distance+1 > possibleTargets[0].distance {
			break
		}
		queue = queue[1:]
		for _, yDelta := range yDeltas {
			if yDelta != 0 {
				if entities.speciesAtLocation(current.x, current.y+yDelta, entity.enemy) {
					possibleTargets = append(possibleTargets, &pathPiece{
						x:        current.x,
						y:        current.y + yDelta,
						distance: current.distance + 1,
						cameFrom: current,
					})
				}
				queue = checkAndAppendToQueueIfNeeded(grid, entities, queue, processedSquares, current, 0, yDelta)
			} else {
				for _, xDelta := range xDeltas {
					if entities.speciesAtLocation(current.x+xDelta, current.y, entity.enemy) {
						possibleTargets = append(possibleTargets, &pathPiece{
							x:        current.x + xDelta,
							y:        current.y,
							distance: current.distance + 1,
							cameFrom: current,
						})
					}
					queue = checkAndAppendToQueueIfNeeded(grid, entities, queue, processedSquares, current, xDelta, yDelta)
				}
			}
		}
	}
	return getFirstTargetByReadingOrder(possibleTargets)
}

func getFirstTargetByReadingOrder(possibleTargets []*pathPiece) *pathPiece {
	if len(possibleTargets) > 0 {
		sort.Slice(possibleTargets, func(i, j int) bool {
			if possibleTargets[i].y < possibleTargets[j].y {
				return true
			}
			if possibleTargets[i].y > possibleTargets[j].y {
				return false
			}
			return possibleTargets[i].x < possibleTargets[j].x
		})
		return followBackwards(possibleTargets[0])
	}
	return nil
}

func checkAndAppendToQueueIfNeeded(grid [][]bool, entities entityMap, queue []*pathPiece, processedSquares map[int]map[int]bool, current *pathPiece, xDelta, yDelta int) []*pathPiece {
	x := current.x + xDelta
	y := current.y + yDelta
	alreadyInQueue := alreadyInQueue(queue, x, y)
	alreadyProcessed := alreadyProcessed(processedSquares, x, y)
	if !alreadyInQueue && !alreadyProcessed && grid[y][x] && !entities.atLocation(x, y) {
		addToAlreadyProcessed(processedSquares, x, y)
		return append(queue, &pathPiece{
			x:        x,
			y:        y,
			distance: current.distance + 1,
			cameFrom: current,
		})
	}
	return queue
}

func alreadyInQueue(queue []*pathPiece, x, y int) (alreadyInQueue bool) {
	for _, gs := range queue {
		if gs.x == x && gs.y == y {
			alreadyInQueue = true
			break
		}
	}
	return
}

func alreadyProcessed(processedSquares map[int]map[int]bool, x, y int) bool {
	col, ok := processedSquares[x]
	if !ok {
		return false
	}
	_, ok = col[y]
	return ok
}

func addToAlreadyProcessed(processedSquares map[int]map[int]bool, x, y int) {
	col, ok := processedSquares[x]
	if ok {
		col[y] = true
	} else {
		processedSquares[x] = map[int]bool{
			y: true,
		}
	}
}

func followBackwards(currentSquare *pathPiece) *pathPiece {
	for currentSquare.cameFrom.cameFrom != nil {
		currentSquare = currentSquare.cameFrom
	}
	return currentSquare
}
