package model

type TaskType uint8

const (
	Read TaskType = iota
	Write
	Unknown
)

func (tt TaskType) String() string {
	switch tt {
	case Read:
		return "read"
	case Write:
		return "write"
	}
	return "unknown"
}
