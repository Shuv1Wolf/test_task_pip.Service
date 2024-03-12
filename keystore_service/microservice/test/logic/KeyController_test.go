package logic_test

import (
	"context"
	"os"
	"testing"

	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	"github.com/stretchr/testify/assert"
	data1 "test-task-pip.service/keystore_service/microservice/data/version1"
	"test-task-pip.service/keystore_service/microservice/logic"
	"test-task-pip.service/keystore_service/microservice/persistence"
)

type keyControllerTest struct {
	KEY1        *data1.KeyV1
	KEY2        *data1.KeyV1
	persistence *persistence.KeySqlitePersistence
	controller  *logic.KeyController
}

func newKeyControllerTest() *keyControllerTest {
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

	persistence := persistence.NewKeysSqlitePersistence()
	persistence.Configure(context.Background(), dbConfig)

	controller := logic.NewKeyController()

	references := cref.NewReferencesFromTuples(
		context.Background(),
		cref.NewDescriptor("key", "persistence", "sqlite", "default", "1.0"), persistence,
		cref.NewDescriptor("key", "controller", "default", "default", "1.0"), controller,
	)

	controller.SetReferences(context.Background(), references)

	return &keyControllerTest{
		KEY1:        KEY1,
		KEY2:        KEY2,
		persistence: persistence,
		controller:  controller,
	}
}

func (c *keyControllerTest) setup(t *testing.T) {
	err := c.persistence.Open(context.Background(), "")
	if err != nil {
		t.Error("Failed to open persistence", err)
	}

	err = c.persistence.Clear(context.Background(), "")
	if err != nil {
		t.Error("Failed to clear persistence", err)
	}
}

func (c *keyControllerTest) teardown(t *testing.T) {
	err := c.persistence.Close(context.Background(), "")
	if err != nil {
		t.Error("Failed to close persistence", err)
	}
}

func (c *keyControllerTest) testCrudOperations(t *testing.T) {
	var key1 data1.KeyV1

	// Create the first key
	key, err := c.controller.CreateKey(context.Background(), "", c.KEY1.Clone())
	assert.Nil(t, err)
	assert.NotEqual(t, data1.KeyV1{}, key)
	assert.Equal(t, c.KEY1.Key, key.Key)
	assert.Equal(t, c.KEY1.Owner, key.Owner)

	// Create the second key
	key, err = c.controller.CreateKey(context.Background(), "", c.KEY2.Clone())
	assert.Nil(t, err)
	assert.NotEqual(t, data1.KeyV1{}, key)
	assert.Equal(t, c.KEY2.Key, key.Key)
	assert.Equal(t, c.KEY2.Owner, key.Owner)

	// Get all keys
	page, err := c.controller.GetKeys(context.Background(), "", *cdata.NewEmptyFilterParams(), *cdata.NewEmptyPagingParams())
	assert.Nil(t, err)
	assert.NotNil(t, page)
	assert.True(t, page.HasData())
	assert.Len(t, page.Data, 2)
	key1 = page.Data[0].Clone()

	// Update the key
	key1.Owner = "ABC"
	key, err = c.controller.UpdateKey(context.Background(), "", key1)
	assert.Nil(t, err)
	assert.NotEqual(t, data1.KeyV1{}, key)
	assert.Equal(t, key1.Id, key.Id)
	assert.Equal(t, "ABC", key.Owner)

	// Get key by owner
	key, err = c.controller.GetKeyByOwner(context.Background(), "", key1.Owner)
	assert.Nil(t, err)
	assert.NotEqual(t, data1.KeyV1{}, key)
	assert.Equal(t, key1.Id, key.Id)

	// Delete the key
	key, err = c.controller.DeleteKeyById(context.Background(), "", key1.Id)
	assert.Nil(t, err)
	assert.NotEqual(t, data1.KeyV1{}, key)
	assert.Equal(t, key1.Id, key.Id)

	// Try to get deleted key
	key, err = c.controller.GetkeyById(context.Background(), "", key.Id)
	assert.Nil(t, err)
	assert.Equal(t, data1.KeyV1{}, key)
}

func TestKeyController(t *testing.T) {
	c := newKeyControllerTest()

	c.setup(t)
	t.Run("CRUD Operations", c.testCrudOperations)
	c.teardown(t)
}
