package clients1_test

import (
	"context"
	"os"
	"testing"

	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	clients1 "test-task-pip.service/jobs_service/client/clients/version1"
	"test-task-pip.service/jobs_service/microservice/logic"
	"test-task-pip.service/jobs_service/microservice/persistence"
	service1 "test-task-pip.service/jobs_service/microservice/service/version1"
)

type jobCommandableHttpClientV1Test struct {
	persistence *persistence.JobSqlitePersistence
	controller  *logic.JobController
	service     *service1.JobHttpServiceV1
	client      *clients1.JobHttpClientV1
	fixture     *JobClientV1Fixture
	ctx         context.Context
}

func newJobHttpClientV1Test() *jobCommandableHttpClientV1Test {
	sqliteDatabase := os.Getenv("SQLITE_DB")
	if sqliteDatabase == "" {
		sqliteDatabase = "D:/go_path/src/test-task-pip.service/jobs_service/microservice/temp/storage.db"
	}

	if sqliteDatabase == "" {
		panic("Connection params losse")
	}

	dbConfig := cconf.NewConfigParamsFromTuples(
		"connection.database", sqliteDatabase,
	)

	ctx := context.Background()
	persistence := persistence.NewJobSqlitePersistence()
	persistence.Configure(ctx, dbConfig)

	controller := logic.NewJobController()

	httpConfig := cconf.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.port", "8081",
		"connection.host", "localhost",
	)

	service := service1.NewJobHttpServiceV1()
	service.Configure(ctx, httpConfig)

	client := clients1.NewJobHttpClientV1()
	client.Configure(ctx, httpConfig)

	references := cref.NewReferencesFromTuples(ctx,
		cref.NewDescriptor("job", "persistence", "memory", "default", "1.0"), persistence,
		cref.NewDescriptor("job", "controller", "default", "default", "1.0"), controller,
		cref.NewDescriptor("job", "service", "http", "default", "1.0"), service,
		cref.NewDescriptor("job", "client", "http", "default", "1.0"), client,
	)
	controller.SetReferences(ctx, references)
	service.SetReferences(ctx, references)
	client.SetReferences(ctx, references)

	fixture := NewjobClientV1Fixture(client)

	return &jobCommandableHttpClientV1Test{
		persistence: persistence,
		controller:  controller,
		service:     service,
		client:      client,
		fixture:     fixture,
		ctx:         ctx,
	}
}

func (c *jobCommandableHttpClientV1Test) setup(t *testing.T) {
	err := c.persistence.Open(c.ctx, "")
	if err != nil {
		t.Error("Failed to open persistence", err)
	}

	err = c.service.Open(c.ctx, "")
	if err != nil {
		t.Error("Failed to open service", err)
	}

	err = c.client.Open(c.ctx, "")
	if err != nil {
		t.Error("Failed to open client", err)
	}

	err = c.persistence.Clear(c.ctx, "")
	if err != nil {
		t.Error("Failed to clear persistence", err)
	}
}

func (c *jobCommandableHttpClientV1Test) teardown(t *testing.T) {
	err := c.client.Close(c.ctx, "")
	if err != nil {
		t.Error("Failed to close client", err)
	}

	err = c.service.Close(c.ctx, "")
	if err != nil {
		t.Error("Failed to close service", err)
	}

	err = c.persistence.Close(c.ctx, "")
	if err != nil {
		t.Error("Failed to close persistence", err)
	}
}

func TestJobHttpClientV1(t *testing.T) {
	c := newJobHttpClientV1Test()

	c.setup(t)
	t.Run("CRUD Operations", c.fixture.TestCrudOperations)
	c.teardown(t)

	c.setup(t)
	t.Run("WithFilters", c.fixture.TestGetWithFilters)
	c.teardown(t)
}
