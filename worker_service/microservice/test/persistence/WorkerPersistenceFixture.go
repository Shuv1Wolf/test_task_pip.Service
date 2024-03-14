package persistence_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	data1 "test-task-pip.service/worker_service/microservice/data/version1"
	"test-task-pip.service/worker_service/microservice/persistence"
)

type WorkerPersistenceFixture struct {
	persistence persistence.IWorkerPersistence
}

func NewWorkerPersistenceFixture(persistence persistence.WorkerPersistence) *WorkerPersistenceFixture {
	c := WorkerPersistenceFixture{}
	c.persistence = &persistence
	return &c
}

func (c *WorkerPersistenceFixture) TestWorker(t *testing.T) {
	// Get status
	status := c.persistence.GetStatus(context.Background(), "")
	assert.Equal(t, data1.Stop, status)

	// Get work alias
	workAlias := c.persistence.GetWorkAlias(context.Background(), "")
	assert.Equal(t, data1.NoWork, workAlias)

	worker := c.persistence.UpdateWorker(context.Background(), "", data1.Working, "1test")
	assert.NotNil(t, worker)
	assert.Equal(t, data1.Working, worker.Status)
	assert.Equal(t, "1test", worker.WorkAlias)

}
