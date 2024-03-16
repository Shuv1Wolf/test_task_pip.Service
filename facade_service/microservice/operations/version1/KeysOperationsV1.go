package operations1

import (
	"context"
	"net/http"

	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
	rpcservices "github.com/pip-services3-gox/pip-services3-rpc-gox/services"
	data1 "test-task-pip.service/keystore_service/microservice/data/version1"
	clients1 "test-task-pip.service/worker_service/microservice/clients/version1"
)

type KeysOperationsV1 struct {
	*rpcservices.RestOperations
	keyClient     *clients1.KeyHttpClientV1
	correlationId string
}

func NewKeysOperationsV1() *KeysOperationsV1 {
	jobHttpConfig := cconf.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.port", "8081",
		"connection.host", "localhost",
	)

	keyClient := clients1.NewKeyHttpClientV1()
	keyClient.Configure(context.Background(), jobHttpConfig)

	c := KeysOperationsV1{
		RestOperations: rpcservices.NewRestOperations(),
		keyClient:      keyClient,
	}
	c.DependencyResolver.Put(context.Background(), "key", cref.NewDescriptor("key", "client", "http", "*", "1.0"))
	c.correlationId = "key_operations"
	return &c
}

func (c *KeysOperationsV1) GetKeys(res http.ResponseWriter, req *http.Request) {
	c.keyClient.Open(context.Background(), c.correlationId)
	defer c.keyClient.Close(context.Background(), c.correlationId)

	var filter = c.GetFilterParams(req)
	var paging = c.GetPagingParams(req)

	page, err := c.keyClient.GetKeys(context.Background(),
		c.correlationId, *filter, *paging)
	if err != nil {
		c.SendError(res, req, err)
	} else {
		c.SendResult(res, req, page, nil)
	}
}

func (c *KeysOperationsV1) GetkeyById(res http.ResponseWriter, req *http.Request) {
	c.keyClient.Open(context.Background(), c.correlationId)
	defer c.keyClient.Close(context.Background(), c.correlationId)

	id := c.GetParam(req, "key_id")

	key, err := c.keyClient.GetkeyById(context.Background(),
		c.correlationId, id)
	if err != nil {
		c.SendError(res, req, err)
	} else {
		c.SendResult(res, req, key, nil)
	}
}

func (c *KeysOperationsV1) GetKeyByOwner(res http.ResponseWriter, req *http.Request) {
	c.keyClient.Open(context.Background(), c.correlationId)
	defer c.keyClient.Close(context.Background(), c.correlationId)

	owner := c.GetParam(req, "owner")

	key, err := c.keyClient.GetKeyByOwner(context.Background(),
		c.correlationId, owner)
	if err != nil {
		c.SendError(res, req, err)
	} else {
		c.SendResult(res, req, key, nil)
	}
}

func (c *KeysOperationsV1) UpdateKey(res http.ResponseWriter, req *http.Request) {
	c.keyClient.Open(context.Background(), c.correlationId)
	defer c.keyClient.Close(context.Background(), c.correlationId)

	data := data1.KeyV1{}
	err := c.DecodeBody(req, &data)
	if err != nil {
		c.SendError(res, req, err)
	}

	key, err := c.keyClient.UpdateKey(context.Background(),
		c.correlationId, data)
	if err != nil {
		c.SendError(res, req, err)
	} else {
		c.SendResult(res, req, key, nil)
	}
}

func (c *KeysOperationsV1) DeleteKeyById(res http.ResponseWriter, req *http.Request) {
	c.keyClient.Open(context.Background(), c.correlationId)
	defer c.keyClient.Close(context.Background(), c.correlationId)

	c.keyClient.Open(context.Background(), c.correlationId)
	defer c.keyClient.Close(context.Background(), c.correlationId)

	id := c.GetParam(req, "key_id")

	key, err := c.keyClient.DeleteKeyById(context.Background(),
		c.correlationId, id)
	if err != nil {
		c.SendError(res, req, err)
	} else {
		c.SendResult(res, req, key, nil)
	}
}
