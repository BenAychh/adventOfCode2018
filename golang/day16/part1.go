package main

func findAllInterpretationsWithOver2Matches(interpretations []interpretation) (count int) {
	for _, interp := range interpretations {
		internalCount := 0
		for _, fn := range opcodes {
			if fn(interp.before, interp.register) == interp.after {
				internalCount++
				if internalCount > 2 {
					break
				}
			}
		}
		if internalCount > 2 {
			count++
		}
	}
	return
}
