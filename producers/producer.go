package producer

import "fmt"

func Produce(threadPool chan int, poolSize int) {
	for i := 0; i < poolSize; i++ {
		threadPool <- i // This waits until there is space in the channel.
		fmt.Println("======>> Produced data: ", i)
	}
	close(threadPool)
}
