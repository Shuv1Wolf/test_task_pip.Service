package clients1

import (
	"context"

	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-rpc-gox/clients"
	cclients "github.com/pip-services3-gox/pip-services3-rpc-gox/clients"
	data1 "test-task-pip.service/jobs_service/microservice/data/version1"
)

type JobHttpClientV1 struct {
	*cclients.CommandableHttpClient
}

func NewJobHttpClientV1() *JobHttpClientV1 {
	c := &JobHttpClientV1{
		CommandableHttpClient: cclients.NewCommandableHttpClient("v1/jobs"),
	}
	return c
}

func (c *JobHttpClientV1) CreateJob(ctx context.Context, correlationId string,
	id string, owner string) (*data1.JobV1, error) {

	params := cdata.NewAnyValueMapFromTuples(
		"job_id", id,
		"owner", owner,
	)

	response, err := c.CallCommand(ctx, "create_job", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*data1.JobV1](response, correlationId)
}

func (c *JobHttpClientV1) GetJobs(ctx context.Context, correlationId string,
	filter cdata.FilterParams, paging cdata.PagingParams) (page *cdata.DataPage[data1.JobV1], err error) {

	params := cdata.NewEmptyStringValueMap()
	c.AddFilterParams(params, &filter)
	c.AddPagingParams(params, &paging)

	response, err := c.CallCommand(ctx, "get_jobs", correlationId, cdata.NewAnyValueMapFromValue(params.Value()))

	if err != nil {
		return cdata.NewEmptyDataPage[data1.JobV1](), err
	}

	return clients.HandleHttpResponse[*cdata.DataPage[data1.JobV1]](response, correlationId)
}

func (c *JobHttpClientV1) GetJobsByStatus(ctx context.Context, correlationId string,
	status string) (page *cdata.DataPage[data1.JobV1], err error) {

	params := cdata.NewAnyValueMapFromTuples(
		"status", status,
	)

	response, err := c.CallCommand(ctx, "get_jobs_by_status", correlationId, cdata.NewAnyValueMapFromValue(params.Value()))

	if err != nil {
		return cdata.NewEmptyDataPage[data1.JobV1](), err
	}

	return clients.HandleHttpResponse[*cdata.DataPage[data1.JobV1]](response, correlationId)
}

// TODO: проверить
func (c *JobHttpClientV1) GetNotStartedJob(ctx context.Context, correlationId string) (*data1.JobV1, error) {
	response, err := c.CallCommand(ctx, "get_not_started_job", correlationId, cdata.NewEmptyAnyValueMap())
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*data1.JobV1](response, correlationId)
}

func (c *JobHttpClientV1) UpdateInProgress(ctx context.Context, correlationId string,
	id string, owner string) (*data1.JobV1, error) {

	params := cdata.NewAnyValueMapFromTuples(
		"job_id", id,
		"owner", owner,
	)

	response, err := c.CallCommand(ctx, "update_in_progress", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*data1.JobV1](response, correlationId)
}

func (c *JobHttpClientV1) UpdateInCompleted(ctx context.Context, correlationId string,
	id string, owner string) (*data1.JobV1, error) {

	params := cdata.NewAnyValueMapFromTuples(
		"job_id", id,
		"owner", owner,
	)

	response, err := c.CallCommand(ctx, "update_in_completed", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*data1.JobV1](response, correlationId)
}
