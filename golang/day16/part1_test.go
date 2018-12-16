package main

import (
	"testing"
)

func TestAddr1(t *testing.T) {
	before := [4]int{3, 2, 1, 1}
	register := [4]int{0, 1, 2, 3}
	expected := [4]int{3, 2, 1, 3}
	actual := addr(before, register)

	if actual != expected {
		t.Errorf("Expected %v, %v to be %v but instead got %v", before, register, expected, actual)
	}
}

func TestAddr2(t *testing.T) {
	before := [4]int{5, 1, 1, 1}
	register := [4]int{0, 0, 1, 1}
	expected := [4]int{5, 6, 1, 1}
	actual := addr(before, register)

	if actual != expected {
		t.Errorf("Expected %v, %v to be %v but instead got %v", before, register, expected, actual)
	}
}

func TestAddi1(t *testing.T) {
	before := [4]int{3, 2, 1, 1}
	register := [4]int{6, 2, 1, 2}
	expected := [4]int{3, 2, 2, 1}
	actual := addi(before, register)

	if actual != expected {
		t.Errorf("Expected %v, %v to be %v but instead got %v", before, register, expected, actual)
	}
}

func TestAddi2(t *testing.T) {
	before := [4]int{5, 1, 1, 1}
	register := [4]int{0, 0, 9, 2}
	expected := [4]int{5, 1, 14, 1}
	actual := addi(before, register)

	if actual != expected {
		t.Errorf("Expected %v, %v to be %v but instead got %v", before, register, expected, actual)
	}
}

func TestMulr1(t *testing.T) {
	before := [4]int{3, 2, 1, 1}
	register := [4]int{9, 2, 1, 2}
	expected := [4]int{3, 2, 2, 1}
	actual := mulr(before, register)

	if actual != expected {
		t.Errorf("Expected %v, %v to be %v but instead got %v", before, register, expected, actual)
	}
}

func TestMulr2(t *testing.T) {
	before := [4]int{5, 1, 1, 1}
	register := [4]int{8, 0, 0, 2}
	expected := [4]int{5, 1, 25, 1}
	actual := mulr(before, register)

	if actual != expected {
		t.Errorf("Expected %v, %v to be %v but instead got %v", before, register, expected, actual)
	}
}

func TestMuli1(t *testing.T) {
	before := [4]int{3, 2, 1, 1}
	register := [4]int{3, 1, 0, 3}
	expected := [4]int{3, 2, 1, 0}
	actual := muli(before, register)

	if actual != expected {
		t.Errorf("Expected %v, %v to be %v but instead got %v", before, register, expected, actual)
	}
}

func TestMuli2(t *testing.T) {
	before := [4]int{5, 1, 1, 1}
	register := [4]int{8, 0, 2, 0}
	expected := [4]int{10, 1, 1, 1}
	actual := muli(before, register)

	if actual != expected {
		t.Errorf("Expected %v, %v to be %v but instead got %v", before, register, expected, actual)
	}
}

func TestBanr1(t *testing.T) {
	before := [4]int{0, 1, 1, 1}
	register := [4]int{6, 2, 1, 0}
	expected := [4]int{1, 1, 1, 1}
	actual := banr(before, register)

	if actual != expected {
		t.Errorf("Expected %v, %v to be %v but instead got %v", before, register, expected, actual)
	}
}

func TestBanr2(t *testing.T) {
	before := [4]int{0, 3, 1, 1}
	register := [4]int{7, 0, 1, 2}
	expected := [4]int{0, 3, 0, 1}
	actual := banr(before, register)

	if actual != expected {
		t.Errorf("Expected %v, %v to be %v but instead got %v", before, register, expected, actual)
	}
}

func TestBani1(t *testing.T) {
	before := [4]int{0, 0, 1, 1}
	register := [4]int{3, 2, 4, 0}
	expected := [4]int{0, 0, 1, 1}
	actual := bani(before, register)

	if actual != expected {
		t.Errorf("Expected %v, %v to be %v but instead got %v", before, register, expected, actual)
	}
}

func TestBani2(t *testing.T) {
	before := [4]int{1, 1, 1, 1}
	register := [4]int{8, 0, 5, 0}
	expected := [4]int{1, 1, 1, 1}
	actual := bani(before, register)

	if actual != expected {
		t.Errorf("Expected %v, %v to be %v but instead got %v", before, register, expected, actual)
	}
}

func TestBorr1(t *testing.T) {
	before := [4]int{5, 1, 1, 1}
	register := [4]int{6, 2, 0, 0}
	expected := [4]int{5, 1, 1, 1}
	actual := borr(before, register)

	if actual != expected {
		t.Errorf("Expected %v, %v to be %v but instead got %v", before, register, expected, actual)
	}
}

func TestBorr2(t *testing.T) {
	before := [4]int{6, 1, 0, 1}
	register := [4]int{7, 0, 2, 1}
	expected := [4]int{6, 6, 0, 1}
	actual := borr(before, register)

	if actual != expected {
		t.Errorf("Expected %v, %v to be %v but instead got %v", before, register, expected, actual)
	}
}

func TestBori1(t *testing.T) {
	before := [4]int{0, 0, 1, 1}
	register := [4]int{3, 2, 3, 0}
	expected := [4]int{3, 0, 1, 1}
	actual := bori(before, register)

	if actual != expected {
		t.Errorf("Expected %v, %v to be %v but instead got %v", before, register, expected, actual)
	}
}

func TestBori2(t *testing.T) {
	before := [4]int{2, 1, 3, 3}
	register := [4]int{15, 1, 3, 0}
	expected := [4]int{3, 1, 3, 3}
	actual := bori(before, register)

	if actual != expected {
		t.Errorf("Expected %v, %v to be %v but instead got %v", before, register, expected, actual)
	}
}

