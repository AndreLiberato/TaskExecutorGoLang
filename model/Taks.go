package model

type Task struct {
	Id    uint64
	Cost  float64
	Type  TaskType
	Value float64
}
