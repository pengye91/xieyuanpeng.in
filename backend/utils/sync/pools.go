package sync

import "sync"

type Worker interface {
	DoTask()
}

type GoroutinePool struct {
	worker chan Worker
	wg sync.WaitGroup
}


func NewGoPool(goroutineNumber int) *GoroutinePool {
	var pool =  GoroutinePool{
		worker: make(chan Worker),
	}

	pool.wg.Add(goroutineNumber)

	for i := 0; i < goroutineNumber; i++ {
		for w := range pool.worker {
			w.DoTask()
		}
		pool.wg.Done()
	}
	return &pool
}

func (pool *GoroutinePool) Run(w Worker) {
	pool.worker <- w
}

func (pool *GoroutinePool) Shutdown() {
	close(pool.worker)
	pool.wg.Wait()
}
