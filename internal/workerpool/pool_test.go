package workerpool

import (
	"context"
	"errors"
	"github.com/wishperera/GVAT/internal/mocks"
	"os"
	"reflect"
	"testing"
)

var pool *Pool
var errorProcess = errors.New("non negative number")

func TestMain(m *testing.M) {
	mockLog := mocks.NewMockLog()
	maxWorkers := 10
	poolBufferSize := 1000
	poolWorkerBuffer := 10
	process := func(ctx context.Context, in interface{}) (out interface{}, err error) {
		input := in.(int)
		if input < 0 {
			return input * -1, nil
		}
		return out, errorProcess
	}
	var err error
	pool, err = NewPool(Config{
		MaxWorkers:       maxWorkers,
		QueueSize:        poolBufferSize,
		WorkerBufferSize: poolWorkerBuffer,
		Process:          process,
	}, mockLog)
	if err != nil {
		os.Exit(1)
	}

	pool.Init()

	code := m.Run()
	pool.ShutDown()
	os.Exit(code)
}

func TestPool_ExecuteJob(t *testing.T) {
	type args struct {
		ctx context.Context
		in  interface{}
	}
	type fields struct {
		pool *Pool
	}

	tests := []struct {
		name    string
		args    args
		fields  fields
		wantOut interface{}
		wantErr error
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				in:  -3,
			},
			fields: fields{
				pool: pool,
			},
			wantOut: 3,
			wantErr: nil,
		},
		{
			name: "job failure",
			args: args{
				ctx: context.Background(),
				in:  2,
			},
			fields: fields{
				pool: pool,
			},
			wantOut: nil,
			wantErr: errorProcess,
		},
	}
	for _, tt := range tests {
		temp := tt
		t.Run(tt.name, func(t *testing.T) {
			gotOutChan, gotErrChan, err := temp.fields.pool.ExecuteJob(temp.args.ctx, temp.args.in)
			if err != nil {
				t.Errorf("failed to execute job due: %s", err)
				return
			}

			var out interface{}

			select {
			case out = <-gotOutChan:
				break
			case err = <-gotErrChan:
				break
			}
			close(gotOutChan)
			close(gotErrChan)

			if !reflect.DeepEqual(out, temp.wantOut) {
				t.Errorf("ExecuteJob() gotOutChan = %v, want %v", out, temp.wantOut)
			}
			if !reflect.DeepEqual(err, temp.wantErr) {
				t.Errorf("ExecuteJob() gotErrChan = %v, want %v", err, temp.wantErr)
			}
		})
	}
}