func TestSetr1(t *testing.T) {
	before := [4]int{3, 2, 1, 1}
	register := [4]int{0, 1, 2, 3}
	expected := [4]int{3, 2, 1, 2}
	actual := setr(before, register)

	if actual != expected {
		t.Errorf("Expected %v, %v to be %v but instead got %v", before, register, expected, actual)
	}
}

func TestSetr2(t *testing.T) {
	before := [4]int{5, 1, 1, 1}
	register := [4]int{0, 0, 1, 1}
	expected := [4]int{5, 5, 1, 1}
	actual := setr(before, register)

	if actual != expected {
		t.Errorf("Expected %v, %v to be %v but instead got %v", before, register, expected, actual)
	}
}

func TestSeti1(t *testing.T) {
	before := [4]int{3, 2, 1, 1}
	register := [4]int{6, 4, 1, 2}
	expected := [4]int{3, 2, 4, 1}
	actual := seti(before, register)

	if actual != expected {
		t.Errorf("Expected %v, %v to be %v but instead got %v", before, register, expected, actual)
	}
}

func TestSeti2(t *testing.T) {
	before := [4]int{5, 1, 1, 1}
	register := [4]int{0, 0, 9, 2}
	expected := [4]int{5, 1, 0, 1}
	actual := seti(before, register)

	if actual != expected {
		t.Errorf("Expected %v, %v to be %v but instead got %v", before, register, expected, actual)
	}
}

func TestGtir1(t *testing.T) {
	before := [4]int{3, 2, 5, 1}
	register := [4]int{6, 4, 1, 2}
	expected := [4]int{3, 2, 1, 1}
	actual := gtir(before, register)

	if actual != expected {
		t.Errorf("Expected %v, %v to be %v but instead got %v", before, register, expected, actual)
	}
}

func TestGtir2(t *testing.T) {
	before := [4]int{5, 1, 1, 3}
	register := [4]int{0, 0, 3, 0}
	expected := [4]int{0, 1, 1, 3}
	actual := gtir(before, register)

	if actual != expected {
		t.Errorf("Expected %v, %v to be %v but instead got %v", before, register, expected, actual)
	}
}

func TestGtri1(t *testing.T) {
	before := [4]int{3, 2, 5, 1}
	register := [4]int{6, 1, 1, 2}
	expected := [4]int{3, 2, 1, 1}
	actual := gtri(before, register)

	if actual != expected {
		t.Errorf("Expected %v, %v to be %v but instead got %v", before, register, expected, actual)
	}
}

func TestGtri2(t *testing.T) {
	before := [4]int{5, 1, 1, 3}
	register := [4]int{0, 1, 3, 0}
	expected := [4]int{0, 1, 1, 3}
	actual := gtri(before, register)

	if actual != expected {
		t.Errorf("Expected %v, %v to be %v but instead got %v", before, register, expected, actual)
	}
}

func TestGtrr1(t *testing.T) {
	before := [4]int{3, 2, 5, 1}
	register := [4]int{6, 1, 2, 2}
	expected := [4]int{3, 2, 0, 1}
	actual := gtrr(before, register)

	if actual != expected {
		t.Errorf("Expected %v, %v to be %v but instead got %v", before, register, expected, actual)
	}
}

func TestGtrr2(t *testing.T) {
	before := [4]int{5, 1, 1, 3}
	register := [4]int{0, 0, 3, 0}
	expected := [4]int{1, 1, 1, 3}
	actual := gtrr(before, register)

	if actual != expected {
		t.Errorf("Expected %v, %v to be %v but instead got %v", before, register, expected, actual)
	}
}

func TestEqir1(t *testing.T) {
	before := [4]int{3, 2, 5, 1}
	register := [4]int{6, 4, 1, 2}
	expected := [4]int{3, 2, 0, 1}
	actual := eqir(before, register)

	if actual != expected {
		t.Errorf("Expected %v, %v to be %v but instead got %v", before, register, expected, actual)
	}
}

func TestEqir2(t *testing.T) {
	before := [4]int{5, 1, 1, 4}
	register := [4]int{0, 4, 3, 0}
	expected := [4]int{1, 1, 1, 4}
	actual := eqir(before, register)

	if actual != expected {
		t.Errorf("Expected %v, %v to be %v but instead got %v", before, register, expected, actual)
	}
}

func TestEqri1(t *testing.T) {
	before := [4]int{3, 2, 5, 1}
	register := [4]int{6, 1, 2, 2}
	expected := [4]int{3, 2, 1, 1}
	actual := eqri(before, register)

	if actual != expected {
		t.Errorf("Expected %v, %v to be %v but instead got %v", before, register, expected, actual)
	}
}

func TestEqri2(t *testing.T) {
	before := [4]int{5, 1, 1, 3}
	register := [4]int{0, 1, 3, 0}
	expected := [4]int{0, 1, 1, 3}
	actual := eqri(before, register)

	if actual != expected {
		t.Errorf("Expected %v, %v to be %v but instead got %v", before, register, expected, actual)
	}
}

func TestEqrr1(t *testing.T) {
	before := [4]int{3, 2, 5, 1}
	register := [4]int{6, 1, 2, 2}
	expected := [4]int{3, 2, 0, 1}
	actual := eqrr(before, register)

	if actual != expected {
		t.Errorf("Expected %v, %v to be %v but instead got %v", before, register, expected, actual)
	}
}

func TestEqrr2(t *testing.T) {
	before := [4]int{5, 1, 1, 0}
	register := [4]int{0, 0, 3, 0}
	expected := [4]int{0, 1, 1, 0}
	actual := eqrr(before, register)

	if actual != expected {
		t.Errorf("Expected %v, %v to be %v but instead got %v", before, register, expected, actual)
	}
}
