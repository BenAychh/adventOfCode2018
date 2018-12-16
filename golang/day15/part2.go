package main

import (
	"fmt"
)

func findLowestPowerNeededToKeepElvesAlive(grid [][]bool, entities entityLocations, print bool) (round, hitpoints int) {
	var elvesLived bool
	lastFailure := 3
	step := 16
	lastSuccess := -1
	for lastSuccess-lastFailure != 1 {
		attackPower := (lastSuccess-lastFailure)/2 + lastFailure
		if lastSuccess == -1 {
			attackPower = lastFailure + step
		}
		copyOfEntities := copyWithPower(entities, attackPower)
		elvesLived, round, hitpoints = doElvesLive(grid, copyOfEntities, print)
		if elvesLived {
			lastSuccess = attackPower
		} else {
			lastFailure = attackPower
		}
	}
	if print {
		fmt.Printf("Attack Power: %d\n", lastSuccess)
	}
	if !elvesLived {
		copyOfEntities := copyWithPower(entities, lastSuccess)
		_, round, hitpoints = doElvesLive(grid, copyOfEntities, print)
	}
	return
}

func copyWithPower(el entityLocations, attackPower int) (newEl entityLocations) {
	newEl = entityLocations{}
	for y, row := range el {
		newRow := entityRow{}
		for x, e := range row {
			newE := *e
			if newE.species == elf {
				newE.attackPower = attackPower
			}
			newRow[x] = &newE
		}
		newEl[y] = newRow
	}
	return
}

func doElvesLive(grid [][]bool, entities entityLocations, print bool) (elvesLived bool, round, hitpoints int) {
	elvesLived = true
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
							if entity.species == goblin {
								elvesLived = false
								return
							}
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
