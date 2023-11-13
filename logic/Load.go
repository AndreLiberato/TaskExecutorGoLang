package logic

import (
	"ConcurrentProgramming/TaskExecutor/model"
	"ConcurrentProgramming/TaskExecutor/random"
	"math"
)

// TaskAlocation configura a capacidade inicial de Tasks.
// Não aloca memória
func TaskAlocation(N uint64) (*model.Tasks, uint64) {
	numberTasks := uint64(math.Pow(10, float64(N)))
	tasks := make(model.Tasks, numberTasks)
	return &tasks, numberTasks
}

// FillTasks cria aleatóriamente as tasks.
func FillTasks(tasks *model.Tasks, numberTasks uint64, E uint64) {
	writingTasks := ((E * numberTasks) / 100)  // Tasks de escrita
	readingTasks := numberTasks - writingTasks // Tasks de leitura

	// Objeto que encapsula a quantidade de tasks de cada tipo
	taskType := random.TaskTypeQuantifier{
		WritingTasks: &writingTasks,
		ReadingTasks: &readingTasks,
	}

	// Criando as tasks
	for i := uint64(0); i < numberTasks; i++ {
		(*tasks)[i] = model.Task{
			Id:    uint64(i + 1),
			Cost:  random.GenerateCost(),
			Type:  taskType.GenerateType(),
			Value: random.GenerateValue(),
		}
	}
}
