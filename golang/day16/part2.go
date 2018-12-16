package main

import "regexp"

func processTestProgram(interpretations []interpretation, data []string) (result [4]int) {
	re := regexp.MustCompile("-?[0-9]+")
	opMap := getOpMap(interpretations)
	registries := [][4]int{}
	for _, datum := range data {
		registries = append(registries, processLine(datum, re))
	}
	for _, registry := range registries {
		result = opMap[registry[0]](result, registry)
	}
	return
}

func getOpMap(interpretations []interpretation) (opMap map[int]func([4]int, [4]int) [4]int) {
	possibilities := getPossibilities(interpretations)
	opNameMap := getNameMap(possibilities)
	opMap = map[int]func([4]int, [4]int) [4]int{}

	for i, name := range opNameMap {
		opMap[i] = opcodesMap[name]
	}

	return
}

func getPossibilities(interpretations []interpretation) (possibilities map[int]map[string]bool) {
	possibilities = map[int]map[string]bool{}
	for _, interp := range interpretations {
		for key, fn := range opcodesMap {
			if fn(interp.before, interp.register) == interp.after {
				code := interp.register[0]
				possibility, ok := possibilities[code]
				if !ok {
					possibilities[code] = map[string]bool{
						key: true,
					}
				} else {
					_, alreadyIncluded := possibility[key]
					if !alreadyIncluded {
						possibility[key] = true
					}
				}
			}
		}
	}
	return
}

func getNameMap(possibilities map[int]map[string]bool) (opNameMap map[int]string) {
	opNameMap = map[int]string{}
	for len(opNameMap) != 16 {
		for i, possibility := range possibilities {
			for _, name := range opNameMap {
				delete(possibility, name)
			}
			if len(possibility) == 1 {
				for key := range possibility {
					opNameMap[i] = key
				}
			}
		}
	}
	return
}
