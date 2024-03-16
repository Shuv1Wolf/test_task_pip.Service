package operations1

import (
	"context"
	"net/http"

	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	rpcservices "github.com/pip-services3-go/pip-services3-rpc-go/services"
	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
	clients1 "test-task-pip.service/worker_service/client/clients/version1"
)

type WorkerOperationsV1 struct {
	*rpcservices.RestOperations
	workerClient  *clients1.WorkerHttpClientV1
	correlationId string
}

func NewWorkerOperationsV1() *WorkerOperationsV1 {
	jobHttpConfig := cconf.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.port", "8083",
		"connection.host", "localhost",
	)

	workerClient := clients1.NewWorkerHttpClientV1()
	workerClient.Configure(context.Background(), jobHttpConfig)

	c := WorkerOperationsV1{
		RestOperations: rpcservices.NewRestOperations(),
		workerClient:   workerClient,
	}
	c.DependencyResolver.Put("worker", cref.NewDescriptor("worker", "client", "http", "*", "1.0"))
	c.correlationId = "worker_operations"
	return &c
}

func (c *WorkerOperationsV1) GetWorkAlias(res http.ResponseWriter, req *http.Request) {
	c.workerClient.Open(context.Background(), c.correlationId)
	defer c.workerClient.Close(context.Background(), c.correlationId)

	status, err := c.workerClient.GetWorkAlias(context.Background(), c.correlationId)
	if err != nil {
		c.SendError(res, req, err)
	} else {
		c.SendResult(res, req, status, nil)
	}
}

func (c *WorkerOperationsV1) GetStatus(res http.ResponseWriter, req *http.Request) {
	c.workerClient.Open(context.Background(), c.correlationId)
	defer c.workerClient.Close(context.Background(), c.correlationId)

	status, err := c.workerClient.GetStatus(context.Background(), c.correlationId)
	if err != nil {
		c.SendError(res, req, err)
	} else {
		c.SendResult(res, req, status, nil)
	}
}

func (c *WorkerOperationsV1) Start(res http.ResponseWriter, req *http.Request) {
	c.workerClient.Open(context.Background(), c.correlationId)
	defer c.workerClient.Close(context.Background(), c.correlationId)

	status, err := c.workerClient.Start(context.Background(), c.correlationId)
	if err != nil {
		c.SendError(res, req, err)
	} else {
		c.SendResult(res, req, status, nil)
	}
}

func (c *WorkerOperationsV1) Stop(res http.ResponseWriter, req *http.Request) {
	c.workerClient.Open(context.Background(), c.correlationId)
	defer c.workerClient.Close(context.Background(), c.correlationId)

	status, err := c.workerClient.Stop(context.Background(), c.correlationId)
	if err != nil {
		c.SendError(res, req, err)
	} else {
		c.SendResult(res, req, status, nil)
	}
}
