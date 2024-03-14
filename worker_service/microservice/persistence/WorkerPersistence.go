package persistence

import (
	"context"

	data1 "test-task-pip.service/worker_service/microservice/data/version1"
)

type WorkerPersistence struct {
	*data1.Worker
}

func NewWorkerPersistence() *WorkerPersistence {
	c := &WorkerPersistence{&data1.Worker{
		Id:        "1",
		Status:    data1.Stop,
		WorkAlias: data1.NoWork,
	}}

	return c
}

func (c *WorkerPersistence) GetWorkAlias(ctx context.Context, correlationId string) (alias string) {
	return c.WorkAlias
}

func (c *WorkerPersistence) GetStatus(ctx context.Context, correlationId string) (status string) {
	return c.Status
}

func (c *WorkerPersistence) UpdateWorker(ctx context.Context, correlationId string,
	status string, workAlias string) (item data1.Worker) {

	c.Status = status
	c.WorkAlias = workAlias

	return *c.Worker
}
