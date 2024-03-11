package containers

import (
	cproc "github.com/pip-services3-gox/pip-services3-container-gox/container"
	rbuild "github.com/pip-services3-gox/pip-services3-rpc-gox/build"
	sqlite "github.com/pip-services3-gox/pip-services3-sqlite-gox/build"
	cswagger "github.com/pip-services3-gox/pip-services3-swagger-gox/build"
	"test_task_pip.Service/keystore_service/microservice/build"
)

type KeyProcess struct {
	cproc.ProcessContainer
}

func NewKeyProcess() *KeyProcess {
	c := &KeyProcess{
		ProcessContainer: *cproc.NewProcessContainer("key", "Key microservice"),
	}

	c.AddFactory(build.NewKeyServiceFactory())
	c.AddFactory(rbuild.NewDefaultRpcFactory())
	c.AddFactory(sqlite.NewDefaultSqliteFactory())
	c.AddFactory(cswagger.NewDefaultSwaggerFactory())

	return c
}
