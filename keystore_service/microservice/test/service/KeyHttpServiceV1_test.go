package service_test

import (
	"context"
	"os"
	"testing"

	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	cclients "github.com/pip-services3-gox/pip-services3-rpc-gox/clients"
	tclients "github.com/pip-services3-gox/pip-services3-rpc-gox/test"
	"github.com/stretchr/testify/assert"
	data1 "test_task_pip.Service/keystore_service/microservice/data/version1"
	"test_task_pip.Service/keystore_service/microservice/logic"
	"test_task_pip.Service/keystore_service/microservice/persistence"
	service1 "test_task_pip.Service/keystore_service/microservice/service/version1"
)

type keyHttpServiceV1Test struct {
	KEY1        *data1.KeyV1
	KEY2        *data1.KeyV1
	persistence *persistence.KeySqlitePersistence
	controller  *logic.KeyController
	service     *service1.KeyHttpServiceV1
	client      *tclients.TestCommandableHttpClient
}

func newKeyHttpServiceV1Test() *keyHttpServiceV1Test {
	KEY1 := &data1.KeyV1{
		Id:    "1",
		Key:   "dwefngerjg-gdf98u89u89dfg-",
		Owner: "Piter",
	}

	KEY2 := &data1.KeyV1{
		Id:    "2",
		Key:   "98sd908fgdsgf-dgfd",
		Owner: "Cat",
	}

	sqliteDatabase := os.Getenv("SQLITE_DB")
	if sqliteDatabase == "" {
		sqliteDatabase = "../../temp/storage.db"
	}

	if sqliteDatabase == "" {
		panic("Connection params losse")
	}

	dbConfig := cconf.NewConfigParamsFromTuples(
		"connection.database", sqliteDatabase,
	)

	restConfig := cconf.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.port", "3001",
		"connection.host", "localhost",
	)

	persistence := persistence.NewKeysSqlitePersistence()
	persistence.Configure(context.Background(), dbConfig)

	controller := logic.NewKeyController()

	service := service1.NewKeyHttpServiceV1()
	service.Configure(context.Background(), restConfig)

	client := tclients.NewTestCommandableHttpClient("v1/keys")
	client.Configure(context.Background(), restConfig)

	references := cref.NewReferencesFromTuples(
		context.Background(),
		cref.NewDescriptor("key", "persistence", "sqlite", "default", "1.0"), persistence,
		cref.NewDescriptor("key", "controller", "default", "default", "1.0"), controller,
		cref.NewDescriptor("key", "service", "http", "default", "1.0"), service,
		cref.NewDescriptor("key", "client", "http", "default", "1.0"), client,
	)

	controller.SetReferences(context.Background(), references)
	service.SetReferences(context.Background(), references)

	return &keyHttpServiceV1Test{
		KEY1:        KEY1,
		KEY2:        KEY2,
		persistence: persistence,
		controller:  controller,
		service:     service,
		client:      client,
	}
}

func (c *keyHttpServiceV1Test) setup(t *testing.T) {
	err := c.persistence.Open(context.Background(), "")
	if err != nil {
		t.Error("Failed to open persistence", err)
	}

	err = c.service.Open(context.Background(), "")
	if err != nil {
		t.Error("Failed to open service", err)
	}

	err = c.client.Open(context.Background(), "")
	if err != nil {
		t.Error("Failed to open client", err)
	}

	err = c.persistence.Clear(context.Background(), "")
	if err != nil {
		t.Error("Failed to clear persistence", err)
	}
}

func (c *keyHttpServiceV1Test) teardown(t *testing.T) {
	err := c.client.Close(context.Background(), "")
	if err != nil {
		t.Error("Failed to close client", err)
	}

	err = c.service.Close(context.Background(), "")
	if err != nil {
		t.Error("Failed to close service", err)
	}

	err = c.persistence.Close(context.Background(), "")
	if err != nil {
		t.Error("Failed to close persistence", err)
	}
}

