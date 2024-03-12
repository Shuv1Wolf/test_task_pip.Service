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
	service1 "test-task-pip.service/keystore_service/microservice/service/version1"
)

type keyCommandableHttpClientV1Test struct {
	persistence *persistence.KeySqlitePersistence
	controller  *logic.KeyController
	service     *service1.KeyHttpServiceV1
	client      *clients1.KeyHttpClientV1
	fixture     *KeyClientV1Fixture
	ctx         context.Context
}

func newKeyHttpClientV1Test() *keyCommandableHttpClientV1Test {
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

	httpConfig := cconf.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.port", "8081",
		"connection.host", "localhost",
	)

	service := service1.NewKeyHttpServiceV1()
	service.Configure(ctx, httpConfig)

	client := clients1.NewKeyHttpClientV1()
	client.Configure(ctx, httpConfig)

	references := cref.NewReferencesFromTuples(ctx,
		cref.NewDescriptor("key", "persistence", "memory", "default", "1.0"), persistence,
		cref.NewDescriptor("key", "controller", "default", "default", "1.0"), controller,
		cref.NewDescriptor("key", "service", "http", "default", "1.0"), service,
		cref.NewDescriptor("key", "client", "http", "default", "1.0"), client,
	)
	controller.SetReferences(ctx, references)
	service.SetReferences(ctx, references)
	client.SetReferences(ctx, references)

	fixture := NewKeyClientV1Fixture(client)

	return &keyCommandableHttpClientV1Test{
		persistence: persistence,
		controller:  controller,
		service:     service,
		client:      client,
		fixture:     fixture,
		ctx:         ctx,
	}
}

func (c *keyCommandableHttpClientV1Test) setup(t *testing.T) {
	err := c.persistence.Open(c.ctx, "")
	if err != nil {
		t.Error("Failed to open persistence", err)
	}

	err = c.service.Open(c.ctx, "")
	if err != nil {
		t.Error("Failed to open service", err)
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

func (c *keyCommandableHttpClientV1Test) teardown(t *testing.T) {
	err := c.client.Close(c.ctx, "")
	if err != nil {
		t.Error("Failed to close client", err)
	}

	err = c.service.Close(c.ctx, "")
	if err != nil {
		t.Error("Failed to close service", err)
	}

	err = c.persistence.Close(c.ctx, "")
	if err != nil {
		t.Error("Failed to close persistence", err)
	}
}

func TestBeaconsHttpClientV1(t *testing.T) {
	c := newKeyHttpClientV1Test()

	c.setup(t)
	t.Run("CRUD Operations", c.fixture.TestCrudOperations)
	c.teardown(t)

	c.setup(t)
	t.Run("WithFilters", c.fixture.TestGetWithFilters)
	c.teardown(t)
}
