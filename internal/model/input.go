package model

type StressTestInput struct {
	Url              string
	NumberOfRequests int
	ConcurrencyLevel int
}
