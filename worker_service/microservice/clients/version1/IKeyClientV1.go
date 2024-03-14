package clients1

import (
	"context"

	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	data1 "test-task-pip.service/keystore_service/microservice/data/version1"
)

type IKeyClientV1 interface {
	CreateKey(ctx context.Context, correlationId string, item data1.KeyV1) (*data1.KeyV1, error)

	GetKeys(ctx context.Context, correlationId string, filter cdata.FilterParams, paging cdata.PagingParams) (page *cdata.DataPage[data1.KeyV1], err error)

	GetkeyById(ctx context.Context, correlationId string, id string) (*data1.KeyV1, error)

	GetKeyByOwner(ctx context.Context, correlationId string, owner string) (*data1.KeyV1, error)

	UpdateKey(ctx context.Context, correlationId string, item data1.KeyV1) (*data1.KeyV1, error)

	DeleteKeyById(ctx context.Context, correlationId string, id string) (*data1.KeyV1, error)
}
