package logic

import (
	"ConcurrentProgramming/TaskExecutor/model"
	"ConcurrentProgramming/TaskExecutor/random"
	"math"
)

func TaskAlocation(N uint64) *[]model.Task {
	numberTasks := uint64(math.Pow(10, float64(N)))
	tasks := make([]model.Task, numberTasks)
	return &tasks
}

func FillTasks(tasks *[]model.Task, E uint64) {
	numberTasks := uint64(len(*tasks))
	writingTasks := ((E * numberTasks) / 100)
	readingTasks := numberTasks - writingTasks
	tasksWrite := random.TaskPercentage{WritingTasks: &writingTasks, ReadingTasks: &readingTasks}

	for i := uint64(0); i < numberTasks; i++ {
		(*tasks)[i].Id = uint64(i + 1)
		(*tasks)[i].Cost = random.GenerateCost()
		(*tasks)[i].Type = tasksWrite.GenerateType()
		(*tasks)[i].Value = random.GenerateValue()
	}
}

func SlicePerThread(tasks *[]model.Task, T uint64) *[][]model.Task {
	numberTasks := uint64(len(*tasks))
	sliceSize := numberTasks / T
	extraTasks := numberTasks % T
	sliceTask := make([][]model.Task, T)
	startIndex := uint64(0)

	for i := uint64(0); i < T; i++ {
		realSliceSize := sliceSize
		if i < extraTasks {
			realSliceSize++
		}
		endIndex := startIndex + realSliceSize
		sliceTask[i] = (*tasks)[startIndex:endIndex]
		startIndex = endIndex
	}
	return &sliceTask
}
