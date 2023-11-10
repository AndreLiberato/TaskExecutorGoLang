package random

import (
	"ConcurrentProgramming/TaskExecutor/model"
	"math/rand"
)

func GenerateCost() float64 {
	return rand.Float64() * 0.01
}

func GenerateValue() float64 {
	return rand.Float64() * 10
}

type TaskPercentage struct {
	WritingTasks *uint64
	ReadingTasks *uint64
}

func (taskPercentage *TaskPercentage) GenerateType() model.TaskType {
	value := rand.Int63n(2)
	switch value {
	case 0:
		if (*taskPercentage.WritingTasks) > 0 {
			(*taskPercentage.WritingTasks)--
			return model.Write
		} else {
			return model.Read
		}
	case 1:
		if (*taskPercentage.ReadingTasks) > 0 {
			(*taskPercentage.ReadingTasks)--
			return model.Read
		} else {
			return model.Write
		}
	}
	return model.Unknown
}
