package persistence_test

import (
	"context"
	"os"
	"testing"

	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
	"test-task-pip.service/jobs_service/microservice/persistence"
)

type JobSqlitePersistenceTest struct {
	persistence *persistence.JobSqlitePersistence
	fixture     *JobPersistenceFixture
}

func newJobSqlitePersistenceTest() *JobSqlitePersistenceTest {
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

	persistence := persistence.NewJobSqlitePersistence()
	persistence.Configure(context.Background(), dbConfig)
	fixture := NewJobPersistenceFixture(*persistence)

	return &JobSqlitePersistenceTest{
		persistence: persistence,
		fixture:     fixture,
	}
}

func (c *JobSqlitePersistenceTest) setup(t *testing.T) {
	err := c.persistence.Open(context.Background(), "")
	if err != nil {
		t.Error("Failed to open persistence", err)
	}

	err = c.persistence.Clear(context.Background(), "")
	if err != nil {
		t.Error("Failed to clear persistence", err)
	}
}

func (c *JobSqlitePersistenceTest) teardown(t *testing.T) {
	err := c.persistence.Close(context.Background(), "")
	if err != nil {
		t.Error("Failed to close persistence", err)
	}
}

func TestJobSqlitePersistence(t *testing.T) {
	c := newJobSqlitePersistenceTest()
	if c == nil {
		return
	}

	c.setup(t)
	t.Run("CRUD Operations", c.fixture.TestCrudOperations)
	c.teardown(t)

	c.setup(t)
	t.Run("WithFilters", c.fixture.TestGetWithFilters)
	c.teardown(t)
}
