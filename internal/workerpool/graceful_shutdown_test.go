package workerpool

import (
	"context"
	"github.com/wishperera/GVAT/internal/mocks"
	"os"
	"sync"
	"sync/atomic"
	"testing"
)

func TestGracefulShutdown(t *testing.T) {
	var responseCount int64
	wg := new(sync.WaitGroup)

	mockLog := mocks.NewMockLog()
	maxWorkers := 10
	poolBufferSize := 1000
	poolWorkerBuffer := 10
	process := func(ctx context.Context, in interface{}) (out interface{}, err error) {
		input := in.(int)
		if input >= 0 {
			return input * -1, nil
		}
		return out, errorProcess
	}

	tempPool, err := NewPool(Config{
		MaxWorkers:       maxWorkers,
		QueueSize:        poolBufferSize,
		WorkerBufferSize: poolWorkerBuffer,
		Process:          process,
	}, mockLog)
	if err != nil {
		os.Exit(1)
	}

	tempPool.Init()

	jobCount := 1000

	for i := 0; i < 1000; i++ {
		oc, ec, err := tempPool.ExecuteJob(context.Background(), i)
		if err != nil {
			t.Error(err)
			t.Fail()
		}
		wg.Add(1)
		go func(t *testing.T) {
			select {
			case <-oc:
				atomic.AddInt64(&responseCount, 1)
			case processError := <-ec:
				t.Log(processError)
			}
			close(oc)
			close(ec)
			wg.Done()
		}(t)
	}
	tempPool.ShutDown()
	wg.Wait()

	if int64(jobCount) != responseCount {
		t.Error("pool did not shutdown gracefully")
		return
	}
}
