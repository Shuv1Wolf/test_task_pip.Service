package persistence

import (
	"context"

	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	data1 "test-task-pip.service/jobs_service/microservice/data/version1"
)

type IJobPersistence interface {
	Create(ctx context.Context, correlationId string, item data1.JobV1) (data1.JobV1, error)

	GetPageByFilter(ctx context.Context, correlationId string, filter cdata.FilterParams, paging cdata.PagingParams) (page cdata.DataPage[data1.JobV1], err error)

	GetPageByStatus(ctx context.Context, correlationId string, status string) (page cdata.DataPage[data1.JobV1], err error)

	GetNotStartedJob(ctx context.Context, correlationId string) (data1.JobV1, error)

	Update(ctx context.Context, correlationId string, item data1.JobV1) (data1.JobV1, error)
}
