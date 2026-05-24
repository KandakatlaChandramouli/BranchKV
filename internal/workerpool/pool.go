package workerpool

import (
	"sync"
)

type Job func()

type WorkerPool struct {
	jobs chan Job
	wg   sync.WaitGroup
}

func NewWorkerPool(
	workers int,
	queue int,
) *WorkerPool {

	p := &WorkerPool{
		jobs: make(chan Job, queue),
	}

	for i := 0; i < workers; i++ {

		go func() {

			for job := range p.jobs {
				job()
				p.wg.Done()
			}
		}()
	}

	return p
}

func (p *WorkerPool) Submit(
	job Job,
) {

	p.wg.Add(1)

	p.jobs <- job
}

func (p *WorkerPool) Wait() {
	p.wg.Wait()
}
