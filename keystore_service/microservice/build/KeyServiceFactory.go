package build

import (
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	cbuild "github.com/pip-services3-gox/pip-services3-components-gox/build"
	"test-task-pip.service/keystore_service/microservice/logic"
	"test-task-pip.service/keystore_service/microservice/persistence"
	service1 "test-task-pip.service/keystore_service/microservice/service/version1"
)

type KeyServiceFactory struct {
	cbuild.Factory
}

func NewKeyServiceFactory() *KeyServiceFactory {
	c := &KeyServiceFactory{
		Factory: *cbuild.NewFactory(),
	}

	sqlitePersistanceDescriptor := cref.NewDescriptor("key", "persistence", "sqlite", "*", "1.0")
	controllerDescriptor := cref.NewDescriptor("key", "controller", "default", "*", "1.0")
	httpServiceV1Descriptor := cref.NewDescriptor("key", "service", "http", "*", "1.0")

	c.RegisterType(sqlitePersistanceDescriptor, persistence.NewKeysSqlitePersistence)
	c.RegisterType(controllerDescriptor, logic.NewKeyController)
	c.RegisterType(httpServiceV1Descriptor, service1.NewKeyHttpServiceV1)

	return c
}
