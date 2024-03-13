package logic

import (
	"context"

	ccmd "github.com/pip-services3-gox/pip-services3-commons-gox/commands"
	cconv "github.com/pip-services3-gox/pip-services3-commons-gox/convert"
	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	crun "github.com/pip-services3-gox/pip-services3-commons-gox/run"
	cvalid "github.com/pip-services3-gox/pip-services3-commons-gox/validate"
	data1 "test-task-pip.service/jobs_service/microservice/data/version1"
)

type JobCommandSet struct {
	ccmd.CommandSet
	controller   IJobController
	keyConvector cconv.IJSONEngine[data1.JobV1]
}

func NewJobCommandSet(controller IJobController) *JobCommandSet {
	c := &JobCommandSet{
		CommandSet:   *ccmd.NewCommandSet(),
		controller:   controller,
		keyConvector: cconv.NewDefaultCustomTypeJsonConvertor[data1.JobV1](),
	}

	c.AddCommand(c.makeGetJobsCommand())
	c.AddCommand(c.makeCreateJobCommand())
	c.AddCommand(c.makeGetNotStartedJobCommand())
	c.AddCommand(c.makeUpdateInProgressCommand())
	c.AddCommand(c.makeUpdateInCompletedCommand())
	c.AddCommand(c.makeGetJobByStatusCommand())

	return c
}

func (c *JobCommandSet) makeGetJobsCommand() ccmd.ICommand {
	return ccmd.NewCommand(
		"get_jobs",
		cvalid.NewObjectSchema().
			WithOptionalProperty("filter", cvalid.NewFilterParamsSchema()).
			WithOptionalProperty("paging", cvalid.NewPagingParamsSchema()),
		func(ctx context.Context, correlationId string, args *crun.Parameters) (result any, err error) {
			filter := cdata.NewEmptyFilterParams()
			paging := cdata.NewEmptyPagingParams()
			if _val, ok := args.Get("filter"); ok {
				filter = cdata.NewFilterParamsFromValue(_val)
			}
			if _val, ok := args.Get("paging"); ok {
				paging = cdata.NewPagingParamsFromValue(_val)
			}
			return c.controller.GetJobs(ctx, correlationId, *filter, *paging)
		})
}

func (c *JobCommandSet) makeCreateJobCommand() ccmd.ICommand {
	return ccmd.NewCommand(
		"create_job",
		cvalid.NewObjectSchema().
			WithOptionalProperty("job_id", cconv.String).
			WithRequiredProperty("owner", cconv.String),
		func(ctx context.Context, correlationId string, args *crun.Parameters) (result any, err error) {
			return c.controller.CreateJob(ctx, correlationId, args.GetAsString("job_id"), args.GetAsString("owner"))
		})
}

func (c *JobCommandSet) makeGetJobByStatusCommand() ccmd.ICommand {
	return ccmd.NewCommand(
		"get_jobs_by_status",
		cvalid.NewObjectSchema().
			WithRequiredProperty("status", cconv.String),
		func(ctx context.Context, correlationId string, args *crun.Parameters) (result any, err error) {
			return c.controller.GetJobsByStatus(ctx, correlationId, args.GetAsString("status"))
		})
}

func (c *JobCommandSet) makeGetNotStartedJobCommand() ccmd.ICommand {
	return ccmd.NewCommand(
		"get_not_started_job",
		cvalid.NewObjectSchema(),
		func(ctx context.Context, correlationId string, args *crun.Parameters) (result any, err error) {
			return c.controller.GetNotStartedJob(ctx, correlationId)
		})
}

func (c *JobCommandSet) makeUpdateInProgressCommand() ccmd.ICommand {
	return ccmd.NewCommand(
		"update_in_progress",
		cvalid.NewObjectSchema().
			WithRequiredProperty("job_id", cconv.String).
			WithRequiredProperty("owner", cconv.String),
		func(ctx context.Context, correlationId string, args *crun.Parameters) (any, error) {
			return c.controller.UpdateInProgress(ctx, correlationId, args.GetAsString("job_id"), args.GetAsString("owner"))
		})
}

func (c *JobCommandSet) makeUpdateInCompletedCommand() ccmd.ICommand {
	return ccmd.NewCommand(
		"update_in_completed",
		cvalid.NewObjectSchema().
			WithRequiredProperty("job_id", cconv.String).
			WithRequiredProperty("owner", cconv.String),
		func(ctx context.Context, correlationId string, args *crun.Parameters) (any, error) {
			return c.controller.UpdateInCompleted(ctx, correlationId, args.GetAsString("job_id"), args.GetAsString("owner"))
		})
}
