package main

import "fmt"

func followInstructionsUntilDone(instructions instructionLines, currentRegisterState [6]int, ipRegister int) [6]int {
	ip := 0
	i := 0
	for instructions.hasInstruction(ip) {
		i++
		currentRegisterState[ipRegister] = ip
		instruction := instructions[ip]
		fn := opcodesMap[instruction.fnName]
		currentRegisterState = fn(currentRegisterState, instruction.set)
		ip = currentRegisterState[ipRegister] + 1
		fmt.Println(i, ip, instruction, currentRegisterState)
	}
	return currentRegisterState
}
