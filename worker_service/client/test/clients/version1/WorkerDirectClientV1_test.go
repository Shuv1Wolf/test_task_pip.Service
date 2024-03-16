package clients1_test

import (
	"context"
	"testing"

	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	clients1 "test-task-pip.service/worker_service/client/clients/version1"
	"test-task-pip.service/worker_service/microservice/logic"
	"test-task-pip.service/worker_service/microservice/persistence"
)

type workerDirectClientV1Test struct {
	persistence *persistence.WorkerPersistence
	controller  *logic.WorkerController
	client      *clients1.WorkerDirectClientV1
	fixture     *WorkerClientV1Fixture
	ctx         context.Context
}

func newWorkerDirectClientV1Test() *workerDirectClientV1Test {
	ctx := context.Background()
	persistence := persistence.NewWorkerPersistence()

	controller := logic.NewWorkerController()

	client := clients1.NewWorkerDirectClientV1()
	client.Configure(ctx, cconf.NewEmptyConfigParams())

	references := cref.NewReferencesFromTuples(ctx,
		cref.NewDescriptor("worker", "persistence", "default", "default", "1.0"), persistence,
		cref.NewDescriptor("worker", "controller", "default", "default", "1.0"), controller,
		cref.NewDescriptor("worker", "client", "direct", "default", "1.0"), client,
	)
	controller.SetReferences(ctx, references)
	client.SetReferences(ctx, references)

	fixture := NewWorkerClientV1Fixture(client)

	return &workerDirectClientV1Test{
		persistence: persistence,
		controller:  controller,
		client:      client,
		fixture:     fixture,
		ctx:         ctx,
	}
}

func (c *workerDirectClientV1Test) setup(t *testing.T) {
	err := c.client.Open(c.ctx, "")
	if err != nil {
		t.Error("Failed to open client", err)
	}
}

func (c *workerDirectClientV1Test) teardown(t *testing.T) {
	err := c.client.Close(c.ctx, "")
	if err != nil {
		t.Error("Failed to close client", err)
	}
}

func TestBeaconsDirectClientV1(t *testing.T) {
	c := newWorkerDirectClientV1Test()

	c.setup(t)
	t.Run("Test", c.fixture.TestWorker)
	c.teardown(t)
}
