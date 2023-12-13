package model

import (
	"fmt"
	"sync"
)

// Executor encapsula as informações necessárias para seu fluxo de excução.
type Executor struct {
	TasksChannel   TaskChannel
	ResultsChannel ResultChannel
	Tasks          *Tasks
	Results        *Results
	WaitGroup      *sync.WaitGroup
	WaitReceiver   *sync.WaitGroup
}

// Execute é a função responsável pelo controle do fluxo de execução do Executor.
func (executor Executor) Execute() {
	fmt.Println("Iniciando trabalho do executor.")

	go executor.sender(executor.TasksChannel) // Inicia GoRoutine para enviar as tasks

	executor.WaitReceiver.Add(1)

	go executor.receiver(executor.ResultsChannel) // Inicia GoRoutine para recever os results

	executor.WaitGroup.Wait() // Espere todas os workers acabarem

	close(executor.ResultsChannel) // Fecha o canal de results

	fmt.Println("Trabalho do executor finalizado.")
}

// sender é o método responsável por enviar as tarefas para o canal de tasks.
func (executor Executor) sender(tasksChannel chan<- Task) {
	defer close(tasksChannel)
	for _, task := range *executor.Tasks {
		tasksChannel <- task // Envia para o canal
	}
}

// receiver é o método responsável por receber os resultados do canal de resultados.
func (executor Executor) receiver(resultsChannel <-chan Result) {
	keyResult := uint64(0) // Chave inicial
	for result := range resultsChannel {
		(*executor.Results)[keyResult] = result
		keyResult++
	}
	executor.WaitReceiver.Done()
}
