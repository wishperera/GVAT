package workerpool

import "context"

type Process func(ctx context.Context, in interface{}) (out interface{}, err error)

type jobInput struct {
	ctx     context.Context
	input   interface{}
	outChan chan interface{}
	errChan chan error
}
