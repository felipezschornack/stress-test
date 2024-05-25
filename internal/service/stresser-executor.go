package service

import (
	"net/http"
	"stress-test/internal/model"
	"sync"
	"sync/atomic"
	"time"
)

func Run(input model.StressTestInput) model.StressTestOutput {
	requests := int32(input.NumberOfRequests)
	executionStart := time.Now()
	requestsCount := atomic.Int32{}
	responseStatusCounter := *NewCounter()

	wg := sync.WaitGroup{}
	wg.Add(input.NumberOfRequests)

	makeRequest := func() bool {
		reqNumber := requestsCount.Add(1)
		if reqNumber <= requests {
			defer wg.Done()
			resp, err := http.Get(input.Url)

			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()
			responseStatusCounter.Add(resp.StatusCode)
			return true
		} else {
			requestsCount.Add(-1)
			return false
		}
	}

	for i := 0; i < input.ConcurrencyLevel; i++ {
		go func() {
			for makeRequest() {
			}
		}()
	}

	wg.Wait()

	return model.StressTestOutput{
		ExecutionTime:         time.Since(executionStart),
		RequestsCount:         int(requestsCount.Load()),
		ResponseStatusCounter: responseStatusCounter.GetAll(),
	}
}