func (c *keyHttpServiceV1Test) testCrudOperations(t *testing.T) {
	var key1 data1.KeyV1

	// Create the first key
	params := cdata.NewAnyValueMapFromTuples(
		"key", c.KEY1.Clone(),
	)
	response, err := c.client.CallCommand(context.Background(), "create_key", "", params)
	assert.Nil(t, err)
	assert.NotNil(t, response)

	key, err := cclients.HandleHttpResponse[data1.KeyV1](response, "")
	assert.Nil(t, err)
	assert.NotEqual(t, data1.KeyV1{}, key)
	assert.Equal(t, c.KEY1.Key, key.Key)
	assert.Equal(t, c.KEY1.Owner, key.Owner)

	// Create the second key
	params = cdata.NewAnyValueMapFromTuples(
		"key", c.KEY2.Clone(),
	)
	response, err = c.client.CallCommand(context.Background(), "create_key", "", params)
	assert.Nil(t, err)
	assert.NotNil(t, response)

	key, err = cclients.HandleHttpResponse[data1.KeyV1](response, "")
	assert.Nil(t, err)
	assert.NotEqual(t, data1.KeyV1{}, key)
	assert.Equal(t, c.KEY2.Key, key.Key)
	assert.Equal(t, c.KEY2.Owner, key.Owner)

	// Get all key
	params = cdata.NewAnyValueMapFromTuples(
		"filter", cdata.NewEmptyFilterParams(),
		"paging", cdata.NewEmptyFilterParams(),
	)
	response, err = c.client.CallCommand(context.Background(), "get_keys", "", params)
	assert.Nil(t, err)
	assert.NotNil(t, response)

	page, err := cclients.HandleHttpResponse[cdata.DataPage[data1.KeyV1]](response, "")
	assert.Nil(t, err)
	assert.True(t, page.HasData())
	assert.Len(t, page.Data, 2)
	key1 = page.Data[0].Clone()

	// Update the beacon
	key1.Owner = "ABC"
	params = cdata.NewAnyValueMapFromTuples(
		"key", key1,
	)
	response, err = c.client.CallCommand(context.Background(), "update_key", "", params)
	assert.Nil(t, err)
	assert.NotNil(t, response)

	key, err = cclients.HandleHttpResponse[data1.KeyV1](response, "")
	assert.Nil(t, err)
	assert.NotEqual(t, data1.KeyV1{}, key)
	assert.Equal(t, c.KEY1.Id, key.Id)
	assert.Equal(t, "ABC", key.Owner)

	// Get key by owner
	params = cdata.NewAnyValueMapFromTuples(
		"owner", key1.Owner,
	)
	response, err = c.client.CallCommand(context.Background(), "get_key_by_owner", "", params)
	assert.Nil(t, err)
	assert.NotNil(t, response)

	key, err = cclients.HandleHttpResponse[data1.KeyV1](response, "")
	assert.Nil(t, err)
	assert.NotEqual(t, data1.KeyV1{}, key)
	assert.Equal(t, c.KEY1.Id, key.Id)

	// Delete the key
	params = cdata.NewAnyValueMapFromTuples(
		"key_id", key1.Id,
	)
	response, err = c.client.CallCommand(context.Background(), "delete_key_by_id", "", params)
	assert.Nil(t, err)

	key, err = cclients.HandleHttpResponse[data1.KeyV1](response, "")
	assert.Nil(t, err)
	assert.NotNil(t, response)

	assert.NotEqual(t, data1.KeyV1{}, key)
	assert.Equal(t, c.KEY1.Id, key.Id)

	// Try to get deleted beacon
	params = cdata.NewAnyValueMapFromTuples(
		"key_id", key1.Id,
	)
	response, err = c.client.CallCommand(context.Background(), "get_key_by_id", "", params)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	key, err = cclients.HandleHttpResponse[data1.KeyV1](response, "")
	assert.Nil(t, err)
	assert.Equal(t, data1.KeyV1{}, key)
}

func TestKeyCommmandableHttpServiceV1(t *testing.T) {
	c := newKeyHttpServiceV1Test()

	c.setup(t)
	t.Run("CRUD Operations", c.testCrudOperations)
	c.teardown(t)
}
