package main

import (
	"sync"

	producer "github.com/macqueen01/worker-pool-exercise/producers"
	"github.com/macqueen01/worker-pool-exercise/setups"
	"github.com/macqueen01/worker-pool-exercise/utils"
	worker "github.com/macqueen01/worker-pool-exercise/workers"
)

var wg = &sync.WaitGroup{}

var threadPool = make(chan int, setups.THREAD_POOL_SIZE)

func workerPool() {
	defer wg.Done()
	for i := 0; i < setups.THREAD_POOL_SIZE; i++ {
		wg.Add(1)
		go worker.Work(
			utils.GenerateRandomString(10),
			wg,
			threadPool,
		)
	}
}

func main() {
	wg.Add(1)
	go workerPool()
	go producer.Produce(threadPool, 10)
	wg.Wait()
}
