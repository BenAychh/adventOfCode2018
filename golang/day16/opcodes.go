package main

func addr(before [4]int, register [4]int) (after [4]int) {
	if !validRegistryAddresses(register[1], register[2], register[3]) {
		return [4]int{-1, -1, -1, -1}
	}
	after = before
	a := before[register[1]]
	b := before[register[2]]
	after[register[3]] = a + b
	return
}

func addi(before [4]int, register [4]int) (after [4]int) {
	if !validRegistryAddresses(register[1], register[3]) {
		return [4]int{-1, -1, -1, -1}
	}
	after = before
	a := before[register[1]]
	b := register[2]
	after[register[3]] = a + b
	return
}

func mulr(before [4]int, register [4]int) (after [4]int) {
	if !validRegistryAddresses(register[1], register[2], register[3]) {
		return [4]int{-1, -1, -1, -1}
	}
	after = before
	a := before[register[1]]
	b := before[register[2]]
	after[register[3]] = a * b
	return
}

func muli(before [4]int, register [4]int) (after [4]int) {
	if !validRegistryAddresses(register[1], register[3]) {
		return [4]int{-1, -1, -1, -1}
	}
	after = before
	a := before[register[1]]
	b := register[2]
	after[register[3]] = a * b
	return
}

func banr(before [4]int, register [4]int) (after [4]int) {
	if !validRegistryAddresses(register[1], register[2], register[3]) {
		return [4]int{-1, -1, -1, -1}
	}
	a := before[register[1]]
	b := before[register[2]]
	after = before
	after[register[3]] = a & b
	return
}

func bani(before [4]int, register [4]int) (after [4]int) {
	if !validRegistryAddresses(register[1], register[3]) {
		return [4]int{-1, -1, -1, -1}
	}
	a := before[register[1]]
	b := register[2]
	after = before
	after[register[3]] = a & b
	return
}

func borr(before [4]int, register [4]int) (after [4]int) {
	if !validRegistryAddresses(register[1], register[2], register[3]) {
		return [4]int{-1, -1, -1, -1}
	}
	a := before[register[1]]
	b := before[register[2]]
	after = before
	after[register[3]] = a | b
	return
}

func bori(before [4]int, register [4]int) (after [4]int) {
	if !validRegistryAddresses(register[1], register[3]) {
		return [4]int{-1, -1, -1, -1}
	}
	a := before[register[1]]
	b := register[2]
	after = before
	after[register[3]] = a | b
	return
}

func setr(before [4]int, register [4]int) (after [4]int) {
	if !validRegistryAddresses(register[1], register[3]) {
		return [4]int{-1, -1, -1, -1}
	}
	a := before[register[1]]
	after = before
	after[register[3]] = a
	return
}

func seti(before [4]int, register [4]int) (after [4]int) {
	if !validRegistryAddresses(register[3]) {
		return [4]int{-1, -1, -1, -1}
	}
	a := register[1]
	after = before
	after[register[3]] = a
	return
}

func gtir(before [4]int, register [4]int) (after [4]int) {
	if !validRegistryAddresses(register[2], register[3]) {
		return [4]int{-1, -1, -1, -1}
	}
	a := register[1]
	b := before[register[2]]
	after = before
	after[register[3]] = 0
	if a > b {
		after[register[3]] = 1
	}
	return
}

func gtri(before [4]int, register [4]int) (after [4]int) {
	if !validRegistryAddresses(register[1], register[3]) {
		return [4]int{-1, -1, -1, -1}
	}
	a := before[register[1]]
	b := register[2]
	after = before
	after[register[3]] = 0
	if a > b {
		after[register[3]] = 1
	}
	return
}

func gtrr(before [4]int, register [4]int) (after [4]int) {
	if !validRegistryAddresses(register[1], register[2], register[3]) {
		return [4]int{-1, -1, -1, -1}
	}
	a := before[register[1]]
	b := before[register[2]]
	after = before
	after[register[3]] = 0
	if a > b {
		after[register[3]] = 1
	}
	return
}

func eqir(before [4]int, register [4]int) (after [4]int) {
	if !validRegistryAddresses(register[2], register[3]) {
		return [4]int{-1, -1, -1, -1}
	}
	a := register[1]
	b := before[register[2]]
	after = before
	after[register[3]] = 0
	if a == b {
		after[register[3]] = 1
	}
	return
}

func eqri(before [4]int, register [4]int) (after [4]int) {
	if !validRegistryAddresses(register[1], register[3]) {
		return [4]int{-1, -1, -1, -1}
	}
	a := before[register[1]]
	b := register[2]
	after = before
	after[register[3]] = 0
	if a == b {
		after[register[3]] = 1
	}
	return
}

func eqrr(before [4]int, register [4]int) (after [4]int) {
	if !validRegistryAddresses(register[1], register[2], register[3]) {
		return [4]int{-1, -1, -1, -1}
	}
	a := before[register[1]]
	b := before[register[2]]
	after = before
	after[register[3]] = 0
	if a == b {
		after[register[3]] = 1
	}
	return
}

func validRegistryAddresses(addresses ...int) bool {
	for _, addr := range addresses {
		if addr > 3 {
			return false
		}
	}
	return true
}

var opcodes = [16]func([4]int, [4]int) [4]int{
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
