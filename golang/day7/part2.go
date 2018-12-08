package main

import (
	"sort"
)

func getTotalTime(data []string, workers, baseTimePerTask int) int {
	parsedSteps := parseSteps(data)
	extraInfoSteps := createExtraInfoStepStatuses(parsedSteps, baseTimePerTask)
	workerQueue := createWorkers(workers)
	time := 0
	for !areAllDone(extraInfoSteps) {
		ready := getAllReadyAndSortedWithTime(extraInfoSteps)
		for _, readyStepName := range ready {
			for _, worker := range workerQueue {
				if worker.assignment == nil {
					extraInfoSteps[readyStepName].worker = worker
					extraInfoSteps[readyStepName].startedAt = time
					break
				}
			}
		}
		time++
		doTick(extraInfoSteps, time)
	}
	return time
}

type worker struct {
	assignment *extraStepStatusInfo
}

type extraStepStatusInfo struct {
	*stepStatus
	startedAt int
	totalTime int
	*worker
}

func doTick(extraInfoMap map[string]*extraStepStatusInfo, currentTime int) {
	for _, info := range extraInfoMap {
		info.tick(currentTime)
	}
}

func (e *extraStepStatusInfo) tick(currentTime int) {
	if e.startedAt > -1 && !e.isDone {
		if currentTime >= e.startedAt+e.totalTime {
			e.worker.assignment = nil
			e.worker = nil
			e.isDone = true
		}
	}
}

func createExtraInfoStepStatuses(parsedSteps map[string]*stepStatus, baseTimePerTask int) (extraInfoMap map[string]*extraStepStatusInfo) {
	extraInfoMap = map[string]*extraStepStatusInfo{}
	for key, value := range parsedSteps {
		extraInfoMap[key] = &extraStepStatusInfo{
			stepStatus: value,
			startedAt:  -1,
			worker:     nil,
			totalTime:  getSecondsFromLetter(key) + baseTimePerTask,
		}
	}
	return
}

func areAllDone(steps map[string]*extraStepStatusInfo) bool {
	for _, step := range steps {
		if !step.isDone {
			return false
		}
	}
	return true
}

func getAllReadyAndSortedWithTime(steps map[string]*extraStepStatusInfo) (areReady []string) {
	areReady = []string{}
	for name, status := range steps {
		if !status.isDone && status.startedAt == -1 && isReadyWithTime(name, steps) {
			areReady = append(areReady, name)
		}
	}
	sort.Strings(areReady)
	return areReady
}

func isReadyWithTime(name string, steps map[string]*extraStepStatusInfo) bool {
	ready := true
	for _, requirement := range steps[name].requirements {
		if !steps[requirement].isDone {
			ready = false
			break
		}
	}
	return ready
}

func createWorkers(count int) (workers []*worker) {
	workers = make([]*worker, count)
	for i := 0; i < count; i++ {
		workers[i] = &worker{}
	}
	return workers
}

func getSecondsFromLetter(letter string) int {
	return int([]rune(letter)[0]) - 64
}
