package random

import (
	"ConcurrentProgramming/TaskExecutor/model"
	"math/rand"
)

// GenerateCost gera um valor aleatório no intervalo (0.0, 0.01(.
func GenerateCost() float32 {
	return rand.Float32() * 0.01
}

// GenerateValue gera um valor aleatório no intervalo (0.0, 10(.
func GenerateValue() float32 {
	return rand.Float32() * 10
}

// TaskTypeQuantifier é uma estrutura que encapsula o valores de cada tipo de task.
type TaskTypeQuantifier struct {
	WritingTasks *uint64
	ReadingTasks *uint64
}

// GenerateType é o método responsável por gerar o tipo de tarefa de maneira aleatória.
func (taskTypeQuantifier *TaskTypeQuantifier) GenerateType() model.TaskType {
	value := rand.Int31n(2) // Valor no intervalo (0, 2(
	switch value {
	case 0:
		if (*taskTypeQuantifier.WritingTasks) > 0 {
			(*taskTypeQuantifier.WritingTasks)--
			return model.Write
		} else {
			return model.Read
		}
	case 1:
		if (*taskTypeQuantifier.ReadingTasks) > 0 {
			(*taskTypeQuantifier.ReadingTasks)--
			return model.Read
		} else {
			return model.Write
		}
	}
	return model.Unknown
}
