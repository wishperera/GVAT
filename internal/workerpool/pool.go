package workerpool

import (
	"context"
	"errors"
	"github.com/wishperera/GVAT/internal/pkg/log"
	"sync"
)

type Pool struct {
	config     Config
	wg         *sync.WaitGroup
	dispatcher *Dispatcher
	log        log.Logger
	shutDown   bool
}

type Config struct {
	MaxWorkers       int
	QueueSize        int
	WorkerBufferSize int
	Process          Process
}

func NewPool(config Config, log log.Logger) (p *Pool, err error) {
	if config.MaxWorkers < 1 {
		return p, errors.New("max workers must at least be 1")
	}
	p = new(Pool)
	p.config = config
	p.wg = new(sync.WaitGroup)
	p.log = log.NewLog("worker-pool")
	p.dispatcher = newDispatcher(dispatcherConfig{
		queueSize:        p.config.QueueSize,
		maxWorkers:       p.config.MaxWorkers,
		workerBufferSize: p.config.WorkerBufferSize,
		wg:               p.wg,
		process:          p.config.Process,
	}, p.log)

	return p, nil
}

func (p *Pool) Init() {
	p.log.Info("worker pool in starting ...")
	p.dispatcher.Init()
	p.log.Info("worker pool in ready state...")
}

func (p *Pool) ShutDown() {
	p.shutDown = true
	p.log.Info("worker pool in shutting down ...")
	p.dispatcher.ShutDown() <- struct{}{}
	p.wg.Wait()
	p.log.Info("worker pool in shutdown gracefully ...")
}

func (p *Pool) ExecuteJob(ctx context.Context, in interface{}) (outChan chan interface{}, errChan chan error, err error) {
	if p.shutDown {
		return nil, nil, errors.New("cannot process request, pool under shutdown")
	}

	if len(p.dispatcher.Queue()) == p.config.QueueSize {
		return nil, nil, errors.New("worker pool exhausted")
	}

	oc := make(chan interface{})
	ec := make(chan error)

	job := jobInput{
		ctx:     ctx,
		input:   in,
		outChan: oc,
		errChan: ec,
	}

	p.dispatcher.Queue() <- job

	return oc, ec, nil
}
