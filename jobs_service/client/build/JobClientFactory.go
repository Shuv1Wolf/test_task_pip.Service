package build

import (
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	cbuild "github.com/pip-services3-gox/pip-services3-components-gox/build"
	clients1 "test-task-pip.service/jobs_service/client/clients/version1"
)

type JobClientFactory struct {
	cbuild.Factory
	NullClientDescriptor   *cref.Descriptor
	DirectClientDescriptor *cref.Descriptor
	HttpClientDescriptor   *cref.Descriptor
}

func NewJobClientFactory() *JobClientFactory {

	bcf := JobClientFactory{}
	bcf.Factory = *cbuild.NewFactory()

	bcf.NullClientDescriptor = cref.NewDescriptor("job", "client", "null", "*", "1.0")
	bcf.DirectClientDescriptor = cref.NewDescriptor("job", "client", "direct", "*", "1.0")
	bcf.HttpClientDescriptor = cref.NewDescriptor("job", "client", "http", "*", "1.0")

	bcf.RegisterType(bcf.HttpClientDescriptor, clients1.NewJobHttpClientV1)
	bcf.RegisterType(bcf.DirectClientDescriptor, clients1.NewJobDirectClientV1)
	bcf.RegisterType(bcf.NullClientDescriptor, clients1.NewJobNullClentV1)

	return &bcf
}
