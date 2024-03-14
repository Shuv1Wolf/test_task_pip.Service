package logic_test

import (
	"context"
	"testing"

	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	"github.com/stretchr/testify/assert"
	data1 "test-task-pip.service/worker_service/microservice/data/version1"
	"test-task-pip.service/worker_service/microservice/logic"
	"test-task-pip.service/worker_service/microservice/persistence"
)

type WorkerControllerTest struct {
	persistence *persistence.WorkerPersistence
	controller  *logic.WorkerController
}

func newWorkerControllerTest() *WorkerControllerTest {
	persistence := persistence.NewWorkerPersistence()

	controller := logic.NewWorkerController()

	references := cref.NewReferencesFromTuples(
		context.Background(),
		cref.NewDescriptor("worker", "persistence", "default", "*", "1.0"), persistence,
		cref.NewDescriptor("worker", "controller", "default", "default", "1.0"), controller,
	)

	controller.SetReferences(context.Background(), references)

	return &WorkerControllerTest{
		persistence: persistence,
		controller:  controller,
	}
}

func (c *WorkerControllerTest) testWorkerOperations(t *testing.T) {
	_, err := c.controller.Start(context.Background(), "")
	assert.Nil(t, err)

	status, err := c.controller.GetStatus(context.Background(), "")
	assert.Nil(t, err)
	assert.Equal(t, status, data1.Waiting)

	alias, err := c.controller.GetWorkAlias(context.Background(), "")
	assert.Nil(t, err)
	assert.Equal(t, alias, data1.NoWork)

	status, err = c.controller.Stop(context.Background(), "")
	assert.Nil(t, err)
	assert.Equal(t, status, "worker №1 stop")
}

func TestWorkerController(t *testing.T) {
	c := newWorkerControllerTest()
	t.Run("Operations", c.testWorkerOperations)
}
