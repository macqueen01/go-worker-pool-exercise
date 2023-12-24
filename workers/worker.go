package worker

import (
	"sync"

	"github.com/macqueen01/worker-pool-exercise/tasks"
)

func Work(id string, wg *sync.WaitGroup, threadPool <-chan int) {
	defer wg.Done()

	for task := range threadPool {
		tasks.SampleTask(id, task)
	}
}
