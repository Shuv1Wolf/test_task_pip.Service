package clients1

import (
	"context"

	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	clients "github.com/pip-services3-gox/pip-services3-rpc-gox/clients"
	data1 "test-task-pip.service/jobs_service/microservice/data/version1"
	"test-task-pip.service/jobs_service/microservice/logic"
)

type JobDirectClientV1 struct {
	clients.DirectClient
	controller logic.IJobController
}

func NewJobDirectClientV1() *JobDirectClientV1 {
	c := &JobDirectClientV1{
		DirectClient: *clients.NewDirectClient(),
	}
	c.DependencyResolver.Put(context.Background(), "controller", cref.NewDescriptor("job", "controller", "*", "*", "1.0"))
	return c
}

func (c *JobDirectClientV1) SetReferences(ctx context.Context, references cref.IReferences) {
	c.DirectClient.SetReferences(ctx, references)

	controller, ok := c.Controller.(logic.IJobController)
	if !ok {
		panic("JobDirectClientV1: Cant't resolv dependency 'controller' to IJobController")
	}
	c.controller = controller
}

func (c *JobDirectClientV1) GetJobs(ctx context.Context, correlationId string,
	filter cdata.FilterParams, paging cdata.PagingParams) (*cdata.DataPage[data1.JobV1], error) {
	timing := c.Instrument(ctx, correlationId, "jobs.get_jobs")
	result, err := c.controller.GetJobs(ctx, correlationId, filter, paging)
	timing.EndTiming(ctx, err)
	return &result, err
}

func (c *JobDirectClientV1) CreateJob(ctx context.Context, correlationId string, id string, owner string) (*data1.JobV1, error) {
	timing := c.Instrument(ctx, correlationId, "jobs.create_job")
	result, err := c.controller.CreateJob(ctx, correlationId, id, owner)
	timing.EndTiming(ctx, err)
	return &result, err
}

func (c *JobDirectClientV1) GetJobsByStatus(ctx context.Context, correlationId string, status string) (page *cdata.DataPage[data1.JobV1], err error) {
	timing := c.Instrument(ctx, correlationId, "jobs.get_jobs_by_statys")
	result, err := c.controller.GetJobsByStatus(ctx, correlationId, status)
	timing.EndTiming(ctx, err)
	return &result, err
}

func (c *JobDirectClientV1) GetNotStartedJob(ctx context.Context, correlationId string) (*data1.JobV1, error) {
	timing := c.Instrument(ctx, correlationId, "jobs.get_not_started_job")
	result, err := c.controller.GetNotStartedJob(ctx, correlationId)
	timing.EndTiming(ctx, err)
	return &result, err
}

func (c *JobDirectClientV1) UpdateInProgress(ctx context.Context, correlationId string, id string, owner string) (*data1.JobV1, error) {
	timing := c.Instrument(ctx, correlationId, "jobs.update_in_progress")
	result, err := c.controller.UpdateInProgress(ctx, correlationId, id, owner)
	timing.EndTiming(ctx, err)
	return &result, err
}

func (c *JobDirectClientV1) UpdateInCompleted(ctx context.Context, correlationId string, id string, owner string) (*data1.JobV1, error) {
	timing := c.Instrument(ctx, correlationId, "jobs.update_in_completed")
	result, err := c.controller.UpdateInCompleted(ctx, correlationId, id, owner)
	timing.EndTiming(ctx, err)
	return &result, err
}

func (c *JobDirectClientV1) UpdateInNotStarted(ctx context.Context, correlationId string, id string, owner string) (*data1.JobV1, error) {
	timing := c.Instrument(ctx, correlationId, "jobs.update_not_started")
	result, err := c.controller.UpdateInNotStarted(ctx, correlationId, id, owner)
	timing.EndTiming(ctx, err)
	return &result, err
}
