package clients1_test

import (
	"context"
	"testing"

	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	clients1 "test-task-pip.service/worker_service/client/clients/version1"
	"test-task-pip.service/worker_service/microservice/logic"
	"test-task-pip.service/worker_service/microservice/persistence"
	service1 "test-task-pip.service/worker_service/microservice/service/version1"
)

type workerCommandableHttpClientV1Test struct {
	persistence *persistence.WorkerPersistence
	controller  *logic.WorkerController
	service     *service1.WorkerHttpServiceV1
	client      *clients1.WorkerHttpClientV1
	fixture     *WorkerClientV1Fixture
	ctx         context.Context
}

func newWorkerHttpClientV1Test() *workerCommandableHttpClientV1Test {
	ctx := context.Background()
	persistence := persistence.NewWorkerPersistence()

	controller := logic.NewWorkerController()

	httpConfig := cconf.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.port", "8083",
		"connection.host", "localhost",
	)

	service := service1.NewWorkerHttpServiceV1()
	service.Configure(ctx, httpConfig)

	client := clients1.NewWorkerHttpClientV1()
	client.Configure(ctx, httpConfig)

	references := cref.NewReferencesFromTuples(ctx,
		cref.NewDescriptor("worker", "persistence", "default", "default", "1.0"), persistence,
		cref.NewDescriptor("worker", "controller", "default", "default", "1.0"), controller,
		cref.NewDescriptor("worker", "service", "http", "default", "1.0"), service,
		cref.NewDescriptor("worker", "client", "http", "default", "1.0"), client,
	)
	controller.SetReferences(ctx, references)
	service.SetReferences(ctx, references)
	client.SetReferences(ctx, references)

	fixture := NewWorkerClientV1Fixture(client)

	return &workerCommandableHttpClientV1Test{
		persistence: persistence,
		controller:  controller,
		service:     service,
		client:      client,
		fixture:     fixture,
		ctx:         ctx,
	}
}

func (c *workerCommandableHttpClientV1Test) setup(t *testing.T) {
	err := c.service.Open(c.ctx, "")
	if err != nil {
		t.Error("Failed to open service", err)
	}

	err = c.client.Open(c.ctx, "")
	if err != nil {
		t.Error("Failed to open client", err)
	}
}

func (c *workerCommandableHttpClientV1Test) teardown(t *testing.T) {
	err := c.client.Close(c.ctx, "")
	if err != nil {
		t.Error("Failed to close client", err)
	}

	err = c.service.Close(c.ctx, "")
	if err != nil {
		t.Error("Failed to close service", err)
	}
}

func TestJobHttpClientV1(t *testing.T) {
	c := newWorkerHttpClientV1Test()

	c.setup(t)
	t.Run("Test", c.fixture.TestWorker)
	c.teardown(t)
}
