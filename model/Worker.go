package model

import (
	"os"
	"sync"
)

// Worker é a estrutura que encapsula as informações necessárias para o seu fluxo de trabalho.
type Worker struct {
	tasksChannel   TaskChannel
	resultsChannel ResultChannel
	waitGroup      *sync.WaitGroup
	fileMutex      *sync.RWMutex
}

// Work é o método que determina o trabalho do Worker.
func (worker Worker) Work(file *os.File) {
	defer worker.waitGroup.Done()
	for task := range worker.tasksChannel {
		worker.resultsChannel <- task.Process(file, worker.fileMutex)
	}
}
