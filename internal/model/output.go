package model

import (
	"time"
)

type StressTestOutput struct {
	ExecutionTime         time.Duration
	RequestsCount         int
	ResponseStatusCounter map[int]int
}
