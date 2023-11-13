package model

import (
	"ConcurrentProgramming/TaskExecutor/handler"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

// Taks é a estrutura que encapsula as tarefas.
type Task struct {
	Id    uint64
	Cost  float64
	Type  TaskType
	Value float64
}

// Tasks é um mapa com chave uint64 e valor Taks
type Tasks map[uint64]Task

// TaskChannel é um canal de task.
type TaskChannel chan Task

// Process é o método de processamento das tarefas.
func (task Task) Process(file *os.File, writeMutex *sync.RWMutex) Result {
	startTime := time.Now() // Marcador de início de excução

	task.waitCost(task.Cost) // Espera o tempo em ms

	value := float64(0.0)

	switch task.Type {
	case Read:
		writeMutex.RLock()           // Efetua lock para leitura
		value = task.readValue(file) // Ler valor do arquivo
		writeMutex.RUnlock()         // Libera a trava arquivo
	case Write:
		writeMutex.Lock()            // Efetua o lock para escrita
		value = task.readValue(file) // Ler valor do arquivo
		value = task.sumValue(value) // Soma o valor do arquivo com o valor da task
		task.writeValue(file, value) //	Escreve o valor no arquivo
		writeMutex.Unlock()          // Libeara a trava de escrita do arquivo
	}

	endTime := time.Now()                 // Marcador de fim de execução
	processTime := endTime.Sub(startTime) // Tempo de execução

	// Retorna o resultado
	return Result{Id: task.Id, Value: value, Time: processTime}
}

// waitCost é o método responsável por aguardar um tempo que varia de 0 até 0.01 ms
func (task Task) waitCost(cost float64) {
	duration := time.Duration(cost * 1000000) // Convertendo ns para ms
	time.Sleep(duration)
}

// readValue é o método responsável pela leitura do valor existente no arquivo.
func (task Task) readValue(file *os.File) float64 {
	_, err := file.Seek(0, 0) // Retorna o ponteiro do arquivo para o início
	handler.Check(err)

	scanner := bufio.NewScanner(file) // Scanner do arquivo
	if scanner.Scan() {
		valueString := scanner.Text()                     // Recupera a informação do arquivo em formato de texto
		value, err := strconv.ParseFloat(valueString, 64) // Converte o valor em string para float64
		handler.Check(err)
		return value
	}
	return 0
}

// sumValue é o método responsável por somar o valor da task com o valor lido no arquivo.
func (task Task) sumValue(value float64) float64 {
	return task.Value + value
}

// writeValue é o método responsável por escrever o valor no arquivo.
func (task Task) writeValue(file *os.File, newValue float64) {
	_, err := file.Seek(0, 0) // Retonar o ponteiro do arquivo para o início
	handler.Check(err)

	_, err = file.WriteString(fmt.Sprintf("%f\n", newValue)) // Escreve o novo valor no lugar do antigo
	handler.Check(err)
}
