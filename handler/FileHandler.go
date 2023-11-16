package handler

import (
	"fmt"
	"os"
	"sync"
)

// SharedFile encapsula o ponteiro para o arquivo e o RWMutex
type SharedFile struct {
	File      *os.File
	FileMutex *sync.RWMutex
}

// SharedFileInstance cria e retorna uma instância de SharedFile
func SharedFileInstance() SharedFile {
	sharedFile := SharedFile{
		File:      createFile(),
		FileMutex: new(sync.RWMutex),
	}
	return sharedFile
}

// createFile cria o arquivo com o nome value.data
func createFile() *os.File {
	file, err := os.Create("value.data")
	Check(err)
	return file
}

// WriteInitalValue escreve o valor inicial no arquivo
func (sharedFile SharedFile) WriteInitalValue() {
	// Escrevendo o valor inicial no arquivo
	_, err := sharedFile.File.WriteString(fmt.Sprintf("%f\n", 0.0))
	Check(err)
}
