package build

import (
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	cbuild "github.com/pip-services3-gox/pip-services3-components-gox/build"
	"test-task-pip.service/worker_service/microservice/logic"
	"test-task-pip.service/worker_service/microservice/persistence"
	service1 "test-task-pip.service/worker_service/microservice/service/version1"
)

type WorkerServiceFactory struct {
	cbuild.Factory
}

func NewWorkerServiceFactory() *WorkerServiceFactory {
	c := &WorkerServiceFactory{
		Factory: *cbuild.NewFactory(),
	}

	persistanceDescriptor := cref.NewDescriptor("worker", "persistence", "default", "*", "1.0")
	controllerDescriptor := cref.NewDescriptor("worker", "controller", "default", "*", "1.0")
	httpServiceV1Descriptor := cref.NewDescriptor("worker", "service", "http", "*", "1.0")

	c.RegisterType(persistanceDescriptor, persistence.NewWorkerPersistence)
	c.RegisterType(controllerDescriptor, logic.NewWorkerController)
	c.RegisterType(httpServiceV1Descriptor, service1.NewWorkerHttpServiceV1)

	return c
}
