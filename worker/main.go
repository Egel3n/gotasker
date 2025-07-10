package main

import (
	"TaskQueueManager/internal/queue"
	"TaskQueueManager/internal/task"
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	rdb := queue.NewRedisClient()
	task.Init()

	for {
		t, err := queue.DequeueTask(ctx,rdb)

		if err != nil {
			fmt.Print(err)
			return
		}
		err = task.Process(t)
		if err != nil {
			if t.Retry > 0 {
				t.Retry--
				queue.EnqueueTask(ctx,rdb,*t,"gotask:default")
			} else{
				queue.EnqueueTask(ctx,rdb,*t,"gotask:dead")
			}
		} else{
			fmt.Println("Task Done", t.Name)
		}
		
	}

}