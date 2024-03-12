package clients1

import (
	"context"

	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	clients "github.com/pip-services3-gox/pip-services3-rpc-gox/clients"
	data1 "test-task-pip.service/keystore_service/microservice/data/version1"
	logic "test-task-pip.service/keystore_service/microservice/logic"
)

type KeyDirectClientV1 struct {
	clients.DirectClient
	controller logic.IKeyController
}

func NewKeyDirectClientV1() *KeyDirectClientV1 {
	c := &KeyDirectClientV1{
		DirectClient: *clients.NewDirectClient(),
	}
	c.DependencyResolver.Put(context.Background(), "controller", cref.NewDescriptor("key", "controller", "*", "*", "1.0"))
	return c
}

func (c *KeyDirectClientV1) SetReferences(ctx context.Context, references cref.IReferences) {
	c.DirectClient.SetReferences(ctx, references)

	controller, ok := c.Controller.(logic.IKeyController)
	if !ok {
		panic("KeyDirectClientV1: Cant't resolv dependency 'controller' to IBeaconsController")
	}
	c.controller = controller
}

func (c *KeyDirectClientV1) GetKeys(ctx context.Context, correlationId string,
	filter cdata.FilterParams, paging cdata.PagingParams) (*cdata.DataPage[data1.KeyV1], error) {
	timing := c.Instrument(ctx, correlationId, "keys.get_keys")
	result, err := c.controller.GetKeys(ctx, correlationId, filter, paging)
	timing.EndTiming(ctx, err)
	return &result, err
}

func (c *KeyDirectClientV1) GetkeyById(ctx context.Context, correlationId string, id string) (*data1.KeyV1, error) {
	timing := c.Instrument(ctx, correlationId, "keys.get_key_by_id")
	result, err := c.controller.GetkeyById(ctx, correlationId, id)
	timing.EndTiming(ctx, err)
	return &result, err
}

func (c *KeyDirectClientV1) GetKeyByOwner(ctx context.Context, correlationId string, owner string) (*data1.KeyV1, error) {
	timing := c.Instrument(ctx, correlationId, "keys.get_key_by_owner")
	result, err := c.controller.GetKeyByOwner(ctx, correlationId, owner)
	timing.EndTiming(ctx, err)
	return &result, err
}

func (c *KeyDirectClientV1) CreateKey(ctx context.Context, correlationId string, item data1.KeyV1) (*data1.KeyV1, error) {
	timing := c.Instrument(ctx, correlationId, "keys.create_key")
	result, err := c.controller.CreateKey(ctx, correlationId, item)
	timing.EndTiming(ctx, err)
	return &result, err
}

func (c *KeyDirectClientV1) UpdateKey(ctx context.Context, correlationId string, item data1.KeyV1) (*data1.KeyV1, error) {
	timing := c.Instrument(ctx, correlationId, "keys.update_key")
	result, err := c.controller.UpdateKey(ctx, correlationId, item)
	timing.EndTiming(ctx, err)
	return &result, err
}

func (c *KeyDirectClientV1) DeleteKeyById(ctx context.Context, correlationId string, id string) (*data1.KeyV1, error) {
	timing := c.Instrument(ctx, correlationId, "keys.delete_key_by_id")
	result, err := c.controller.DeleteKeyById(ctx, correlationId, id)
	timing.EndTiming(ctx, err)
	return &result, err
}
