package clients1_test

import (
	"context"
	"testing"

	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/stretchr/testify/assert"
	clients1 "test-task-pip.service/jobs_service/client/clients/version1"
	data1 "test-task-pip.service/jobs_service/microservice/data/version1"
)

type JobClientV1Fixture struct {
	JOB1   *data1.JobV1
	JOB2   *data1.JobV1
	JOB3   *data1.JobV1
	client clients1.IJobClientV1
	ctx    context.Context
}

func NewjobClientV1Fixture(client clients1.IJobClientV1) *JobClientV1Fixture {
	c := JobClientV1Fixture{}

	c.JOB1 = &data1.JobV1{
		Id:     "5",
		Owner:  "Alex",
		Status: data1.NotStarted,
	}

	c.JOB2 = &data1.JobV1{
		Id:     "6",
		Owner:  "Tom",
		Status: data1.NotStarted,
	}

	c.JOB3 = &data1.JobV1{
		Id:     "7",
		Owner:  "Harry",
		Status: data1.NotStarted,
	}

	c.client = client
	c.ctx = context.Background()
	return &c
}

func (c *JobClientV1Fixture) testCreateJobs(t *testing.T) data1.JobV1 {
	// Create the first job
	job, err := c.client.CreateJob(context.Background(), "", c.JOB1.Id, c.JOB1.Owner)
	assert.Nil(t, err)
	assert.NotEqual(t, data1.JobV1{}, job)
	assert.Equal(t, c.JOB1.Status, job.Status)
	assert.Equal(t, c.JOB1.Owner, job.Owner)

	// Create the second job
	job, err = c.client.CreateJob(context.Background(), "", c.JOB2.Id, c.JOB2.Owner)
	assert.Nil(t, err)
	assert.NotEqual(t, data1.JobV1{}, job)
	assert.Equal(t, c.JOB2.Status, job.Status)
	assert.Equal(t, c.JOB2.Owner, job.Owner)

	// Create the third job
	job, err = c.client.CreateJob(context.Background(), "", c.JOB3.Id, c.JOB3.Owner)
	assert.Nil(t, err)
	assert.NotEqual(t, data1.JobV1{}, job)
	assert.Equal(t, c.JOB3.Status, job.Status)
	assert.Equal(t, c.JOB3.Owner, job.Owner)

	return *job
}

func (c *JobClientV1Fixture) TestCrudOperations(t *testing.T) {
	var job1 data1.JobV1

	// Create items
	c.testCreateJobs(t)

	page, err := c.client.GetJobs(context.Background(), "", *cdata.NewEmptyFilterParams(), *cdata.NewEmptyPagingParams())
	assert.Nil(t, err)
	assert.NotNil(t, page)
	assert.True(t, page.HasData())
	assert.Len(t, page.Data, 3)
	job1 = page.Data[0].Clone()

	// Update the job (progress)
	job1.Status = data1.Progress
	job, err := c.client.UpdateInProgress(context.Background(), "", job1.Id, job1.Owner)
	assert.Nil(t, err)
	assert.NotEqual(t, data1.JobV1{}, job)
	assert.Equal(t, job1.Id, job.Id)
	assert.Equal(t, data1.Progress, job.Status)

	// Update the job (completed)
	job1.Status = data1.Completed
	job, err = c.client.UpdateInCompleted(context.Background(), "", job1.Id, job1.Owner)
	assert.Nil(t, err)
	assert.NotEqual(t, data1.JobV1{}, job)
	assert.Equal(t, job1.Id, job.Id)
	assert.Equal(t, data1.Completed, job.Status)

	// Get Not Started Job
	job, err = c.client.GetNotStartedJob(context.Background(), "")
	assert.Nil(t, err)
	assert.NotEqual(t, data1.JobV1{}, job)
	assert.Equal(t, data1.NotStarted, job.Status)
	assert.Equal(t, "6", job.Id)

	// Get Jobs By Status
	page, err = c.client.GetJobsByStatus(context.Background(), "", data1.Completed)
	assert.Nil(t, err)
	assert.NotNil(t, page)
	assert.True(t, page.HasData())
	assert.Len(t, page.Data, 1)
}

func (c *JobClientV1Fixture) TestGetWithFilters(t *testing.T) {
	// Create items
	c.testCreateJobs(t)

	filter := *cdata.NewFilterParamsFromTuples(
		"id", "1",
	)
	// Filter by id
	page, err := c.client.GetJobs(context.Background(), "",
		filter,
		*cdata.NewEmptyPagingParams())
	assert.Nil(t, err)
	assert.True(t, page.HasData())
}
