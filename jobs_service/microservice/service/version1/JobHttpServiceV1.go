package service1

import (
	"context"

	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	cservices "github.com/pip-services3-gox/pip-services3-rpc-gox/services"
)

type JobHttpServiceV1 struct {
	cservices.CommandableHttpService
}

func NewJobHttpServiceV1() *JobHttpServiceV1 {
	c := &JobHttpServiceV1{}
	c.CommandableHttpService = *cservices.InheritCommandableHttpService(c, "v1/jobs")
	c.DependencyResolver.Put(context.Background(), "controller", cref.NewDescriptor("job", "controller", "*", "*", "1.0"))
	return c
}
