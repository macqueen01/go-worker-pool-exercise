package tasks

import (
	"fmt"
	"time"
)

func SampleTask(id string, data interface{}) {
	fmt.Println("======>> Thread Started -- SampleTask: ", id)
	for i := 0; i < 5; i++ {
		time.Sleep(3 * time.Second)
		fmt.Println("======>> Thread in progress -- SampleTask: ", id)
	}
	fmt.Println("======>> Thread Ended -- SampleTask: ", id)
}
