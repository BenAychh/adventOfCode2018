package main

// 01 seti 1 0 4	r4 = 1
// 					LAND
// 02 seti 1 0 1	r1 = 1

// 03 mulr 4 1 3	r3 = r4 * r1
// 04 eqrr 3 5 3	r3 = (r4 * r1) == r5
// 05 addr 3 2 2	r2 = ((r4 * r1) == r5) + r2		IF-1
// 06 addi 2 1 2	r2 = r2 + 1										ELSE-1
//				 	IF-1
// 07 addr 4 0 0	r0 = r4 + r0
//					ELSE-1
// 08 addi 1 1 1	r1 = r1 + 1

// 09 gtrr 1 5 3	r3 = r1 > r5
// 10 addr 2 3 2	r2 = r2 + (r1 > r5)						IF-2
// 11 seti 2 5 2	r2 = 2												JUMP
//					IF-2
// 12 addi 4 1 4	r4 = r4 + 1
// 13 gtrr 4 5 3	r3 = r4 > r5
// 14 addr 3 2 2	r2 = r2 + r4 > r5
// 15 seti 1 9 2	r2 = 1												JUMP to 2
// 16 mulr 2 2 2				DEST

func getOptimizedBreakIpRegister2(zerothInt int) int {
	r0 := 0
	// r1 := 0
	r4 := 0
	r5 := 0

	if zerothInt == 1 {
		r4, r5 = 1, 10551417
	} else {
		r4, r5 = 1, 1017
	}

	for r4 <= r5 {
		// for r1 <= r5 {
		// 	if r4*r1 == r5 {
		// 		r0 = r4 + r0
		// 	}
		// 	r1 = r1 + 1
		// }
		if r5%r4 == 0 {
			r0 = r0 + r4
		}
		r4 = r4 + 1
	}
	return r0
}
