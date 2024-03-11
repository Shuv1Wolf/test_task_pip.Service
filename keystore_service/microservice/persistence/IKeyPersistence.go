package persistence

import (
	"context"

	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	data1 "test_task_pip.Service/keystore_service/microservice/data/version1"
)

type IKeyPersistence interface {
	Create(ctx context.Context, correlationId string, item data1.KeyV1) (data1.KeyV1, error)

	GetPageByFilter(ctx context.Context, correlationId string, filter cdata.FilterParams, paging cdata.PagingParams) (page cdata.DataPage[data1.KeyV1], err error)

	GetOneById(ctx context.Context, correlationId string, id string) (data1.KeyV1, error)

	GetOneByOwner(ctx context.Context, correlationId string, owner string) (data1.KeyV1, error)

	Update(ctx context.Context, correlationId string, item data1.KeyV1) (data1.KeyV1, error)

	DeleteById(ctx context.Context, correlationId string, id string) (data1.KeyV1, error)
}
