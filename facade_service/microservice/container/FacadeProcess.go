package container

import (
	cproc "github.com/pip-services3-go/pip-services3-container-go/container"
	rpcbuild "github.com/pip-services3-go/pip-services3-rpc-go/build"
	cswagger "github.com/pip-services3-gox/pip-services3-swagger-gox/build"
	ffactory "test-task-pip.service/facade_service/microservice/build"
)

type FacadeProcess struct {
	*cproc.ProcessContainer
}

func NewFacadeProcess() *FacadeProcess {
	c := FacadeProcess{}
	c.ProcessContainer = cproc.NewProcessContainer("facade", "Public facade for pip-vault 1.0")
	c.AddFactory(rpcbuild.NewDefaultRpcFactory())
	c.AddFactory(ffactory.NewFacadeFactory())
	c.AddFactory(cswagger.NewDefaultSwaggerFactory())

	return &c
}
