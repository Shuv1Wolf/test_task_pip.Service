package persistence_test

import (
	"context"
	"testing"

	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/stretchr/testify/assert"
	data1 "test-task-pip.service/jobs_service/microservice/data/version1"
	"test-task-pip.service/jobs_service/microservice/persistence"
)

type JobPersistenceFixture struct {
	JOB1        *data1.JobV1
	JOB2        *data1.JobV1
	JOB3        *data1.JobV1
	persistence persistence.IJobPersistence
}

func NewJobPersistenceFixture(persistence persistence.JobSqlitePersistence) *JobPersistenceFixture {
	c := JobPersistenceFixture{}

	c.JOB1 = &data1.JobV1{
		Id:     "1",
		Owner:  "Alex",
		Status: data1.Progress,
	}

	c.JOB2 = &data1.JobV1{
		Id:     "2",
		Owner:  "Tomac",
		Status: data1.Progress,
	}

	c.JOB3 = &data1.JobV1{
		Id:     "3",
		Owner:  "Anna",
		Status: data1.NotStarted,
	}

	c.persistence = &persistence
	return &c
}

func (c *JobPersistenceFixture) testCreatejobs(t *testing.T) data1.JobV1 {
	// Create the first job
	job, err := c.persistence.Create(context.Background(), "", *c.JOB1)
	assert.Nil(t, err)
	assert.NotEqual(t, data1.JobV1{}, job)
	assert.Equal(t, c.JOB1.Status, job.Status)
	assert.Equal(t, c.JOB1.Owner, job.Owner)

	// Create the second job
	job, err = c.persistence.Create(context.Background(), "", *c.JOB2)
	assert.Nil(t, err)
	assert.NotEqual(t, data1.JobV1{}, job)
	assert.Equal(t, c.JOB2.Status, job.Status)
	assert.Equal(t, c.JOB2.Owner, job.Owner)

	// Create the third job
	job, err = c.persistence.Create(context.Background(), "", *c.JOB3)
	assert.Nil(t, err)
	assert.NotEqual(t, data1.JobV1{}, job)
	assert.Equal(t, c.JOB3.Status, job.Status)
	assert.Equal(t, c.JOB3.Owner, job.Owner)

	return job
}

func (c *JobPersistenceFixture) TestCrudOperations(t *testing.T) {
	// Create items
	var job1 data1.JobV1

	// Create items
	c.testCreatejobs(t)

	// Get all beacons
	page, err := c.persistence.GetPageByFilter(context.Background(), "",
		*cdata.NewEmptyFilterParams(), *cdata.NewEmptyPagingParams())
	assert.Nil(t, err)
	assert.NotNil(t, page)
	assert.True(t, page.HasData())
	assert.Len(t, page.Data, 3)
	job1 = page.Data[0].Clone()

	// Update the job
	job1.Status = data1.Completed
	job, err := c.persistence.Update(context.Background(), "", job1)
	assert.Nil(t, err)
	assert.NotEqual(t, data1.JobV1{}, job)
	assert.Equal(t, job1.Id, job.Id)
	assert.Equal(t, data1.Completed, job.Status)

	// Get page job by status
	page, err = c.persistence.GetPageByStatus(context.Background(), "", job1.Status)
	assert.Nil(t, err)
	assert.NotNil(t, page)
	assert.True(t, page.HasData())
	assert.Len(t, page.Data, 1)

	// Get not_started job
	job, err = c.persistence.GetNotStartedJob(context.Background(), "")
	assert.Nil(t, err)
	assert.NotEqual(t, data1.JobV1{}, job)
	assert.Equal(t, "3", job.Id)

}

func (c *JobPersistenceFixture) TestGetWithFilters(t *testing.T) {
	// Create items
	c.testCreatejobs(t)

	filter := *cdata.NewFilterParamsFromTuples(
		"id", "1",
	)
	// Filter by id
	page, err := c.persistence.GetPageByFilter(context.Background(), "",
		filter,
		*cdata.NewEmptyPagingParams())
	assert.Nil(t, err)
	assert.True(t, page.HasData())
	assert.Len(t, page.Data, 1)
}
