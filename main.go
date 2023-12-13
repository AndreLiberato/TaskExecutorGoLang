package main

import (
	"ConcurrentProgramming/TaskExecutor/handler"
	"ConcurrentProgramming/TaskExecutor/logic"
	"fmt"
	"os"
	"runtime/pprof"
	"strconv"
	"time"
)

// Main é o entrypoint do programa.
// Responsável  pelo fluxo principal de execução
func main() {
	fcpu, err := os.Create("cpu.prof")
	handler.Check(err)
	defer fcpu.Close()
	pprof.StartCPUProfile(fcpu)
	defer pprof.StopCPUProfile()

	startTime := time.Now()

	fmt.Println("Iniciando programa.")

	args := os.Args

	N, errN := strconv.ParseFloat(args[1], 64)
	handler.Check(errN)

	E, errE := strconv.ParseUint(args[2], 10, 64)
	handler.Check(errE)

	T, errT := strconv.ParseUint(args[3], 10, 64)
	handler.Check(errT)

	results := logic.Init(N, E, T)

	logic.CalculateExecutionTime(results)

	endTime := time.Now()

	fmt.Println("Tempo de execução total do programa:", endTime.Sub(startTime).Milliseconds(), "ms.")

	fmem, err := os.Create("mem.prof")
	handler.Check(err)
	defer fmem.Close()
	pprof.WriteHeapProfile(fmem)
}
