package clients1_test

import (
	"context"
	"os"
	"testing"

	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	clients1 "test-task-pip.service/keystore_service/client/clients/version1"
	"test-task-pip.service/keystore_service/microservice/logic"
	"test-task-pip.service/keystore_service/microservice/persistence"
)

type keyDirectClientV1Test struct {
	persistence *persistence.KeySqlitePersistence
	controller  *logic.KeyController
	client      *clients1.KeyDirectClientV1
	fixture     *KeyClientV1Fixture
	ctx         context.Context
}

func newBeaconsDirectClientV1Test() *keyDirectClientV1Test {
	sqliteDatabase := os.Getenv("SQLITE_DB")
	if sqliteDatabase == "" {
		sqliteDatabase = "D:/go_path/src/test-task-pip.service/keystore_service/microservice/temp/storage.db"
	}

	if sqliteDatabase == "" {
		panic("Connection params losse")
	}

	dbConfig := cconf.NewConfigParamsFromTuples(
		"connection.database", sqliteDatabase,
	)

	ctx := context.Background()
	persistence := persistence.NewKeysSqlitePersistence()
	persistence.Configure(ctx, dbConfig)

	controller := logic.NewKeyController()

	client := clients1.NewKeyDirectClientV1()
	client.Configure(ctx, cconf.NewEmptyConfigParams())

	references := cref.NewReferencesFromTuples(ctx,
		cref.NewDescriptor("key", "persistence", "memory", "default", "1.0"), persistence,
		cref.NewDescriptor("key", "controller", "default", "default", "1.0"), controller,
		cref.NewDescriptor("key", "client", "direct", "default", "1.0"), client,
	)
	controller.SetReferences(ctx, references)
	client.SetReferences(ctx, references)

	fixture := NewKeyClientV1Fixture(client)

	return &keyDirectClientV1Test{
		persistence: persistence,
		controller:  controller,
		client:      client,
		fixture:     fixture,
		ctx:         ctx,
	}
}

func (c *keyDirectClientV1Test) setup(t *testing.T) {
	err := c.persistence.Open(c.ctx, "")
	if err != nil {
		t.Error("Failed to open persistence", err)
	}

	err = c.client.Open(c.ctx, "")
	if err != nil {
		t.Error("Failed to open client", err)
	}

	err = c.persistence.Clear(c.ctx, "")
	if err != nil {
		t.Error("Failed to clear persistence", err)
	}
}

func (c *keyDirectClientV1Test) teardown(t *testing.T) {
	err := c.client.Close(c.ctx, "")
	if err != nil {
		t.Error("Failed to close client", err)
	}

	err = c.persistence.Close(c.ctx, "")
	if err != nil {
		t.Error("Failed to close persistence", err)
	}
}

func TestBeaconsDirectClientV1(t *testing.T) {
	c := newBeaconsDirectClientV1Test()

	c.setup(t)
	t.Run("CRUD Operations", c.fixture.TestCrudOperations)
	c.teardown(t)

	c.setup(t)
	t.Run("WithFilters", c.fixture.TestGetWithFilters)
	c.teardown(t)
}
