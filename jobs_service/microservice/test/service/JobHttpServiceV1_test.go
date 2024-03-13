package service_test

import (
	"context"
	"os"
	"testing"

	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	tclients "github.com/pip-services3-gox/pip-services3-rpc-gox/test"
	"github.com/stretchr/testify/assert"
	data1 "test-task-pip.service/jobs_service/microservice/data/version1"
	"test-task-pip.service/jobs_service/microservice/logic"
	"test-task-pip.service/jobs_service/microservice/persistence"
	service1 "test-task-pip.service/jobs_service/microservice/service/version1"
)

type jobHttpServiceV1Test struct {
	JOB1        *data1.JobV1
	JOB2        *data1.JobV1
	persistence *persistence.JobSqlitePersistence
	controller  *logic.JobController
	service     *service1.JobHttpServiceV1
	client      *tclients.TestCommandableHttpClient
}

func newHttpServiceV1Test() *jobHttpServiceV1Test {
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

	restConfig := cconf.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.port", "3001",
		"connection.host", "localhost",
	)

	persistence := persistence.NewJobSqlitePersistence()
	persistence.Configure(context.Background(), dbConfig)

	controller := logic.NewJobController()

	service := service1.NewJobHttpServiceV1()
	service.Configure(context.Background(), restConfig)

	client := tclients.NewTestCommandableHttpClient("v1/jobs")
	client.Configure(context.Background(), restConfig)

	references := cref.NewReferencesFromTuples(
		context.Background(),
		cref.NewDescriptor("job", "persistence", "sqlite", "default", "1.0"), persistence,
		cref.NewDescriptor("job", "controller", "default", "default", "1.0"), controller,
		cref.NewDescriptor("job", "service", "http", "default", "1.0"), service,
		cref.NewDescriptor("job", "client", "http", "default", "1.0"), client,
	)

	controller.SetReferences(context.Background(), references)
	service.SetReferences(context.Background(), references)

	return &jobHttpServiceV1Test{
		JOB1:        JOB1,
		JOB2:        JOB2,
		persistence: persistence,
		controller:  controller,
		service:     service,
		client:      client,
	}
}

func (c *jobHttpServiceV1Test) setup(t *testing.T) {
	err := c.persistence.Open(context.Background(), "")
	if err != nil {
		t.Error("Failed to open persistence", err)
	}

	err = c.service.Open(context.Background(), "")
	if err != nil {
		t.Error("Failed to open service", err)
	}

	err = c.client.Open(context.Background(), "")
	if err != nil {
		t.Error("Failed to open client", err)
	}

	err = c.persistence.Clear(context.Background(), "")
	if err != nil {
		t.Error("Failed to clear persistence", err)
	}
}

func (c *jobHttpServiceV1Test) teardown(t *testing.T) {
	err := c.client.Close(context.Background(), "")
	if err != nil {
		t.Error("Failed to close client", err)
	}

	err = c.service.Close(context.Background(), "")
	if err != nil {
		t.Error("Failed to close service", err)
	}

	err = c.persistence.Close(context.Background(), "")
	if err != nil {
		t.Error("Failed to close persistence", err)
	}
}

func (c *jobHttpServiceV1Test) testCrudOperations(t *testing.T) {
	params := cdata.NewAnyValueMapFromTuples(
		"job_id", "4",
		"owner", "Robert",
	)

	response, err := c.client.CallCommand(context.Background(), "create_job", "", params)
	assert.Nil(t, err)
	assert.NotNil(t, response)
}

func TestJovCommmandableHttpServiceV1(t *testing.T) {
	c := newHttpServiceV1Test()

	c.setup(t)
	t.Run("CRUD Operations", c.testCrudOperations)
	c.teardown(t)
}
