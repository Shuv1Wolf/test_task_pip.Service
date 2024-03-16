package clients1

import (
	"context"

	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	data1 "test-task-pip.service/jobs_service/microservice/data/version1"
)

type JobNullClientV1 struct {
}

func NewKeyNullClentV1() *JobNullClientV1 {
	return &JobNullClientV1{}
}

func (c *JobNullClientV1) GetJobs(ctx context.Context, correlationId string,
	filter cdata.FilterParams, paging cdata.PagingParams) (*cdata.DataPage[data1.JobV1], error) {
	return cdata.NewEmptyDataPage[data1.JobV1](), nil
}

func (c *JobNullClientV1) CreateJob(ctx context.Context, correlationId string, id string, owner string) (*data1.JobV1, error) {
	return nil, nil
}

func (c *JobNullClientV1) GetJobsByStatus(ctx context.Context, correlationId string, status string) (page *cdata.DataPage[data1.JobV1], err error) {
	return nil, nil
}

func (c *JobNullClientV1) GetNotStartedJob(ctx context.Context, correlationId string) (*data1.JobV1, error) {
	return nil, nil
}

func (c *JobNullClientV1) UpdateInProgress(ctx context.Context, correlationId string, id string, owner string) (*data1.JobV1, error) {
	return nil, nil
}

func (c *JobNullClientV1) UpdateInCompleted(ctx context.Context, correlationId string, id string, owner string) (*data1.JobV1, error) {
	return nil, nil
}

func (c *JobNullClientV1) UpdateInNotStarted(ctx context.Context, correlationId string, id string, owner string) (*data1.JobV1, error) {
	return nil, nil
}
