package model

// TaksType é um enumerador que representa os tipos de tarefa (leitura ou escrita).
type TaskType uint8

// Enumerador
const (
	Read TaskType = iota
	Write
	Unknown
)

// Função para conversão de valor em string
func (tt TaskType) String() string {
	switch tt {
	case Read:
		return "read"
	case Write:
		return "write"
	}
	return "unknown"
}
