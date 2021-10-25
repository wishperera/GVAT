package workerpool

import (
	"fmt"
	"github.com/wishperera/GVAT/internal/pkg/log"
)

type Worker struct {
	id       int
	input    chan jobInput
	shutDown chan struct{}
	process  Process
	pool     *Pool
	log      log.Logger
}

func NewWorker(id int, process Process, pool *Pool, log log.Logger) *Worker {
	w := new(Worker)
	w.id = id
	w.process = process
	w.pool = pool
	w.input = make(chan jobInput)
	w.shutDown = make(chan struct{})
	w.log = log.NewLog(fmt.Sprintf("worker-%d", w.id))

	return w
}

func (w *Worker) Input() chan jobInput {
	return w.input
}

func (w *Worker) ShutDown() chan struct{} {
	return w.shutDown
}

func (w *Worker) ID() int {
	return w.id
}

func (w *Worker) Run() {
	w.pool.wg.Add(1)
	w.log.Info("worker started...")
	for {
		select {
		case in := <-w.input:
			out, err := w.process(in.ctx, in.input)
			if err != nil {
				in.errChan <- err
			} else {
				in.outChan <- out
			}
			w.log.InfoContext(in.ctx, "job executed...")
		case <-w.shutDown:
			w.log.Info("worker shutdown...")
			w.pool.wg.Done()
			return
		}
	}
}
