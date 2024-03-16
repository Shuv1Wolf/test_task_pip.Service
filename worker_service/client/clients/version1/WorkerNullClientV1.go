package clients1

import "context"

type WorkerNullClientV1 struct{}

func NewWorkerNullClientV1() *WorkerNullClientV1 {
	return &WorkerNullClientV1{}
}

func (c *WorkerNullClientV1) GetWorkAlias(ctx context.Context, correlationId string) (alias string, err error) {
	return "", nil
}

func (c *WorkerNullClientV1) GetStatus(ctx context.Context, correlationId string) (status string, err error) {
	return "", nil
}

func (c *WorkerNullClientV1) Start(ctx context.Context, correlationId string) (status string, err error) {
	return "", nil
}

func (c *WorkerNullClientV1) Stop(ctx context.Context, correlationId string) (status string, err error) {
	return "", nil
}
