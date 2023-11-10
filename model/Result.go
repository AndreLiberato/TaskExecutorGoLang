package model

import "time"

type Result struct {
	Id    uint64
	Value uint8
	Time  time.Duration
}
