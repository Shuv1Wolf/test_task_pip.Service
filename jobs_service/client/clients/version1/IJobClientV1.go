package clients1

import (
	"context"

	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	data1 "test-task-pip.service/jobs_service/microservice/data/version1"
)

type IJobClientV1 interface {
	CreateJob(ctx context.Context, correlationId string, id string, owner string) (*data1.JobV1, error)

	GetJobs(ctx context.Context, correlationId string, filter cdata.FilterParams, paging cdata.PagingParams) (page *cdata.DataPage[data1.JobV1], err error)

	GetJobsByStatus(ctx context.Context, correlationId string, status string) (page *cdata.DataPage[data1.JobV1], err error)

	GetNotStartedJob(ctx context.Context, correlationId string) (*data1.JobV1, error)

	UpdateInProgress(ctx context.Context, correlationId string, id string, owner string) (*data1.JobV1, error)
	UpdateInCompleted(ctx context.Context, correlationId string, id string, owner string) (*data1.JobV1, error)
}
