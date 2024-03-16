package clients1

import "context"

type IWorkerClientV1 interface {
	GetWorkAlias(ctx context.Context, correlationId string) (alias string, err error)

	GetStatus(ctx context.Context, correlationId string) (status string, err error)

	Start(ctx context.Context, correlationId string) (status string, err error)

	Stop(ctx context.Context, correlationId string) (status string, err error)
}
