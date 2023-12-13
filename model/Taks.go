package model

import (
	"ConcurrentProgramming/TaskExecutor/handler"
	"fmt"
	"time"
)

// Taks é a estrutura que encapsula as tarefas.
type Task struct {
	Id    uint64
	Cost  float32
	Type  TaskType
	Value float32
}

// Tasks é um mapa com chave uint64 e valor Taks
type Tasks []Task

// TaskChannel é um canal de task.
type TaskChannel chan Task

// Process é o método de processamento das tarefas.
func (task Task) Process(sharedFile *handler.SharedFile) Result {
	startTime := time.Now() // Marcador de início de excução

	task.waitCost(task.Cost) // Espera o tempo em ms

	var value float32

	switch task.Type {
	case Read:
		sharedFile.FileMutex.RLock()      // Efetua o lock para leitura
		value = sharedFile.LastWriteValue // Recupera o valor da ultima escrita da memória
		sharedFile.FileMutex.RUnlock()    // Libera o lock de leitura
	case Write:
		sharedFile.FileMutex.Lock()                      // Efetua o lock para escrita
		value = task.sumValue(sharedFile.LastWriteValue) // Soma o valor da última escrita com o valor da taks
		task.writeValue(sharedFile, value)               //	Escreve o valor no arquivo e atualiza o valor salvo
		sharedFile.FileMutex.Unlock()                    // Libeara o lock de escrita
	}

	endTime := time.Now()                 // Marcador de fim de execução
	processTime := endTime.Sub(startTime) // Tempo de execução

	// Retorna o resultado
	return Result{Id: task.Id, Value: value, Time: processTime}
}

// waitCost é o método responsável por aguardar um tempo que varia de 0 até 0.01 ms
func (task Task) waitCost(cost float32) {
	duration := time.Duration(cost * 1000000) // Convertendo ns para ms
	time.Sleep(duration)
}

// sumValue é o método responsável por somar o valor da task com o valor lido no arquivo.
func (task Task) sumValue(value float32) float32 {
	return task.Value + value
}

// writeValue é o método responsável por escrever o valor no arquivo.
func (task Task) writeValue(sharedFile *handler.SharedFile, newValue float32) {
	_, err := sharedFile.File.Seek(0, 0) // Retonar o ponteiro do arquivo para o início
	handler.Check(err)

	_, err = sharedFile.File.WriteString(fmt.Sprintf("%f\n", newValue)) // Escreve o novo valor no lugar do antigo
	handler.Check(err)

	sharedFile.LastWriteValue = newValue
}
