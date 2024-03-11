package main

import (
	"context"
	"os"

	"test_task_pip.Service/keystore_service/microservice/containers"
)

func main() {
	proc := containers.NewKeyProcess()
	proc.SetConfigPath("./config/config.yml")
	proc.Run(context.Background(), os.Args)
}
