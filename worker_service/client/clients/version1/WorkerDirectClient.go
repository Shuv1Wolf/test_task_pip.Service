package clients1

import (
	"context"

	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	clients "github.com/pip-services3-gox/pip-services3-rpc-gox/clients"
	"test-task-pip.service/worker_service/microservice/logic"
)

type WorkerDirectClientV1 struct {
	clients.DirectClient
	controller logic.IWorkerController
}

func NewWorkerDirectClientV1() *WorkerDirectClientV1 {
	c := &WorkerDirectClientV1{
		DirectClient: *clients.NewDirectClient(),
	}
	c.DependencyResolver.Put(context.Background(), "controller", cref.NewDescriptor("worker", "controller", "*", "*", "1.0"))
	return c
}

func (c *WorkerDirectClientV1) SetReferences(ctx context.Context, references cref.IReferences) {
	c.DirectClient.SetReferences(ctx, references)

	controller, ok := c.Controller.(logic.IWorkerController)
	if !ok {
		panic("WorkerDirectClientV1: Cant't resolv dependency 'controller' to IWorkerController")
	}
	c.controller = controller
}

func (c *WorkerDirectClientV1) GetWorkAlias(ctx context.Context, correlationId string) (alias string, err error) {
	timing := c.Instrument(ctx, correlationId, "workers.get_work_alias")
	result, err := c.controller.GetWorkAlias(ctx, correlationId)
	timing.EndTiming(ctx, err)
	return result, err
}

func (c *WorkerDirectClientV1) GetStatus(ctx context.Context, correlationId string) (status string, err error) {
	timing := c.Instrument(ctx, correlationId, "workers.get_status")
	result, err := c.controller.GetStatus(ctx, correlationId)
	timing.EndTiming(ctx, err)
	return result, err
}

func (c *WorkerDirectClientV1) Start(ctx context.Context, correlationId string) (status string, err error) {
	timing := c.Instrument(ctx, correlationId, "workers.start")
	result, err := c.controller.Start(ctx, correlationId)
	timing.EndTiming(ctx, err)
	return result, err
}

func (c *WorkerDirectClientV1) Stop(ctx context.Context, correlationId string) (status string, err error) {
	timing := c.Instrument(ctx, correlationId, "workers.stop")
	result, err := c.controller.Stop(ctx, correlationId)
	timing.EndTiming(ctx, err)
	return result, err
}
