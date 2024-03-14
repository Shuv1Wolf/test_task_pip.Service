package main

import (
	"context"
	"os"

	"test-task-pip.service/worker_service/microservice/containers"
)

func main() {
	proc := containers.NewWorkerProcess()
	proc.SetConfigPath("./config/config.yml")
	proc.Run(context.Background(), os.Args)
}
