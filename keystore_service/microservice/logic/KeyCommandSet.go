package logic

import (
	"context"

	ccmd "github.com/pip-services3-gox/pip-services3-commons-gox/commands"
	cconv "github.com/pip-services3-gox/pip-services3-commons-gox/convert"
	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	crun "github.com/pip-services3-gox/pip-services3-commons-gox/run"
	cvalid "github.com/pip-services3-gox/pip-services3-commons-gox/validate"
	data1 "test_task_pip.Service/keystore_service/microservice/data/version1"
)

type KeyCommandSet struct {
	ccmd.CommandSet
	controller   IKeyController
	keyConvector cconv.IJSONEngine[data1.KeyV1]
}

func NewKeyCommandSet(controller IKeyController) *KeyCommandSet {
	c := &KeyCommandSet{
		CommandSet:   *ccmd.NewCommandSet(),
		controller:   controller,
		keyConvector: cconv.NewDefaultCustomTypeJsonConvertor[data1.KeyV1](),
	}

	c.AddCommand(c.makeGetKeysCommand())
	c.AddCommand(c.makeGetKeyByIdCommand())
	c.AddCommand(c.makeGetKeyByOwnerCommand())
	c.AddCommand(c.makeCreateKeyCommand())
	c.AddCommand(c.makeDeleteKeyByIdCommand())
	c.AddCommand(c.makeUpdateKeyCommand())

	return c
}

func (c *KeyCommandSet) makeGetKeysCommand() ccmd.ICommand {
	return ccmd.NewCommand(
		"get_keys",
		cvalid.NewObjectSchema().
			WithOptionalProperty("filter", cvalid.NewFilterParamsSchema()).
			WithOptionalProperty("paging", cvalid.NewPagingParamsSchema()),
		func(ctx context.Context, correlationId string, args *crun.Parameters) (result any, err error) {
			filter := cdata.NewEmptyFilterParams()
			paging := cdata.NewEmptyPagingParams()
			if _val, ok := args.Get("filter"); ok {
				filter = cdata.NewFilterParamsFromValue(_val)
			}
			if _val, ok := args.Get("paging"); ok {
				paging = cdata.NewPagingParamsFromValue(_val)
			}
			return c.controller.GetKeys(ctx, correlationId, *filter, *paging)
		})
}

func (c *KeyCommandSet) makeGetKeyByIdCommand() ccmd.ICommand {
	return ccmd.NewCommand(
		"get_key_by_id",
		cvalid.NewObjectSchema().
			WithRequiredProperty("key_id", cconv.String),
		func(ctx context.Context, correlationId string, args *crun.Parameters) (result any, err error) {
			return c.controller.GetkeyById(ctx, correlationId, args.GetAsString("key_id"))
		})
}

func (c *KeyCommandSet) makeGetKeyByOwnerCommand() ccmd.ICommand {
	return ccmd.NewCommand(
		"get_key_by_owner",
		cvalid.NewObjectSchema().
			WithRequiredProperty("owner", cconv.String),
		func(ctx context.Context, correlationId string, args *crun.Parameters) (result any, err error) {
			return c.controller.GetKeyByOwner(ctx, correlationId, args.GetAsString("owner"))
		})
}

func (c *KeyCommandSet) makeCreateKeyCommand() ccmd.ICommand {
	return ccmd.NewCommand(
		"create_key",
		cvalid.NewObjectSchema().
			WithRequiredProperty("key", data1.NewKeyV1schema()),
		func(ctx context.Context, correlationId string, args *crun.Parameters) (result any, err error) {

			var key data1.KeyV1
			if _key, ok := args.GetAsObject("key"); ok {
				buf, err := cconv.JsonConverter.ToJson(_key)
				if err != nil {
					return nil, err
				}
				key, err = c.keyConvector.FromJson(buf)
				if err != nil {
					return nil, err
				}
			}
			return c.controller.CreateKey(ctx, correlationId, key)
		})
}

func (c *KeyCommandSet) makeUpdateKeyCommand() ccmd.ICommand {
	return ccmd.NewCommand(
		"update_key",
		cvalid.NewObjectSchema().
			WithRequiredProperty("key", data1.NewKeyV1schema()),
		func(ctx context.Context, correlationId string, args *crun.Parameters) (result any, err error) {
			var key data1.KeyV1
			if _key, ok := args.GetAsObject("key"); ok {
				buf, err := cconv.JsonConverter.ToJson(_key)
				if err != nil {
					return nil, err
				}
				key, err = c.keyConvector.FromJson(buf)
				if err != nil {
					return nil, err
				}
			}
			return c.controller.UpdateKey(ctx, correlationId, key)
		})
}

func (c *KeyCommandSet) makeDeleteKeyByIdCommand() ccmd.ICommand {
	return ccmd.NewCommand(
		"delete_key_by_id",
		cvalid.NewObjectSchema().
			WithRequiredProperty("key_id", cconv.String),
		func(ctx context.Context, correlationId string, args *crun.Parameters) (result any, err error) {
			return c.controller.DeleteKeyById(ctx, correlationId, args.GetAsString("key_id"))
		})
}
