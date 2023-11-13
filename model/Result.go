package model

import "time"

// Result é a estrutura que encapsula um resultado.
type Result struct {
	Id    uint64
	Value float64
	Time  time.Duration
}

// Results é um mapa com chave uint64 e valor Result.
type Results map[uint64]Result

// ResultChannvel é um cananl de result.
type ResultChannel chan Result
