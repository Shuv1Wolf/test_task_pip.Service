package service1

import (
	"context"

	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	cservices "github.com/pip-services3-gox/pip-services3-rpc-gox/services"
)

type WorkerHttpServiceV1 struct {
	cservices.CommandableHttpService
}

func NewWorkerHttpServiceV1() *WorkerHttpServiceV1 {
	c := &WorkerHttpServiceV1{}
	c.CommandableHttpService = *cservices.InheritCommandableHttpService(c, "v1/workers")
	c.DependencyResolver.Put(context.Background(), "controller", cref.NewDescriptor("worker", "controller", "*", "*", "1.0"))
	return c
}
