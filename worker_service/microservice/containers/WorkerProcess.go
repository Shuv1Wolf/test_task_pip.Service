package containers

import (
	cproc "github.com/pip-services3-gox/pip-services3-container-gox/container"
	rbuild "github.com/pip-services3-gox/pip-services3-rpc-gox/build"
	cswagger "github.com/pip-services3-gox/pip-services3-swagger-gox/build"
	"test-task-pip.service/worker_service/microservice/build"
)

type WorkerProcess struct {
	cproc.ProcessContainer
}

func NewWorkerProcess() *WorkerProcess {
	c := &WorkerProcess{
		ProcessContainer: *cproc.NewProcessContainer("worker", "Worker microservice"),
	}

	c.AddFactory(build.NewWorkerServiceFactory())
	c.AddFactory(rbuild.NewDefaultRpcFactory())
	c.AddFactory(build.NewJobClientFactory())
	c.AddFactory(build.NewKeyClientFactory())
	c.AddFactory(cswagger.NewDefaultSwaggerFactory())

	return c
}
