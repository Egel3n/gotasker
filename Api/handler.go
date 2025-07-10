package api

import (
	"TaskQueueManager/internal/queue"
	"TaskQueueManager/internal/task"
	"context"
	"encoding/json"
	"net/http"
)

var rdb = queue.NewRedisClient();

func HandleEnqueueTask(w http.ResponseWriter, r *http.Request) {

	if(r.Method != http.MethodPost) {
		http.Error(w,"Method Not Allowed Only POST",http.StatusMethodNotAllowed)
	}

	var t task.Task
	err := json.NewDecoder(r.Body).Decode(&t)
	if(err != nil) {
		http.Error(w,"Invalid Json",http.StatusBadRequest)
	}

	if t.Name == "" || t.Args == nil {
        http.Error(w, "Missing required fields", http.StatusBadRequest)
        return
    }

	ctx := context.Background();



	err = queue.EnqueueTask(ctx,rdb,t,"gotask:default");

	if err != nil{
		http.Error(w, "Failed to Enqueue Task. "+ err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "status": "queued",
    })

}