package build

import (
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	cbuild "github.com/pip-services3-go/pip-services3-components-go/build"
	services1 "test-task-pip.service/facade_service/microservice/services/version1"
)

type FacadeFactory struct {
	cbuild.Factory
	FacadeServiceV1Descriptor *cref.Descriptor
}

func NewFacadeFactory() *FacadeFactory {

	c := FacadeFactory{
		Factory: *cbuild.NewFactory(),
	}
	c.FacadeServiceV1Descriptor = cref.NewDescriptor("facade", "service", "http", "*", "1.0")
	c.RegisterType(c.FacadeServiceV1Descriptor, services1.NewFacadeServiceV1)
	return &c
}
