package main

import (
	"ConcurrentProgramming/TaskExecutor/handler"
	"ConcurrentProgramming/TaskExecutor/logic"
	"ConcurrentProgramming/TaskExecutor/model"
	"fmt"
	"os"
	"strconv"
	"time"
)

// Main é o entrypoint do programa.
// Responsável  pelo fluxo principal de execução
func main() {
	startTime := time.Now()
	fmt.Println("Iniciando programa.")
	args := os.Args

	N, errN := strconv.ParseUint(args[1], 10, 64)
	handler.Check(errN)

	E, errE := strconv.ParseUint(args[2], 10, 64)
	handler.Check(errE)

	T, errT := strconv.ParseUint(args[3], 10, 64)
	handler.Check(errT)

	tasks, numberTasks := logic.TaskAlocation(N)

	fmt.Println("Gerando tarefas.")
	logic.FillTasks(tasks, numberTasks, E)

	results := model.Results{}

	executor := model.Executor{
		Tasks:       tasks,
		Results:     results,
		T:           T,
		NumberTasks: numberTasks,
	}

	fmt.Println("Iniciando trabalho do executor.")
	executor.Execute()
	fmt.Println("Trabalho do executor finalizado.")

	totalProcessingTime := time.Duration(0)
	fmt.Println("Calculando tempo total de processamento de tarefas.")
	for _, result := range results {
		totalProcessingTime += result.Time
	}

	fmt.Println("Tempo total de processamento de tarefas:", totalProcessingTime.Milliseconds(), "ms.")

	endTime := time.Now()
	fmt.Println("Tempo total do programa:", endTime.Sub(startTime).Milliseconds(), "ms.")
}
