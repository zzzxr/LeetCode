package main

/*
线程池
*/
import (
	"fmt"
	"sync"
)

type WorkerPool struct {
	Workers   []*Worker
	WorkQueue chan func()
	wg        sync.WaitGroup
}

type Worker struct {
	ID   int
	Pool *WorkerPool
	Job  chan func()
	wg   sync.WaitGroup
}

func NewWorker(ID int, pool *WorkerPool) *Worker {
	worker := &Worker{
		ID:   ID,
		Pool: pool,
		Job:  make(chan func()),
	}
	go worker.start()
	return worker
}

func (w *Worker) start() {
	for job := range w.Job {
		job()
		w.Pool.wg.Done()
	}
	w.wg.Done()
}

func NewWorkerPool(numWorkers int) *WorkerPool {
	pool := &WorkerPool{
		WorkQueue: make(chan func(), numWorkers),
	}

	for i := 0; i < numWorkers; i++ {
		worker := NewWorker(i, pool)
		pool.Workers = append(pool.Workers, worker)
	}

	pool.wg.Add(numWorkers)
	for _, worker := range pool.Workers {
		worker.wg.Wait()
	}
	return pool
}

func (p *WorkerPool) Submit(job func()) {
	p.wg.Add(1)
	p.WorkQueue <- job
}

func main() {
	pool := NewWorkerPool(5) // 创建一个包含 5 个工作线程的线程池

	for i := 0; i < 10; i++ {
		jobID := i
		pool.Submit(func() {
			fmt.Printf("Job %d is processed by worker %d\n", jobID, pool.Workers[jobID%5].ID)
		})
	}

	pool.wg.Wait() // 等待所有任务完成
}
