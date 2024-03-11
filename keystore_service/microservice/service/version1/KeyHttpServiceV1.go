package service1

import (
	"context"

	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	cservices "github.com/pip-services3-gox/pip-services3-rpc-gox/services"
)

type KeyHttpServiceV1 struct {
	cservices.CommandableHttpService
}

func NewKeyHttpServiceV1() *KeyHttpServiceV1 {
	c := &KeyHttpServiceV1{}
	c.CommandableHttpService = *cservices.InheritCommandableHttpService(c, "v1/keys")
	c.DependencyResolver.Put(context.Background(), "controller", cref.NewDescriptor("key", "controller", "*", "*", "1.0"))
	return c
}
