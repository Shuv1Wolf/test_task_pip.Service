package main

import (
	"os"

	"test-task-pip.service/facade_service/microservice/container"
)

func main() {
	proc := container.NewFacadeProcess()
	proc.SetConfigPath("./config/config.yml")
	proc.Run(os.Args)
}
