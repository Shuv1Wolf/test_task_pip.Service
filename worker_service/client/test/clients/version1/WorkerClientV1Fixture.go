package clients1_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	clients1 "test-task-pip.service/worker_service/client/clients/version1"
	data1 "test-task-pip.service/worker_service/microservice/data/version1"
)

type WorkerClientV1Fixture struct {
	client clients1.IWorkerClientV1
	ctx    context.Context
}

func NewWorkerClientV1Fixture(client clients1.IWorkerClientV1) *WorkerClientV1Fixture {
	c := WorkerClientV1Fixture{}
	c.client = client
	c.ctx = context.Background()
	return &c
}

func (c *WorkerClientV1Fixture) TestWorker(t *testing.T) {
	status, err := c.client.GetStatus(c.ctx, "")
	assert.Nil(t, err)
	assert.Equal(t, status, data1.Stop)

	alias, err := c.client.GetWorkAlias(c.ctx, "")
	assert.Nil(t, err)
	assert.Equal(t, alias, data1.NoWork)

	c.client.Start(c.ctx, "")

	status, err = c.client.GetStatus(c.ctx, "")
	assert.Nil(t, err)
	assert.Equal(t, status, data1.Waiting)

	c.client.Stop(c.ctx, "")

	status, err = c.client.GetStatus(c.ctx, "")
	assert.Nil(t, err)
	assert.Equal(t, status, data1.Stop)
}
