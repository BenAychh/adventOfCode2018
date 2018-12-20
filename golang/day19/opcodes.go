package main

import "log"

func addr(before [6]int, register [3]int) (after [6]int) {
	if !validRegistryAddresses(register[0], register[1], register[2]) {
		log.Fatal("error", before, register)
	}
	after = before
	a := before[register[0]]
	b := before[register[1]]
	after[register[2]] = a + b
	return
}

func addi(before [6]int, register [3]int) (after [6]int) {
	if !validRegistryAddresses(register[0], register[2]) {
		log.Fatal("error", before, register)
	}
	after = before
	a := before[register[0]]
	b := register[1]
	after[register[2]] = a + b
	return
}

func mulr(before [6]int, register [3]int) (after [6]int) {
	if !validRegistryAddresses(register[0], register[1], register[2]) {
		log.Fatal("error", before, register)
	}
	after = before
	a := before[register[0]]
	b := before[register[1]]
	after[register[2]] = a * b
	return
}

func muli(before [6]int, register [3]int) (after [6]int) {
	if !validRegistryAddresses(register[0], register[2]) {
		log.Fatal("error", before, register)
	}
	after = before
	a := before[register[0]]
	b := register[1]
	after[register[2]] = a * b
	return
}

func banr(before [6]int, register [3]int) (after [6]int) {
	if !validRegistryAddresses(register[0], register[1], register[2]) {
		log.Fatal("error", before, register)
	}
	a := before[register[0]]
	b := before[register[1]]
	after = before
	after[register[2]] = a & b
	return
}

func bani(before [6]int, register [3]int) (after [6]int) {
	if !validRegistryAddresses(register[0], register[2]) {
		log.Fatal("error", before, register)
	}
	a := before[register[0]]
	b := register[1]
	after = before
	after[register[2]] = a & b
	return
}

func borr(before [6]int, register [3]int) (after [6]int) {
	if !validRegistryAddresses(register[0], register[1], register[2]) {
		log.Fatal("error", before, register)
	}
	a := before[register[0]]
	b := before[register[1]]
	after = before
	after[register[2]] = a | b
	return
}

func bori(before [6]int, register [3]int) (after [6]int) {
	if !validRegistryAddresses(register[0], register[2]) {
		log.Fatal("error", before, register)
	}
	a := before[register[0]]
	b := register[1]
	after = before
	after[register[2]] = a | b
	return
}

func setr(before [6]int, register [3]int) (after [6]int) {
	if !validRegistryAddresses(register[0], register[2]) {
		log.Fatal("error", before, register)
	}
	a := before[register[0]]
	after = before
	after[register[2]] = a
	return
}

func seti(before [6]int, register [3]int) (after [6]int) {
	if !validRegistryAddresses(register[2]) {
		log.Fatal("error", before, register)
	}
	a := register[0]
	after = before
	after[register[2]] = a
	return
}

func gtir(before [6]int, register [3]int) (after [6]int) {
	if !validRegistryAddresses(register[1], register[2]) {
		log.Fatal("error", before, register)
	}
	a := register[0]
	b := before[register[1]]
	after = before
	after[register[2]] = 0
	if a > b {
		after[register[2]] = 1
	}
	return
}

func gtri(before [6]int, register [3]int) (after [6]int) {
	if !validRegistryAddresses(register[0], register[2]) {
		log.Fatal("error", before, register)
	}
	a := before[register[0]]
	b := register[1]
	after = before
	after[register[2]] = 0
	if a > b {
		after[register[2]] = 1
	}
	return
}

func gtrr(before [6]int, register [3]int) (after [6]int) {
	if !validRegistryAddresses(register[0], register[1], register[2]) {
		log.Fatal("error", before, register)
	}
	a := before[register[0]]
	b := before[register[1]]
	after = before
	after[register[2]] = 0
	if a > b {
		after[register[2]] = 1
	}
	return
}

func eqir(before [6]int, register [3]int) (after [6]int) {
	if !validRegistryAddresses(register[1], register[2]) {
		log.Fatal("error", before, register)
	}
	a := register[0]
	b := before[register[1]]
	after = before
	after[register[2]] = 0
	if a == b {
		after[register[2]] = 1
	}
	return
}

func eqri(before [6]int, register [3]int) (after [6]int) {
	if !validRegistryAddresses(register[0], register[2]) {
		log.Fatal("error", before, register)
	}
	a := before[register[0]]
	b := register[1]
	after = before
	after[register[2]] = 0
	if a == b {
		after[register[2]] = 1
	}
	return
}

func eqrr(before [6]int, register [3]int) (after [6]int) {
	if !validRegistryAddresses(register[0], register[1], register[2]) {
		log.Fatal("error", before, register)
	}
	a := before[register[0]]
	b := before[register[1]]
	after = before
	after[register[2]] = 0
	if a == b {
		after[register[2]] = 1
	}
	return
}

func validRegistryAddresses(addresses ...int) bool {
	for _, addr := range addresses {
		if addr > 5 {
			return false
		}
	}
	return true
}

var opcodes = [16]func([6]int, [3]int) [6]int{
	addr,
	addi,
	mulr,
	muli,
	banr,
	bani,
	borr,
	bori,
	setr,
	seti,
	gtir,
	gtri,
	gtrr,
	eqir,
	eqri,
	eqrr,
}

var opcodesMap = map[string]func([6]int, [3]int) [6]int{
	"addr": addr,
	"addi": addi,
	"mulr": mulr,
	"muli": muli,
	"banr": banr,
	"bani": bani,
	"borr": borr,
	"bori": bori,
	"setr": setr,
	"seti": seti,
	"gtir": gtir,
	"gtri": gtri,
	"gtrr": gtrr,
	"eqir": eqir,
	"eqri": eqri,
	"eqrr": eqrr,
}
