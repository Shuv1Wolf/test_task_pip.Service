package persistence_test

import (
	"context"
	"os"
	"testing"

	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
	"test_task_pip.Service/keystore_service/microservice/persistence"
)

type KeySqlitePersistenceTest struct {
	persistence *persistence.KeySqlitePersistence
	fixture     *KeyPersistenceFixture
}

func newKeySqlitePersistenceTest() *KeySqlitePersistenceTest {
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
	fixture := NewKeyPersistenceFixture(*persistence)

	return &KeySqlitePersistenceTest{
		persistence: persistence,
		fixture:     fixture,
	}
}

func (c *KeySqlitePersistenceTest) setup(t *testing.T) {
	err := c.persistence.Open(context.Background(), "")
	if err != nil {
		t.Error("Failed to open persistence", err)
	}

	err = c.persistence.Clear(context.Background(), "")
	if err != nil {
		t.Error("Failed to clear persistence", err)
	}
}

func (c *KeySqlitePersistenceTest) teardown(t *testing.T) {
	err := c.persistence.Close(context.Background(), "")
	if err != nil {
		t.Error("Failed to close persistence", err)
	}
}

func TestKeySqlitePersistence(t *testing.T) {
	c := newKeySqlitePersistenceTest()
	if c == nil {
		return
	}

	c.setup(t)
	t.Run("CRUD Operations", c.fixture.TestCrudOperations)
	c.teardown(t)

	c.setup(t)
	t.Run("WithFilters", c.fixture.TestGetWithFilters)
}
