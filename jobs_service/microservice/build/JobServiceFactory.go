package build

import (
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	cbuild "github.com/pip-services3-gox/pip-services3-components-gox/build"
	"test-task-pip.service/jobs_service/microservice/logic"
	"test-task-pip.service/jobs_service/microservice/persistence"
	service1 "test-task-pip.service/jobs_service/microservice/service/version1"
)

type JobServiceFactory struct {
	cbuild.Factory
}

func NewJobServiceFactory() *JobServiceFactory {
	c := &JobServiceFactory{
		Factory: *cbuild.NewFactory(),
	}

	sqlitePersistanceDescriptor := cref.NewDescriptor("job", "persistence", "sqlite", "*", "1.0")
	controllerDescriptor := cref.NewDescriptor("job", "controller", "default", "*", "1.0")
	httpServiceV1Descriptor := cref.NewDescriptor("job", "service", "http", "*", "1.0")

	c.RegisterType(sqlitePersistanceDescriptor, persistence.NewJobSqlitePersistence)
	c.RegisterType(controllerDescriptor, logic.NewJobController)
	c.RegisterType(httpServiceV1Descriptor, service1.NewJobHttpServiceV1)

	return c
}
