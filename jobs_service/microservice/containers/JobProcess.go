package containers

import (
	cproc "github.com/pip-services3-gox/pip-services3-container-gox/container"
	rbuild "github.com/pip-services3-gox/pip-services3-rpc-gox/build"
	sqlite "github.com/pip-services3-gox/pip-services3-sqlite-gox/build"
	cswagger "github.com/pip-services3-gox/pip-services3-swagger-gox/build"
	"test-task-pip.service/jobs_service/microservice/build"
)

type JobProcess struct {
	cproc.ProcessContainer
}

func NewJobProcess() *JobProcess {
	c := &JobProcess{
		ProcessContainer: *cproc.NewProcessContainer("job", "Job microservice"),
	}

	c.AddFactory(build.NewJobServiceFactory())
	c.AddFactory(rbuild.NewDefaultRpcFactory())
	c.AddFactory(sqlite.NewDefaultSqliteFactory())
	c.AddFactory(cswagger.NewDefaultSwaggerFactory())

	return c
}
