package main

import (
	"fmt"
)

func getRoundsAndRemainingHitpoints(grid [][]bool, entities entityLocations, print bool) (round, hitpoints int) {
	areBothSpeciesStillAlive := bothSpeciesStillAlive(entities)
	for areBothSpeciesStillAlive {
		for _, y := range entities.sortedKeys() {
			row, ok := entities[y]
			if ok {
				for _, x := range row.sortedKeys() {
					entity, ok := row[x]
					if ok {
						wasSomeoneKilled := takeTurn(grid, entities, entity, x, y, round)
						if wasSomeoneKilled {
							areBothSpeciesStillAlive = bothSpeciesStillAlive(entities)
							if !areBothSpeciesStillAlive {
								break
							}
						}
					}
				}
			}
			if !areBothSpeciesStillAlive {
				break
			}
		}
		if print {
			fmt.Print(getPrintableString(grid, entities))
		}
		areBothSpeciesStillAlive = bothSpeciesStillAlive(entities)
		round++
	}
	round--
	if !wasCompleteRound(entities) {
		round--
	}
	hitpoints = countHitpoints(entities)
	return
}

func wasCompleteRound(entities entityLocations) bool {
	round := -1
	for _, row := range entities {
		for _, entity := range row {
			if round != -1 {
				if entity.tickCount != round {
					return false
				}
			} else {
				round = entity.tickCount
			}
		}
	}
	return true
}

func bothSpeciesStillAlive(entities entityLocations) bool {
	elves, goblins := 0, 0
	for _, row := range entities {
		for _, e := range row {
			if e.species == goblin {
				goblins++
			} else if e.species == elf {
				elves++
			}
			if elves > 0 && goblins > 0 {
				return true
			}
		}
	}
	return false
}

func takeTurn(grid [][]bool, entities entityLocations, entity *entity, x, y, tickCount int) (wasSomeoneKilled bool) {
	if entity.tickCount < tickCount {
		entity.tickCount = tickCount
		foe, enemyX, enemyY := getAttackableEnemies(entities, entity.enemy, x, y)
		if foe != nil {
			foe.hitPoints -= entity.attackPower
			if foe.hitPoints <= 0 {
				entities.removeAt(enemyX, enemyY)
				wasSomeoneKilled = true
			}
		} else {
			nextMove := getNextMove(grid, entities, entity, x, y)
			if nextMove != nil {
				entities.moveTo(x, y, nextMove.x, nextMove.y)
				foe, enemyX, enemyY := getAttackableEnemies(entities, entity.enemy, nextMove.x, nextMove.y)
				if foe != nil {
					foe.hitPoints -= entity.attackPower
					if foe.hitPoints <= 0 {
						entities.removeAt(enemyX, enemyY)
						wasSomeoneKilled = true
					}
				}
			}
		}
	}
	return
}

func countHitpoints(entities entityLocations) (sum int) {
	for _, row := range entities {
		for _, entity := range row {
			sum += entity.hitPoints
		}
	}
	return
}
