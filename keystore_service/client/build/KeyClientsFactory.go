package build

import (
	bclients "github.com/pip-services-samples/client-beacons-gox/clients/version1"
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	cbuild "github.com/pip-services3-gox/pip-services3-components-gox/build"
)

type KeyClientFactory struct {
	cbuild.Factory
	NullClientDescriptor   *cref.Descriptor
	DirectClientDescriptor *cref.Descriptor
	HttpClientDescriptor   *cref.Descriptor
}

func NewKeyClientFactory() *KeyClientFactory {

	bcf := KeyClientFactory{}
	bcf.Factory = *cbuild.NewFactory()

	bcf.NullClientDescriptor = cref.NewDescriptor("key", "client", "null", "*", "1.0")
	bcf.DirectClientDescriptor = cref.NewDescriptor("key", "client", "direct", "*", "1.0")
	bcf.HttpClientDescriptor = cref.NewDescriptor("key", "client", "http", "*", "1.0")

	bcf.RegisterType(bcf.NullClientDescriptor, bclients.NewBeaconsNullClientV1)
	bcf.RegisterType(bcf.DirectClientDescriptor, bclients.NewBeaconsDirectClientV1)
	bcf.RegisterType(bcf.HttpClientDescriptor, bclients.NewBeaconsHttpClientV1)

	return &bcf
}
