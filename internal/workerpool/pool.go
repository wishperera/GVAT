package workerpool

import (
	"context"
	"errors"
	"github.com/wishperera/GVAT/internal/pkg/log"
	"sync"
)

type Pool struct {
	index      int
	maxWorkers int
	process    Process
	workers    []*Worker
	wg         *sync.WaitGroup
	log        log.Logger
}

func NewPool(maxWorkers int, process Process, log log.Logger) (p *Pool, err error) {
	if maxWorkers < 1 {
		return p, errors.New("max workers must at least be 1")
	}
	p = new(Pool)
	p.maxWorkers = maxWorkers
	p.process = process
	p.workers = make([]*Worker, maxWorkers)
	p.wg = new(sync.WaitGroup)
	p.log = log.NewLog("worker-pool")

	for i := 0; i < maxWorkers; i++ {
		wk := NewWorker(i, p.process, p, p.log)
		p.workers[i] = wk
	}

	return p, nil
}

func (p *Pool) Init() {
	p.log.Info("worker pool in starting ...")
	for i := range p.workers {
		go p.workers[i].Run()
	}
	p.log.Info("worker pool in ready state...")
}

func (p *Pool) ShutDown() {
	p.log.Info("worker pool in shutting down ...")
	for i := range p.workers {
		p.workers[i].ShutDown() <- struct{}{}
	}
	p.wg.Wait()
	p.log.Info("worker pool in shutdown gracefully ...")
}

func (p *Pool) ExecuteJob(ctx context.Context, in interface{}) (outChan chan interface{}, errChan chan error) {
	wk := p.workers[p.index]
	p.index += 1
	if p.index > p.maxWorkers-1 {
		p.index = 0
	}
	outChan = make(chan interface{})
	errChan = make(chan error)

	p.log.InfoContext(ctx, "job assigned..", p.log.Param("worker-id", wk.id))

	wk.Input() <- jobInput{
		ctx:     ctx,
		input:   in,
		outChan: outChan,
		errChan: errChan,
	}

	return outChan, errChan
}
