package service_test

import (
	"context"
	"testing"

	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	cclients "github.com/pip-services3-gox/pip-services3-rpc-gox/clients"
	tclients "github.com/pip-services3-gox/pip-services3-rpc-gox/test"
	"github.com/stretchr/testify/assert"
	data1 "test-task-pip.service/worker_service/microservice/data/version1"
	"test-task-pip.service/worker_service/microservice/logic"
	"test-task-pip.service/worker_service/microservice/persistence"
	service1 "test-task-pip.service/worker_service/microservice/service/version1"
)

type workerServiceV1Test struct {
	persistence *persistence.WorkerPersistence
	controller  *logic.WorkerController
	service     *service1.WorkerHttpServiceV1
	client      *tclients.TestCommandableHttpClient
}

func newHttpServiceV1Test() *workerServiceV1Test {
	restConfig := cconf.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.port", "3001",
		"connection.host", "localhost",
	)

	persistence := persistence.NewWorkerPersistence()

	controller := logic.NewWorkerController()

	service := service1.NewWorkerHttpServiceV1()
	service.Configure(context.Background(), restConfig)

	client := tclients.NewTestCommandableHttpClient("v1/workers")
	client.Configure(context.Background(), restConfig)

	references := cref.NewReferencesFromTuples(
		context.Background(),
		cref.NewDescriptor("worker", "persistence", "default", "default", "1.0"), persistence,
		cref.NewDescriptor("worker", "controller", "default", "default", "1.0"), controller,
		cref.NewDescriptor("worker", "service", "http", "default", "1.0"), service,
		cref.NewDescriptor("worker", "client", "http", "default", "1.0"), client,
	)

	controller.SetReferences(context.Background(), references)
	service.SetReferences(context.Background(), references)

	return &workerServiceV1Test{
		persistence: persistence,
		controller:  controller,
		service:     service,
		client:      client,
	}
}

func (c *workerServiceV1Test) setup(t *testing.T) {
	err := c.service.Open(context.Background(), "")
	if err != nil {
		t.Error("Failed to open service", err)
	}

	err = c.client.Open(context.Background(), "")
	if err != nil {
		t.Error("Failed to open client", err)
	}
}

func (c *workerServiceV1Test) teardown(t *testing.T) {
	err := c.client.Close(context.Background(), "")
	if err != nil {
		t.Error("Failed to close client", err)
	}

	err = c.service.Close(context.Background(), "")
	if err != nil {
		t.Error("Failed to close service", err)
	}
}

func (c *workerServiceV1Test) testOperations(t *testing.T) {
	response, err := c.client.CallCommand(context.Background(), "start", "", cdata.NewEmptyAnyValueMap())
	assert.Nil(t, err)
	assert.NotNil(t, response)

	response, err = c.client.CallCommand(context.Background(), "get_status", "", cdata.NewEmptyAnyValueMap())
	assert.Nil(t, err)
	status, err := cclients.HandleHttpResponse[string](response, "")
	assert.Equal(t, status, data1.Waiting)

	response, err = c.client.CallCommand(context.Background(), "get_work_alias", "", cdata.NewEmptyAnyValueMap())
	assert.Nil(t, err)
	status, err = cclients.HandleHttpResponse[string](response, "")
	assert.Equal(t, status, data1.NoWork)

	response, err = c.client.CallCommand(context.Background(), "stop", "", cdata.NewEmptyAnyValueMap())
	assert.Nil(t, err)
	status, err = cclients.HandleHttpResponse[string](response, "")
	assert.Equal(t, status, "worker №1 stop")
}

func TestJovCommmandableHttpServiceV1(t *testing.T) {
	c := newHttpServiceV1Test()

	c.setup(t)
	t.Run("Operations", c.testOperations)
	c.teardown(t)
}
