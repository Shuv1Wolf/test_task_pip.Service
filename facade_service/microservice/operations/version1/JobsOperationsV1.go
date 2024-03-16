package operations1

import (
	"context"
	"net/http"

	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
	rpcservices "github.com/pip-services3-gox/pip-services3-rpc-gox/services"
	clients1 "test-task-pip.service/worker_service/microservice/clients/version1"
)

type JobsOperationsV1 struct {
	*rpcservices.RestOperations
	jobsClient    *clients1.JobHttpClientV1
	correlationId string
}

func NewJobsOperationsV1() *JobsOperationsV1 {
	jobHttpConfig := cconf.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.port", "8082",
		"connection.host", "localhost",
	)

	jobsClient := clients1.NewJobHttpClientV1()
	jobsClient.Configure(context.Background(), jobHttpConfig)

	c := JobsOperationsV1{
		RestOperations: rpcservices.NewRestOperations(),
		jobsClient:     jobsClient,
	}
	c.DependencyResolver.Put(context.Background(), "job", cref.NewDescriptor("job", "client", "http", "*", "1.0"))
	c.correlationId = "job_operations"
	return &c
}

func (c *JobsOperationsV1) GetJobs(res http.ResponseWriter, req *http.Request) {
	c.jobsClient.Open(context.Background(), c.correlationId)
	defer c.jobsClient.Close(context.Background(), c.correlationId)

	var filter = c.GetFilterParams(req)
	var paging = c.GetPagingParams(req)

	page, err := c.jobsClient.GetJobs(context.Background(),
		c.correlationId, *filter, *paging)
	if err != nil {
		c.SendError(res, req, err)
	} else {
		c.SendResult(res, req, page, nil)
	}
}

func (c *JobsOperationsV1) GetNotStartedJobs(res http.ResponseWriter, req *http.Request) {
	c.jobsClient.Open(context.Background(), c.correlationId)
	defer c.jobsClient.Close(context.Background(), c.correlationId)

	key, err := c.jobsClient.GetJobsByStatus(context.Background(),
		c.correlationId, "not_started")
	if err != nil {
		c.SendError(res, req, err)
	} else {
		c.SendResult(res, req, key, nil)
	}
}

func (c *JobsOperationsV1) GetCompletedJobs(res http.ResponseWriter, req *http.Request) {
	c.jobsClient.Open(context.Background(), c.correlationId)
	defer c.jobsClient.Close(context.Background(), c.correlationId)

	key, err := c.jobsClient.GetJobsByStatus(context.Background(),
		c.correlationId, "completed")
	if err != nil {
		c.SendError(res, req, err)
	} else {
		c.SendResult(res, req, key, nil)
	}
}

func (c *JobsOperationsV1) GetProgressJobs(res http.ResponseWriter, req *http.Request) {
	c.jobsClient.Open(context.Background(), c.correlationId)
	defer c.jobsClient.Close(context.Background(), c.correlationId)

	key, err := c.jobsClient.GetJobsByStatus(context.Background(),
		c.correlationId, "progress")
	if err != nil {
		c.SendError(res, req, err)
	} else {
		c.SendResult(res, req, key, nil)
	}
}
