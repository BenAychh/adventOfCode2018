package main

import (
	"log"
	"sort"
)

type magicalCreature int64

const (
	elf    magicalCreature = iota
	goblin magicalCreature = iota
)

type entity struct {
	attackPower int
	hitPoints   int
	species     magicalCreature
	enemy       magicalCreature
	tickCount   int
}

type entityLocations map[int]entityRow
type entityRow map[int]*entity

type entityMap interface {
	atLocation(int, int) bool
	speciesAtLocation(int, int, magicalCreature) bool
	getAt(int, int) *entity
	removeAt(int, int)
	placeAt(int, int, *entity)
	moveTo(int, int, int, int)
	sortedKeys() []int
}

type entityLine interface {
	sortedKeys() []int
}

func (el entityLocations) atLocation(x, y int) bool {
	row, ok := el[y]
	if !ok {
		return false
	}
	_, ok = row[x]
	return ok
}

func (el entityLocations) speciesAtLocation(x, y int, species magicalCreature) bool {
	row, ok := el[y]
	if !ok {
		return false
	}
	e, ok := row[x]
	if !ok {
		return false
	}
	return e.species == species
}

func (el entityLocations) getAt(x, y int) *entity {
	row, ok := el[y]
	if !ok {
		log.Fatalf("No entity at %d, %d", x, y)
	}
	cell, ok := row[x]
	if !ok {
		log.Fatalf("No entity at %d, %d", x, y)
	}
	return cell
}

func (el entityLocations) removeAt(x, y int) {
	row, ok := el[y]
	if !ok {
		log.Fatalf("No entity at %d, %d", x, y)
	}
	_, ok = row[x]
	if !ok {
		log.Fatalf("No entity at %d, %d", x, y)
	}
	if len(row) == 1 {
		delete(el, y)
	} else {
		delete(row, x)
	}
}

func (el entityLocations) placeAt(x, y int, e *entity) {
	row, ok := el[y]
	if !ok {
		el[y] = entityRow{
			x: e,
		}
	} else {
		_, ok = row[x]
		if ok {
			log.Fatalf("entity already exists at %d, %d", x, y)
		}
		row[x] = e
	}
}

func (el entityLocations) moveTo(fromX, fromY, toX, toY int) {
	entity := el.getAt(fromX, fromY)
	el.placeAt(toX, toY, entity)
	el.removeAt(fromX, fromY)
}

func (el entityLocations) sortedKeys() []int {
	keys := []int{}
	for key := range el {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	return keys
}

func (er entityRow) sortedKeys() []int {
	keys := []int{}
	for key := range er {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	return keys
}
