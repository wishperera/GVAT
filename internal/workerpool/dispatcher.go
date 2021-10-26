package workerpool

import (
	"github.com/wishperera/GVAT/internal/pkg/log"
	"sync"
)

type Dispatcher struct {
	index    int
	config   dispatcherConfig
	queue    chan jobInput
	workers  []*worker
	wg       *sync.WaitGroup
	log      log.Logger
	shutdown chan struct{}
}

type dispatcherConfig struct {
	queueSize        int
	maxWorkers       int
	workerBufferSize int
	wg               *sync.WaitGroup
	process          Process
}

func newDispatcher(config dispatcherConfig, log log.Logger) *Dispatcher {
	d := new(Dispatcher)
	d.config = config
	d.workers = make([]*worker, d.config.maxWorkers)
	d.queue = make(chan jobInput, d.config.queueSize)
	d.wg = new(sync.WaitGroup)
	d.log = log.NewLog("dispatcher")
	d.shutdown = make(chan struct{})

	for i := 0; i < d.config.maxWorkers; i++ {
		wk := newWorker(workerConfig{
			id:               i,
			workerBufferSize: d.config.workerBufferSize,
			process:          d.config.process,
			wg:               d.wg,
		}, log)
		d.workers[i] = wk
	}

	return d
}

func (d *Dispatcher) Init() {
	for i := range d.workers {
		go d.workers[i].Run()
	}
	go d.run()
}

// nolint //intentional
func (d *Dispatcher) Queue() chan<- jobInput {
	return d.queue
}

func (d *Dispatcher) ShutDown() chan<- struct{} {
	return d.shutdown
}

func (d *Dispatcher) run() {
	d.log.Info("dispatcher started...")
	d.config.wg.Add(1)
	for {
		select {
		case in := <-d.queue:
			wk := d.workers[d.index]
			d.index++
			if d.index > d.config.maxWorkers-1 {
				d.index = 0
			}

			d.log.InfoContext(in.ctx, "job assigned..", d.log.Param("worker-id", wk.config.id))

			wk.Input() <- in
		case <-d.shutdown:
			for _, v := range d.workers {
				v.ShutDown() <- struct{}{}
			}
			d.wg.Wait()
			d.log.Info("dispatcher shut down...")
			d.config.wg.Done()
		}
	}
}
