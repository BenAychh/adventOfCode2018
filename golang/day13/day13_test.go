package main

import (
	"testing"

	"github.com/benaychh/aocutils"
)

func TestFirstCrash(t *testing.T) {
	fileName := "test.data"
	rawData := aocutils.LoadFileAsStringArray(fileName)
	tracks, carts := makeTracks(rawData)
	expected := aocutils.Vertex{X: 7, Y: 3}
	actual := *getFirstCrash(tracks, carts, true)

	if actual != expected {
		t.Errorf("Expected %v to be %+v but instead got %+v", fileName, expected, actual)
	}
}

func TestLastCart(t *testing.T) {
	fileName := "test2.data"
	rawData := aocutils.LoadFileAsStringArray(fileName)
	tracks, carts := makeTracks(rawData)
	expected := aocutils.Vertex{X: 6, Y: 4}
	actual := *getLastCart(tracks, carts, true)

	if actual != expected {
		t.Errorf("Expected %v to be %+v but instead got %+v", fileName, expected, actual)
	}
}
