package logic

import (
	"context"

	ccmd "github.com/pip-services3-gox/pip-services3-commons-gox/commands"
	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	data1 "test-task-pip.service/keystore_service/microservice/data/version1"
	"test-task-pip.service/keystore_service/microservice/persistence"
)

type KeyController struct {
	persistence persistence.IKeyPersistence
	commandSet  *KeyCommandSet
}

func NewKeyController() *KeyController {
	c := &KeyController{}
	return c
}

func (c *KeyController) SetReferences(ctx context.Context, references cref.IReferences) {
	locator := cref.NewDescriptor("key", "persistence", "*", "*", "1.0")
	p, err := references.GetOneRequired(locator)
	if p != nil && err == nil {
		if _pers, ok := p.(persistence.IKeyPersistence); ok {
			c.persistence = _pers
			return
		}
	}
	panic(cref.NewReferenceError("key.controller.SetReferences", locator))
}

func (c *KeyController) GetCommandSet() *ccmd.CommandSet {
	if c.commandSet == nil {
		c.commandSet = NewKeyCommandSet(c)
	}
	return &c.commandSet.CommandSet
}

func (c *KeyController) CreateKey(ctx context.Context, correlationId string,
	item data1.KeyV1) (data1.KeyV1, error) {
	if item.Id == "" {
		item.Id = cdata.IdGenerator.NextLong()
	}

	return c.persistence.Create(ctx, correlationId, item)
}

func (c *KeyController) GetKeys(ctx context.Context, correlationId string,
	filter cdata.FilterParams, paging cdata.PagingParams) (cdata.DataPage[data1.KeyV1], error) {

	return c.persistence.GetPageByFilter(ctx, correlationId, filter, paging)
}

func (c *KeyController) GetkeyById(ctx context.Context, correlationId string,
	id string) (data1.KeyV1, error) {

	return c.persistence.GetOneById(ctx, correlationId, id)
}

func (c *KeyController) GetKeyByOwner(ctx context.Context, correlationId string,
	owner string) (data1.KeyV1, error) {

	return c.persistence.GetOneByOwner(ctx, correlationId, owner)
}

func (c *KeyController) UpdateKey(ctx context.Context, correlationId string,
	key data1.KeyV1) (data1.KeyV1, error) {

	return c.persistence.Update(ctx, correlationId, key)
}

func (c *KeyController) DeleteKeyById(ctx context.Context, correlationId string,
	keyId string) (data1.KeyV1, error) {

	return c.persistence.DeleteById(ctx, correlationId, keyId)
}
