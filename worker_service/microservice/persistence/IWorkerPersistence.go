package persistence

import (
	"context"

	data1 "test-task-pip.service/worker_service/microservice/data/version1"
)

type IWorkerPersistence interface {
	GetWorkAlias(ctx context.Context, correlationId string) (alias string)

	GetStatus(ctx context.Context, correlationId string) (status string)

	UpdateWorker(ctx context.Context, correlationId string, status string, workAlias string) (item data1.Worker)
}
