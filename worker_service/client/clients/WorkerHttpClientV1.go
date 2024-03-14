package clients

import (
	"context"

	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-rpc-gox/clients"
	cclients "github.com/pip-services3-gox/pip-services3-rpc-gox/clients"
)

type WorkerHttpClientV1 struct {
	*cclients.CommandableHttpClient
}

func NewWorkerHttpClientV1() *WorkerHttpClientV1 {
	c := &WorkerHttpClientV1{
		CommandableHttpClient: cclients.NewCommandableHttpClient("v1/workers"),
	}
	return c
}

func (c *WorkerHttpClientV1) GetWorkAlias(ctx context.Context, correlationId string) (alias string, err error) {
	response, err := c.CallCommand(ctx, "get_work_alias", correlationId, cdata.NewEmptyAnyValueMap())

	if err != nil {
		return "", err
	}

	return clients.HandleHttpResponse[string](response, correlationId)
}

func (c *WorkerHttpClientV1) GetStatus(ctx context.Context, correlationId string) (status string, err error) {
	response, err := c.CallCommand(ctx, "get_status", correlationId, cdata.NewEmptyAnyValueMap())

	if err != nil {
		return "", err
	}

	return clients.HandleHttpResponse[string](response, correlationId)
}

func (c *WorkerHttpClientV1) Start(ctx context.Context, correlationId string) (status string, err error) {
	response, err := c.CallCommand(ctx, "start", correlationId, cdata.NewEmptyAnyValueMap())

	if err != nil {
		return "", err
	}

	return clients.HandleHttpResponse[string](response, correlationId)
}

func (c *WorkerHttpClientV1) Stop(ctx context.Context, correlationId string) (status string, err error) {
	response, err := c.CallCommand(ctx, "stop", correlationId, cdata.NewEmptyAnyValueMap())

	if err != nil {
		return "", err
	}

	return clients.HandleHttpResponse[string](response, correlationId)
}
