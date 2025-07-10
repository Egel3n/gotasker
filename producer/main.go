package main

import (
	api "TaskQueueManager/api"
	"log"
	"net/http"
)




func main() {

	mux := http.NewServeMux()
    mux.HandleFunc("/enqueue", api.HandleEnqueueTask)

    log.Println("Server started on :8080")
    http.ListenAndServe(":8080", mux)
	
	
}

