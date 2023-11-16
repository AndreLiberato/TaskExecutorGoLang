package main

import (
	"ConcurrentProgramming/TaskExecutor/handler"
	"ConcurrentProgramming/TaskExecutor/logic"
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

	results := logic.Init(N, E, T)

	logic.CalculateExecutionTime(results)

	endTime := time.Now()

	fmt.Println("Tempo de execução total do programa:", endTime.Sub(startTime).Milliseconds(), "ms.")
}
