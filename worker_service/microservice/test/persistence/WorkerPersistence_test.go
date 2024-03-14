package persistence_test

import (
	"testing"

	"test-task-pip.service/worker_service/microservice/persistence"
)

type WorkerPersistenceTest struct {
	persistence *persistence.WorkerPersistence
	fixture     *WorkerPersistenceFixture
}

func newWorkerPersistenceTest() *WorkerPersistenceTest {

	persistence := persistence.NewWorkerPersistence()

	fixture := NewWorkerPersistenceFixture(*persistence)

	return &WorkerPersistenceTest{
		persistence: persistence,
		fixture:     fixture,
	}
}

func TestWorkerPersistence(t *testing.T) {
	c := newWorkerPersistenceTest()
	if c == nil {
		return
	}
	t.Run("CRUD Operations", c.fixture.TestWorker)
}
