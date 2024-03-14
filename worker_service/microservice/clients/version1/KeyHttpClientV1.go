package clients1

import (
	"context"

	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-rpc-gox/clients"
	cclients "github.com/pip-services3-gox/pip-services3-rpc-gox/clients"
	data1 "test-task-pip.service/keystore_service/microservice/data/version1"
)

type KeyHttpClientV1 struct {
	*cclients.CommandableHttpClient
}

func NewKeyHttpClientV1() *KeyHttpClientV1 {
	c := &KeyHttpClientV1{
		CommandableHttpClient: cclients.NewCommandableHttpClient("v1/keys"),
	}
	return c
}

func (c *KeyHttpClientV1) GetKeys(ctx context.Context, correlationId string,
	filter cdata.FilterParams, paging cdata.PagingParams) (*cdata.DataPage[data1.KeyV1], error) {

	params := cdata.NewEmptyStringValueMap()
	c.AddFilterParams(params, &filter)
	c.AddPagingParams(params, &paging)

	response, err := c.CallCommand(ctx, "get_keys", correlationId, cdata.NewAnyValueMapFromValue(params.Value()))

	if err != nil {
		return cdata.NewEmptyDataPage[data1.KeyV1](), err
	}

	return clients.HandleHttpResponse[*cdata.DataPage[data1.KeyV1]](response, correlationId)
}

func (c *KeyHttpClientV1) GetkeyById(ctx context.Context, correlationId string, id string) (*data1.KeyV1, error) {
	params := cdata.NewAnyValueMapFromTuples(
		"key_id", id,
	)

	response, err := c.CallCommand(ctx, "get_key_by_id", correlationId, params)

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*data1.KeyV1](response, correlationId)
}

func (c *KeyHttpClientV1) GetKeyByOwner(ctx context.Context, correlationId string, owner string) (*data1.KeyV1, error) {
	params := cdata.NewAnyValueMapFromTuples(
		"owner", owner,
	)

	response, err := c.CallCommand(ctx, "get_key_by_owner", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*data1.KeyV1](response, correlationId)
}

func (c *KeyHttpClientV1) CreateKey(ctx context.Context, correlationId string, item data1.KeyV1) (*data1.KeyV1, error) {
	params := cdata.NewAnyValueMapFromTuples(
		"key", item,
	)

	response, err := c.CallCommand(ctx, "create_key", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*data1.KeyV1](response, correlationId)
}

func (c *KeyHttpClientV1) UpdateKey(ctx context.Context, correlationId string, item data1.KeyV1) (*data1.KeyV1, error) {
	params := cdata.NewAnyValueMapFromTuples(
		"key", item,
	)

	response, err := c.CallCommand(ctx, "update_key", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*data1.KeyV1](response, correlationId)
}

func (c *KeyHttpClientV1) DeleteKeyById(ctx context.Context, correlationId string, id string) (*data1.KeyV1, error) {
	params := cdata.NewAnyValueMapFromTuples(
		"key_id", id,
	)

	response, err := c.CallCommand(ctx, "delete_key_by_id", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*data1.KeyV1](response, correlationId)
}
