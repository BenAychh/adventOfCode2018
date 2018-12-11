package main

import (
	"testing"
)

// [{9 1 0 2} {7 0 1 0} {3 2 1 1} {6 10 2 1} {2 4 2 2} {6 10 2 2} {1 8 1 1} {1 7 1 0} {3 11 1 2} {7 6 1 1} {2 3 1 0} {4 3 2 0} {10 3 1 1} {5 11 1 2} {4 7 0 1} {8 2 0 1} {15 0 2 0} {1 6 1 0} {8 9 0 1} {3 3 1 1} {0 5 0 1} {2 2 2 0} {5 2 1 2} {1 4 2 1} {2 7 2 2} {3 6 1 1} {5 0 1 0} {6 0 2 0} {5 9 1 2} {14 7 2 0} {3 6 2 1}]

var lines = []string{
	"position=< 9,  1> velocity=< 0,  2>",
	"position=< 7,  0> velocity=<-1,  0>",
	"position=< 3, -2> velocity=<-1,  1>",
	"position=< 6, 10> velocity=<-2, -1>",
	"position=< 2, -4> velocity=< 2,  2>",
	"position=<-6, 10> velocity=< 2, -2>",
	"position=< 1,  8> velocity=< 1, -1>",
	"position=< 1,  7> velocity=< 1,  0>",
	"position=<-3, 11> velocity=< 1, -2>",
	"position=< 7,  6> velocity=<-1, -1>",
	"position=<-2,  3> velocity=< 1,  0>",
	"position=<-4,  3> velocity=< 2,  0>",
	"position=<10, -3> velocity=<-1,  1>",
	"position=< 5, 11> velocity=< 1, -2>",
	"position=< 4,  7> velocity=< 0, -1>",
	"position=< 8, -2> velocity=< 0,  1>",
	"position=<15,  0> velocity=<-2,  0>",
	"position=< 1,  6> velocity=< 1,  0>",
	"position=< 8,  9> velocity=< 0, -1>",
	"position=< 3,  3> velocity=<-1,  1>",
	"position=< 0,  5> velocity=< 0, -1>",
	"position=<-2,  2> velocity=< 2,  0>",
	"position=< 5, -2> velocity=< 1,  2>",
	"position=< 1,  4> velocity=< 2,  1>",
	"position=<-2,  7> velocity=< 2, -2>",
	"position=< 3,  6> velocity=<-1, -1>",
	"position=< 5,  0> velocity=< 1,  0>",
	"position=<-6,  0> velocity=< 2,  0>",
	"position=< 5,  9> velocity=< 1, -2>",
	"position=<14,  7> velocity=<-2,  0>",
	"position=<-3,  6> velocity=< 2, -1>",
}

func TestMessage(t *testing.T) {
	data := []movingPoint{}
	for _, line := range lines {
		data = append(data, processLine(line))
	}
	expected := 3
	actual := getMessage(data)
	if actual != expected {
		t.Errorf("Expected %v to be %d but instead got %d", data, expected, actual)
	}
}
