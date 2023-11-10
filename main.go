package main

import (
	"ConcurrentProgramming/TaskExecutor/logic"
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	N, errN := strconv.ParseUint(args[1], 10, 64)
	E, errE := strconv.ParseUint(args[2], 10, 64)
	T, errT := strconv.ParseUint(args[3], 10, 64)

	if errN != nil {
		fmt.Println("Erro na conversão dos valor N:", errN)
	}
	if errE != nil {
		fmt.Println("Erro na conversão dos valor E:", errE)
	}
	if errT != nil {
		fmt.Println("Erro na conversão dos valor T:", errT)
	}

	tasks := logic.TaskAlocation(N)

	logic.FillTasks(tasks, E)

	sliceTask := logic.SlicePerThread(tasks, T)

	for i := uint64(0); i < T; i++ {
		for j := uint64(0); j < uint64(len((*sliceTask)[i])); j++ {
			fmt.Println((*sliceTask)[i][j])
		}
	}
}
