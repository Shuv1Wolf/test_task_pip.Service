package logic

import (
	"context"

	ccmd "github.com/pip-services3-gox/pip-services3-commons-gox/commands"
	cconv "github.com/pip-services3-gox/pip-services3-commons-gox/convert"
	crun "github.com/pip-services3-gox/pip-services3-commons-gox/run"
	cvalid "github.com/pip-services3-gox/pip-services3-commons-gox/validate"
	data1 "test-task-pip.service/worker_service/microservice/data/version1"
)

type WorkerCommandSet struct {
	ccmd.CommandSet
	controller      IWorkerController
	workerConvector cconv.IJSONEngine[data1.Worker]
}

func NewWorkerCommandSet(controller IWorkerController) *WorkerCommandSet {
	c := &WorkerCommandSet{
		CommandSet:      *ccmd.NewCommandSet(),
		controller:      controller,
		workerConvector: cconv.NewDefaultCustomTypeJsonConvertor[data1.Worker](),
	}

	c.AddCommand(c.makeGetStatusCommand())
	c.AddCommand(c.makeGetWorkAliasCommand())
	c.AddCommand(c.makeStartCommand())
	c.AddCommand(c.makeStopCommand())

	return c
}

func (c *WorkerCommandSet) makeGetStatusCommand() ccmd.ICommand {
	return ccmd.NewCommand(
		"get_status",
		cvalid.NewObjectSchema(),
		func(ctx context.Context, correlationId string, args *crun.Parameters) (result any, err error) {
			return c.controller.GetStatus(ctx, correlationId)
		})
}

func (c *WorkerCommandSet) makeGetWorkAliasCommand() ccmd.ICommand {
	return ccmd.NewCommand(
		"get_work_alias",
		cvalid.NewObjectSchema(),
		func(ctx context.Context, correlationId string, args *crun.Parameters) (result any, err error) {
			return c.controller.GetWorkAlias(ctx, correlationId)
		})
}

func (c *WorkerCommandSet) makeStartCommand() ccmd.ICommand {
	return ccmd.NewCommand(
		"start",
		cvalid.NewObjectSchema(),
		func(ctx context.Context, correlationId string, args *crun.Parameters) (result any, err error) {
			return c.controller.Start(ctx, correlationId)
		})
}

func (c *WorkerCommandSet) makeStopCommand() ccmd.ICommand {
	return ccmd.NewCommand(
		"stop",
		cvalid.NewObjectSchema(),
		func(ctx context.Context, correlationId string, args *crun.Parameters) (result any, err error) {
			return c.controller.Stop(ctx, correlationId)
		})
}
