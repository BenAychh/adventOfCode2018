package main

import (
	"testing"

	"github.com/benaychh/aocutils"
)

func TestFinalRegisterState(t *testing.T) {
	fileName := "day19.test.data"
	rawData := aocutils.LoadFileAsStringArray(fileName)
	lines, ipRegister := getInstructionLines(rawData)
	expected := [6]int{6, 5, 6, 0, 0, 9}
	actual := followInstructionsUntilDone(lines, [6]int{0, 0, 0, 0, 0, 0}, ipRegister)

	if actual != expected {
		t.Errorf("Expected %v to be %v but instead got %v", fileName, expected, actual)
	}
}

func TestOptimizedFn(t *testing.T) {
	expected := 1482
	actual := getOptimizedBreakIpRegister2(0)

	if actual != expected {
		t.Errorf("Expected optimized (0) to be %v but instead got %v", expected, actual)
	}
}
