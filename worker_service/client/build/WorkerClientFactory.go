package build

import (
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	cbuild "github.com/pip-services3-gox/pip-services3-components-gox/build"
	"test-task-pip.service/worker_service/client/clients"
)

type KeyClientFactory struct {
	cbuild.Factory
	// NullClientDescriptor   *cref.Descriptor
	// DirectClientDescriptor *cref.Descriptor
	HttpClientDescriptor *cref.Descriptor
}

func NewKeyClientFactory() *KeyClientFactory {

	bcf := KeyClientFactory{}
	bcf.Factory = *cbuild.NewFactory()

	// bcf.NullClientDescriptor = cref.NewDescriptor("worker", "client", "null", "*", "1.0")
	// bcf.DirectClientDescriptor = cref.NewDescriptor("worker", "client", "direct", "*", "1.0")
	bcf.HttpClientDescriptor = cref.NewDescriptor("worker", "client", "http", "*", "1.0")

	// bcf.RegisterType(bcf.NullClientDescriptor, clients1.NewKeyNullClentV1)
	// bcf.RegisterType(bcf.DirectClientDescriptor, clients1.NewKeyDirectClientV1)
	bcf.RegisterType(bcf.HttpClientDescriptor, clients.NewWorkerHttpClientV1)

	return &bcf
}
