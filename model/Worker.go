package model

import (
	"ConcurrentProgramming/TaskExecutor/handler"
	"sync"
)

// Worker é a estrutura que encapsula as informações necessárias para o seu fluxo de trabalho.
type Worker struct {
	TasksChannel   TaskChannel
	ResultsChannel ResultChannel
	WaitGroup      *sync.WaitGroup
	SharedFile     *handler.SharedFile
}

// Work é o método que determina o trabalho do Worker.
func (worker Worker) Work() {
	defer worker.WaitGroup.Done()
	for task := range worker.TasksChannel {
		worker.ResultsChannel <- task.Process(worker.SharedFile)
	}
}
