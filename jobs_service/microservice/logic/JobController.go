package logic

import (
	"context"

	ccmd "github.com/pip-services3-gox/pip-services3-commons-gox/commands"
	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	data1 "test-task-pip.service/jobs_service/microservice/data/version1"
	"test-task-pip.service/jobs_service/microservice/persistence"
)

type JobController struct {
	persistence persistence.IJobPersistence
	commandSet  *JobCommandSet
}

func NewJobController() *JobController {
	c := &JobController{}
	return c
}

func (c *JobController) SetReferences(ctx context.Context, references cref.IReferences) {
	locator := cref.NewDescriptor("job", "persistence", "*", "*", "1.0")
	p, err := references.GetOneRequired(locator)
	if p != nil && err == nil {
		if _pers, ok := p.(persistence.IJobPersistence); ok {
			c.persistence = _pers
			return
		}
	}
	panic(cref.NewReferenceError("job.controller.SetReferences", locator))
}

func (c *JobController) GetCommandSet() *ccmd.CommandSet {
	if c.commandSet == nil {
		c.commandSet = NewJobCommandSet(c)
	}
	return &c.commandSet.CommandSet
}

// TODO: подключить воркер
func (c *JobController) CreateJob(ctx context.Context, correlationId string,
	id string, owner string) (data1.JobV1, error) {
	if id == "" {
		id = cdata.IdGenerator.NextLong()
	}

	job := data1.JobV1{
		Id:     id,
		Owner:  owner,
		Status: data1.NotStarted,
	}
	return c.persistence.Create(ctx, correlationId, job)
}

func (c *JobController) GetJobs(ctx context.Context, correlationId string,
	filter cdata.FilterParams, paging cdata.PagingParams) (page cdata.DataPage[data1.JobV1], err error) {

	return c.persistence.GetPageByFilter(ctx, correlationId, filter, paging)
}

func (c *JobController) GetJobsByStatus(ctx context.Context, correlationId string,
	status string) (page cdata.DataPage[data1.JobV1], err error) {

	return c.persistence.GetPageByStatus(ctx, correlationId, status)
}

func (c *JobController) GetNotStartedJob(ctx context.Context, correlationId string) (data1.JobV1, error) {

	return c.persistence.GetNotStartedJob(ctx, correlationId)
}

func (c *JobController) UpdateInProgress(ctx context.Context, correlationId string,
	id string, owner string) (data1.JobV1, error) {
	job := data1.JobV1{
		Id:     id,
		Owner:  owner,
		Status: data1.Progress,
	}
	return c.persistence.Update(ctx, correlationId, job)
}

func (c *JobController) UpdateInCompleted(ctx context.Context, correlationId string,
	id string, owner string) (data1.JobV1, error) {
	job := data1.JobV1{
		Id:     id,
		Owner:  owner,
		Status: data1.Completed,
	}
	return c.persistence.Update(ctx, correlationId, job)
}
