package clients1

import (
	"context"

	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	data1 "test-task-pip.service/keystore_service/microservice/data/version1"
)

type KeyNullClientV1 struct {
}

func NewKeyNullClentV1() *KeyNullClientV1 {
	return &KeyNullClientV1{}
}
func (c *KeyNullClientV1) GetKeys(ctx context.Context, correlationId string,
	filter cdata.FilterParams, paging cdata.PagingParams) (*cdata.DataPage[data1.KeyV1], error) {
	return cdata.NewEmptyDataPage[data1.KeyV1](), nil
}

func (c *KeyNullClientV1) GetkeyById(ctx context.Context, correlationId string, id string) (*data1.KeyV1, error) {
	return nil, nil
}

func (c *KeyNullClientV1) GetKeyByOwner(ctx context.Context, correlationId string, owner string) (*data1.KeyV1, error) {
	return nil, nil
}

func (c *KeyNullClientV1) CreateKey(ctx context.Context, correlationId string, item data1.KeyV1) (*data1.KeyV1, error) {
	return nil, nil
}

func (c *KeyNullClientV1) UpdateKey(ctx context.Context, correlationId string, item data1.KeyV1) (*data1.KeyV1, error) {
	return nil, nil
}

func (c *KeyNullClientV1) DeleteKeyById(ctx context.Context, correlationId string, id string) (*data1.KeyV1, error) {
	return nil, nil
}
