package workerpool

import (
	"fmt"
	"github.com/wishperera/GVAT/internal/pkg/log"
	"sync"
)

type worker struct {
	config workerConfig
	input  chan jobInput
	log    log.Logger
}

type workerConfig struct {
	id               int
	workerBufferSize int
	process          Process
	wg               *sync.WaitGroup
}

func newWorker(conf workerConfig, log log.Logger) *worker {
	w := new(worker)
	w.config = conf
	w.input = make(chan jobInput, conf.workerBufferSize)
	w.log = log.NewLog(fmt.Sprintf("worker-%d", w.config.id))

	return w
}

func (w *worker) Input() chan<- jobInput {
	return w.input
}

func (w *worker) ID() int {
	return w.config.id
}

func (w *worker) Run() {
	w.config.wg.Add(1)
	w.log.Info("worker started...")
	for in := range w.input {
		out, err := w.config.process(in.ctx, in.input)
		if err != nil {
			in.errChan <- err
		} else {
			in.outChan <- out
		}
		w.log.InfoContext(in.ctx, "job executed...")
	}
	w.log.Info("worker shutdown...")
	w.config.wg.Done()
}
