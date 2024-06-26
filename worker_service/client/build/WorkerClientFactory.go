package build

import (
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	cbuild "github.com/pip-services3-gox/pip-services3-components-gox/build"
	clients1 "test-task-pip.service/worker_service/client/clients/version1"
)

type WorkerClientFactory struct {
	cbuild.Factory
	NullClientDescriptor   *cref.Descriptor
	DirectClientDescriptor *cref.Descriptor
	HttpClientDescriptor   *cref.Descriptor
}

func NewWorkerClientFactory() *WorkerClientFactory {

	bcf := WorkerClientFactory{}
	bcf.Factory = *cbuild.NewFactory()

	bcf.NullClientDescriptor = cref.NewDescriptor("worker", "client", "null", "*", "1.0")
	bcf.DirectClientDescriptor = cref.NewDescriptor("worker", "client", "direct", "*", "1.0")
	bcf.HttpClientDescriptor = cref.NewDescriptor("worker", "client", "http", "*", "1.0")

	bcf.RegisterType(bcf.NullClientDescriptor, clients1.NewWorkerNullClientV1)
	bcf.RegisterType(bcf.DirectClientDescriptor, clients1.NewWorkerDirectClientV1)
	bcf.RegisterType(bcf.HttpClientDescriptor, clients1.NewWorkerHttpClientV1)

	return &bcf
}
