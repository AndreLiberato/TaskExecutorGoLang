package model

import (
	"ConcurrentProgramming/TaskExecutor/handler"
	"fmt"
	"os"
	"sync"
)

// Executor encapsula as informações necessárias para seu fluxo de excução.
type Executor struct {
	Tasks       *Tasks
	Results     Results
	T           uint64
	NumberTasks uint64
}

// Execute é a função responsável pelo controle do fluxo de execução do Executor.
func (executor Executor) Execute() {
	tasksChannel := make(TaskChannel, executor.T)     // Canal de tasks
	resultsChannel := make(ResultChannel, executor.T) // Canal de results

	// Criando arquivo
	file, err := os.Create("value.data")
	handler.Check(err)
	defer file.Close()

	// Escrevendo o valor inicial no arquivo
	_, err = file.WriteString(fmt.Sprintf("%f\n", 0.0))
	handler.Check(err)

	waitGroup := new(sync.WaitGroup) // Ponto de sincronização

	fileMutex := new(sync.RWMutex) // Mutex de controle de escrita e leitura do arquivo

	// Iteração para criação einicialização dos workers
	for i := uint64(0); i < executor.T; i++ {
		worker := Worker{
			tasksChannel:   tasksChannel,
			resultsChannel: resultsChannel,
			waitGroup:      waitGroup,
			fileMutex:      fileMutex,
		}
		waitGroup.Add(1)
		go worker.Work(file)
	}

	go executor.sender(tasksChannel)     // Inicia GoRoutine para enviar as tasks
	go executor.receiver(resultsChannel) // Inicia GoRoutine para recever os results

	waitGroup.Wait() // Espere todas os workers acabarem

	defer close(resultsChannel) // Fecha o canal de results
}

// sender é o método responsável por enviar as tarefas para o canal de tasks.
func (executor Executor) sender(tasksChannel chan<- Task) {
	defer close(tasksChannel)
	for i := uint64(0); i < executor.NumberTasks; i++ {
		task := (*executor.Tasks)[i] // Atribuí a task
		tasksChannel <- task         // Envia para o canal
		delete((*executor.Tasks), i) // Remove da estrutura que armazena as tasks
	}
}

// receiver é o método responsável por receber os resultados do canal de resultados.
func (executor Executor) receiver(resultsChannel <-chan Result) {
	keyResult := uint64(0) // Chave inicial
	for result := range resultsChannel {
		executor.Results[keyResult] = result
		keyResult++
	}
}
