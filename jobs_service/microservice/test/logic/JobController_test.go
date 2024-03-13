package logic_test

import (
	"context"
	"os"
	"testing"

	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	"github.com/stretchr/testify/assert"
	data1 "test-task-pip.service/jobs_service/microservice/data/version1"
	"test-task-pip.service/jobs_service/microservice/logic"
	"test-task-pip.service/jobs_service/microservice/persistence"
)

type jobControllerTest struct {
	JOB1        *data1.JobV1
	JOB2        *data1.JobV1
	persistence *persistence.JobSqlitePersistence
	controller  *logic.JobController
}

func newJobControllerTest() *jobControllerTest {
	JOB1 := &data1.JobV1{
		Id:     "1",
		Status: data1.NotStarted,
		Owner:  "Piter",
	}

	JOB2 := &data1.JobV1{
		Id:     "2",
		Status: data1.NotStarted,
		Owner:  "Cat",
	}

	sqliteDatabase := os.Getenv("SQLITE_DB")
	if sqliteDatabase == "" {
		sqliteDatabase = "../../temp/storage.db"
	}

	if sqliteDatabase == "" {
		panic("Connection params losse")
	}

	dbConfig := cconf.NewConfigParamsFromTuples(
		"connection.database", sqliteDatabase,
	)

	persistence := persistence.NewJobSqlitePersistence()
	persistence.Configure(context.Background(), dbConfig)

	controller := logic.NewJobController()

	references := cref.NewReferencesFromTuples(
		context.Background(),
		cref.NewDescriptor("job", "persistence", "sqlite", "default", "1.0"), persistence,
		cref.NewDescriptor("job", "controller", "default", "default", "1.0"), controller,
	)

	controller.SetReferences(context.Background(), references)

	return &jobControllerTest{
		JOB1:        JOB1,
		JOB2:        JOB2,
		persistence: persistence,
		controller:  controller,
	}
}

func (c *jobControllerTest) setup(t *testing.T) {
	err := c.persistence.Open(context.Background(), "")
	if err != nil {
		t.Error("Failed to open persistence", err)
	}

	err = c.persistence.Clear(context.Background(), "")
	if err != nil {
		t.Error("Failed to clear persistence", err)
	}
}

func (c *jobControllerTest) teardown(t *testing.T) {
	err := c.persistence.Close(context.Background(), "")
	if err != nil {
		t.Error("Failed to close persistence", err)
	}
}

func (c *jobControllerTest) testCrudOperations(t *testing.T) {
	// Create the first job
	job, err := c.controller.CreateJob(context.Background(), "", c.JOB1.Clone().Id, c.JOB1.Clone().Owner)
	assert.Nil(t, err)
	assert.NotEqual(t, data1.JobV1{}, job)
	assert.Equal(t, c.JOB1.Status, job.Status)
	assert.Equal(t, c.JOB1.Owner, job.Owner)

	// Create the second job
	job, err = c.controller.CreateJob(context.Background(), "", c.JOB2.Clone().Id, c.JOB2.Clone().Owner)
	assert.Nil(t, err)
	assert.NotEqual(t, data1.JobV1{}, job)
	assert.Equal(t, c.JOB2.Status, job.Status)
	assert.Equal(t, c.JOB2.Owner, job.Owner)

	// Get all jobs
	page, err := c.controller.GetJobs(context.Background(), "", *cdata.NewEmptyFilterParams(), *cdata.NewEmptyPagingParams())
	assert.Nil(t, err)
	assert.NotNil(t, page)
	assert.True(t, page.HasData())
	assert.Len(t, page.Data, 2)
	job1 := page.Data[0].Clone()

	// Update the job (progress)
	job1.Status = data1.Progress
	job, err = c.controller.UpdateInProgress(context.Background(), "", job1.Id, job1.Owner)
	assert.Nil(t, err)
	assert.NotEqual(t, data1.JobV1{}, job)
	assert.Equal(t, job1.Id, job.Id)
	assert.Equal(t, data1.Progress, job.Status)

	// Update the job (completed)
	job1.Status = data1.Completed
	job, err = c.controller.UpdateInCompleted(context.Background(), "", job1.Id, job1.Owner)
	assert.Nil(t, err)
	assert.NotEqual(t, data1.JobV1{}, job)
	assert.Equal(t, job1.Id, job.Id)
	assert.Equal(t, data1.Completed, job.Status)

	// Get Not Started Job
	job, err = c.controller.GetNotStartedJob(context.Background(), "")
	assert.Nil(t, err)
	assert.NotEqual(t, data1.JobV1{}, job)
	assert.Equal(t, data1.NotStarted, job.Status)
	assert.Equal(t, "2", job.Id)

	// Get Jobs By Status
	page, err = c.controller.GetJobsByStatus(context.Background(), "", data1.Completed)
	assert.Nil(t, err)
	assert.NotNil(t, page)
	assert.True(t, page.HasData())
	assert.Len(t, page.Data, 1)
}

func TestJobController(t *testing.T) {
	c := newJobControllerTest()

	c.setup(t)
	t.Run("CRUD Operations", c.testCrudOperations)
	c.teardown(t)
}
