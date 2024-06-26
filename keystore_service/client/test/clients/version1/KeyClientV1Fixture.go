package clients1_test

import (
	"context"
	"testing"

	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/stretchr/testify/assert"
	clients1 "test-task-pip.service/keystore_service/client/clients/version1"
	data1 "test-task-pip.service/keystore_service/microservice/data/version1"
)

type KeyClientV1Fixture struct {
	KEY1   *data1.KeyV1
	KEY2   *data1.KeyV1
	KEY3   *data1.KeyV1
	client clients1.IKeyClientV1
	ctx    context.Context
}

func NewKeyClientV1Fixture(client clients1.IKeyClientV1) *KeyClientV1Fixture {
	c := KeyClientV1Fixture{}

	c.KEY1 = &data1.KeyV1{
		Id:    "1",
		Key:   "345646yhgfbmkegmgkfbmv-5434t-45mvg",
		Owner: "Alex",
	}

	c.KEY2 = &data1.KeyV1{
		Id:    "2",
		Key:   "3fwefyhgggbg-455tmgvdfw",
		Owner: "Tom",
	}

	c.KEY3 = &data1.KeyV1{
		Id:    "3",
		Key:   "3fwfwfferf-ferfgvdw-efq",
		Owner: "Harry",
	}

	c.client = client
	c.ctx = context.Background()
	return &c
}

func (c *KeyClientV1Fixture) testCreateKeys(t *testing.T) data1.KeyV1 {
	// Create the first Key
	key, err := c.client.CreateKey(context.Background(), "", *c.KEY1)
	assert.Nil(t, err)
	assert.NotEqual(t, data1.KeyV1{}, key)
	assert.Equal(t, c.KEY1.Key, key.Key)
	assert.Equal(t, c.KEY1.Owner, key.Owner)

	// Create the second key
	key, err = c.client.CreateKey(context.Background(), "", *c.KEY2)
	assert.Nil(t, err)
	assert.NotEqual(t, data1.KeyV1{}, key)
	assert.Equal(t, c.KEY2.Key, key.Key)
	assert.Equal(t, c.KEY2.Owner, key.Owner)

	// Create the third key
	key, err = c.client.CreateKey(context.Background(), "", *c.KEY3)
	assert.Nil(t, err)
	assert.NotEqual(t, data1.KeyV1{}, key)
	assert.Equal(t, c.KEY3.Key, key.Key)
	assert.Equal(t, c.KEY3.Owner, key.Owner)

	return *key
}

func (c *KeyClientV1Fixture) TestCrudOperations(t *testing.T) {
	var key1 data1.KeyV1

	// Create items
	c.testCreateKeys(t)

	// Get all beacons
	page, err := c.client.GetKeys(context.Background(), "",
		*cdata.NewEmptyFilterParams(), *cdata.NewEmptyPagingParams())
	assert.Nil(t, err)
	assert.NotNil(t, page)
	assert.True(t, page.HasData())
	assert.Len(t, page.Data, 3)
	key1 = page.Data[0].Clone()

	// Update the key
	key1.Owner = "ABC"
	key, err := c.client.UpdateKey(context.Background(), "", key1)
	assert.Nil(t, err)
	assert.NotEqual(t, data1.KeyV1{}, key)
	assert.Equal(t, key1.Id, key.Id)
	assert.Equal(t, "ABC", key.Owner)

	// Get key by owner
	key, err = c.client.GetKeyByOwner(context.Background(), "", key1.Owner)
	assert.Nil(t, err)
	assert.NotEqual(t, data1.KeyV1{}, key)
	assert.Equal(t, key1.Id, key.Id)

	// Delete the key
	key, err = c.client.DeleteKeyById(context.Background(), "", key1.Id)
	assert.Nil(t, err)
	assert.NotEqual(t, data1.KeyV1{}, key)
	assert.Equal(t, key1.Id, key.Id)

	// Try to get deleted key
	key, err = c.client.GetkeyById(context.Background(), "", key1.Id)
	assert.Nil(t, err)
	assert.Equal(t, data1.KeyV1{}, *key)
}

func (c *KeyClientV1Fixture) TestGetWithFilters(t *testing.T) {
	// Create items
	c.testCreateKeys(t)

	filter := *cdata.NewFilterParamsFromTuples(
		"id", "1",
	)
	// Filter by id
	page, err := c.client.GetKeys(context.Background(), "",
		filter,
		*cdata.NewEmptyPagingParams())
	assert.Nil(t, err)
	assert.True(t, page.HasData())
}
