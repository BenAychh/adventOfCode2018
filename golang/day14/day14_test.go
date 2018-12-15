package main

import (
	"testing"
)

func Test10RecipesAfter9(t *testing.T) {
	data := 9
	expected := "5158916779"
	actual := getNext10(data, false)

	if actual != expected {
		t.Errorf("Expected %d to be %v but instead got %v", data, expected, actual)
	}
}

func Test10RecipesAfter5(t *testing.T) {
	data := 5
	expected := "0124515891"
	actual := getNext10(data, false)

	if actual != expected {
		t.Errorf("Expected %d to be %v but instead got %v", data, expected, actual)
	}
}

func Test10RecipesAfter18(t *testing.T) {
	data := 18
	expected := "9251071085"
	actual := getNext10(data, false)

	if actual != expected {
		t.Errorf("Expected %d to be %v but instead got %v", data, expected, actual)
	}
}

func Test10RecipesAfter2018(t *testing.T) {
	data := 2018
	expected := "5941429882"
	actual := getNext10(data, false)

	if actual != expected {
		t.Errorf("Expected %d to be %v but instead got %v", data, expected, actual)
	}
}

func Test5RecipesBefore9(t *testing.T) {
	data := []int{5, 1, 5, 8, 9}
	expected := 9
	actual := getPrevious5(data, false)

	if actual != expected {
		t.Errorf("Expected %d to be %v but instead got %v", data, expected, actual)
	}
}

func Test5RecipesBefore5(t *testing.T) {
	data := []int{0, 1, 2, 4, 5}
	expected := 5
	actual := getPrevious5(data, false)

	if actual != expected {
		t.Errorf("Expected %d to be %v but instead got %v", data, expected, actual)
	}
}

func Test5RecipesBefore18(t *testing.T) {
	data := []int{9, 2, 5, 1, 0}
	expected := 18
	actual := getPrevious5(data, false)

	if actual != expected {
		t.Errorf("Expected %d to be %v but instead got %v", data, expected, actual)
	}
}

func Test5RecipesBefore2018(t *testing.T) {
	data := []int{5, 9, 4, 1, 4}
	expected := 2018
	actual := getPrevious5(data, false)

	if actual != expected {
		t.Errorf("Expected %d to be %v but instead got %v", data, expected, actual)
	}
}
