package queue

import (
	"TaskQueueManager/internal/task"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)


func NewRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:	  "localhost:6379",
		Password: "", // No password set
		DB:		  0,  // Use default DB
		Protocol: 2,  // Connection protocol
	})

	return client;
}

func EnqueueTask(ctx context.Context, rdb *redis.Client, t task.Task, queueName string) error {
	//queueName: gotask:default, gotask:dead
	taskJson, err := json.Marshal(t);
	if(err!=nil){
		fmt.Printf("Marshalling Error");
	}
	return rdb.RPush(ctx,queueName,taskJson).Err()
}



func DequeueTask(ctx context.Context, rdb *redis.Client) (*task.Task, error){
	res,err := rdb.BRPop(ctx,0*time.Second,"gotask:default").Result();
	if err != nil {
		return nil,err;
	}

	var t task.Task;
	if err := json.Unmarshal([]byte(res[1]),&t); err != nil {
		return nil, err;
	}
	
	return &t , nil;
}



