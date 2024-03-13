package main

import (
	"context"
	"os"

	"test-task-pip.service/jobs_service/microservice/containers"
)

func main() {
	proc := containers.NewJobProcess()
	proc.SetConfigPath("./config/config.yml")
	proc.Run(context.Background(), os.Args)
}
