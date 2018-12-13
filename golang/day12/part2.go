package main

func getGenerationOptimized(currentState string, startStart, currentGenerationNumber int, notes map[string]rune, targetGeneration int) (sum int) {
	currentGeneration := map[int]rune{}
	for i, r := range currentState {
		currentGeneration[i+startStart] = r
	}
	for i := currentGenerationNumber + 1; i < targetGeneration; i++ {
		currentGeneration, _ = getNextGeneration(currentGeneration, notes)
	}

	for key, value := range currentGeneration {
		if value == '#' {
			sum += key
		}
	}
	return
}
