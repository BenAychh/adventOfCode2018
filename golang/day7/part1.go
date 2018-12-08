package main

import (
	"sort"
	"strings"
)

func getOrder(data []string) string {
	parsedSteps := parseSteps(data)
	ready := getAllReadyAndSorted(parsedSteps)
	var sb strings.Builder
	for len(ready) != 0 {
		parsedSteps[ready[0]].isDone = true
		sb.WriteString(ready[0])
		ready = getAllReadyAndSorted(parsedSteps)
	}
	return sb.String()
}

func parseSteps(data []string) (steps map[string]*stepStatus) {
	steps = map[string]*stepStatus{}
	for _, line := range data {
		split := strings.Split(line, " ")
		name := split[7]
		requirement := split[1]
		step, nameAlreadyExists := steps[name]
		if nameAlreadyExists {
			step.requirements = append(step.requirements, requirement)
		} else {
			newStep := stepStatus{requirements: []string{requirement}}
			steps[name] = &newStep
		}

		_, requirementAlreadyExists := steps[requirement]
		if !requirementAlreadyExists {
			steps[requirement] = &stepStatus{}
		}
	}
	return
}

func isReady(name string, steps map[string]*stepStatus) bool {
	ready := true
	for _, requirement := range steps[name].requirements {
		if !steps[requirement].isDone {
			ready = false
			break
		}
	}
	return ready
}

func getAllReadyAndSorted(steps map[string]*stepStatus) (areReady []string) {
	areReady = []string{}
	for name, status := range steps {
		if !status.isDone && isReady(name, steps) {
			areReady = append(areReady, name)
		}
	}
	sort.Strings(areReady)
	return areReady
}

type stepStatus struct {
	isDone       bool
	requirements []string
}
