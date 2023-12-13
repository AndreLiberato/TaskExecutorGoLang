package logic

import (
	"ConcurrentProgramming/TaskExecutor/handler"
	"ConcurrentProgramming/TaskExecutor/model"
	"fmt"
	"sync"
	"time"
)

func Init(N float64, E uint64, T uint64) *model.Results {
	tasks, numberTasks := TasksInstance(N)
	results := make(model.Results, numberTasks)

	FillTasks(tasks, numberTasks, E)

	tasksChannel := make(model.TaskChannel, 3*T)     // Canal de tasks
	resultsChannel := make(model.ResultChannel, 5*T) // Canal de results

	waitGroup := new(sync.WaitGroup) // Ponto de sincronização
	waitReceiver := new(sync.WaitGroup)
	// Executor
	executor := model.Executor{
		Tasks:          tasks,
		Results:        &results,
		TasksChannel:   tasksChannel,
		ResultsChannel: resultsChannel,
		WaitGroup:      waitGroup,
		WaitReceiver:   waitReceiver,
	}

	sharedFile := handler.SharedFileInstance() // Arquivo compartilhado

	sharedFile.WriteInitalValue() // Escrevendo o valor inicial no arquivo

	// Trabalhador
	worker := model.Worker{
		TasksChannel:   tasksChannel,
		ResultsChannel: resultsChannel,
		WaitGroup:      waitGroup,
		SharedFile:     sharedFile,
	}

	// Iteração inicialização dos workers
	for i := uint64(0); i < T; i++ {
		waitGroup.Add(1)
		go worker.Work() // Inicia e fica aguardando o executor inciar o sender
	}

	startTime := time.Now() // Marcador do tempo de inicio da etapa de processamento

	executor.Execute() // Executa o processamento

	endTime := time.Now() // Marcador de tempo de fim da etapa de processamento

	processingStageDuration := endTime.Sub(startTime).Milliseconds() // Calcula a duração da etapa de processamento

	fmt.Println("Duração da etapa de processamento das tarefas:", processingStageDuration)

	waitReceiver.Wait()

	return &results
}

func CalculateExecutionTime(results *model.Results) {
	totalProcessingTime := time.Duration(0)

	for _, result := range *results {
		totalProcessingTime += result.Time
	}

	fmt.Println("Soma do tempo de processamento das tarefas:", totalProcessingTime.Milliseconds(), "ms.")
}
