package logic

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	ccmd "github.com/pip-services3-gox/pip-services3-commons-gox/commands"
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	"github.com/pip-services3-gox/pip-services3-commons-gox/run"
	data1 "test-task-pip.service/jobs_service/microservice/data/version1"
	data1Key "test-task-pip.service/keystore_service/microservice/data/version1"
	clients1 "test-task-pip.service/worker_service/microservice/clients/version1"
	data1Worker "test-task-pip.service/worker_service/microservice/data/version1"
	workerPersist "test-task-pip.service/worker_service/microservice/persistence"
)

type WorkerController struct {
	workerPersist workerPersist.IWorkerPersistence
	commandSet    *WorkerCommandSet
	timer         run.FixedRateTimer
	jobClient     clients1.IJobClientV1
	keyClient     clients1.IKeyClientV1
}

func NewWorkerController() *WorkerController {
	c := &WorkerController{}

	return c
}

func (c *WorkerController) SetReferences(ctx context.Context, references cref.IReferences) {
	workerLocator := cref.NewDescriptor("worker", "persistence", "default", "*", "1.0")
	p, err := references.GetOneRequired(workerLocator)
	if p != nil && err == nil {
		if _pers, ok := p.(workerPersist.IWorkerPersistence); ok {
			c.workerPersist = _pers
		}
	}

	jobLocator := cref.NewDescriptor("job", "client", "*", "*", "1.0")
	p, err = references.GetOneRequired(jobLocator)
	if p != nil && err == nil {
		if _client, ok := p.(clients1.IJobClientV1); ok {
			c.jobClient = _client
		}
	}

	keyLocator := cref.NewDescriptor("key", "client", "*", "*", "1.0")
	p, err = references.GetOneRequired(keyLocator)
	if p != nil && err == nil {
		if _client, ok := p.(clients1.IKeyClientV1); ok {
			c.keyClient = _client
			return
		}
	}

	panic(cref.NewReferenceError("worker.controller.SetReferences", fmt.Sprintf("%s or %s or %s", workerLocator, jobLocator, keyLocator)))
}

func (c *WorkerController) GetCommandSet() *ccmd.CommandSet {
	if c.commandSet == nil {
		c.commandSet = NewWorkerCommandSet(c)
	}
	return &c.commandSet.CommandSet
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-=_$%*0123456789"

func randomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func generateAndSleep() string {
	randomStr := randomString(20)
	sleepTime := rand.Intn(6) + 10 // генерируем случайное время от 10 до 15 секунд
	time.Sleep(time.Duration(sleepTime) * time.Second)
	return randomStr
}

// TODO:
func (c *WorkerController) Start(ctx context.Context, correlationId string) (status string, err error) {

	updateWorker := c.workerPersist.UpdateWorker(ctx, correlationId, data1Worker.Waiting, data1Worker.NoWork)
	if updateWorker.Id == "" {
		return
	}

	c.timer = *run.NewFixedRateTimerFromCallback(func(ctx context.Context) {
		job, err := c.jobClient.GetNotStartedJob(ctx, correlationId)
		if err != nil {
			return
		}
		if job.Id == "" {
			return
		}

		updateJob, err := c.jobClient.UpdateInProgress(ctx, correlationId, job.Id, job.Owner)
		if err != nil {
			return
		}
		if updateJob.Id == "" {
			return
		}

		updateWorker = c.workerPersist.UpdateWorker(ctx, correlationId, data1Worker.Working,
			fmt.Sprintf("%s%s", job.Id, job.Owner))
		if updateWorker.Id == "" {
			return
		}

		genString := generateAndSleep() // генерация рандомного числа и сон

		key := data1Key.KeyV1{
			Id:    job.Id,
			Owner: job.Owner,
			Key:   genString,
		}

		createKey, err := c.keyClient.CreateKey(ctx, correlationId, key)
		if err != nil {
			return
		}
		if createKey.Id == "" {
			return
		}

		updateJob, err = c.jobClient.UpdateInCompleted(ctx, correlationId, job.Id, job.Owner)
		if err != nil {
			return
		}
		if updateJob.Id == "" {
			return
		}

		updateWorker = c.workerPersist.UpdateWorker(ctx, correlationId, data1Worker.Waiting, data1Worker.NoWork)
		if updateWorker.Id == "" {
			return
		}

	}, 15000, 0, 1)

	c.timer.Start(ctx)

	return fmt.Sprintf("worker №%s start", updateWorker.Id), err
}

func (c *WorkerController) GetWorkAlias(ctx context.Context, correlationId string) (alias string, err error) {
	status := c.workerPersist.GetWorkAlias(ctx, correlationId)
	return status, err
}

func (c *WorkerController) GetStatus(ctx context.Context, correlationId string) (status string, err error) {
	status = c.workerPersist.GetStatus(ctx, correlationId)
	return status, err
}

func (c *WorkerController) Stop(ctx context.Context, correlationId string) (status string, err error) {
	c.timer.Stop(ctx)

	// если есть job progress, то ставим not_started
	page, err := c.jobClient.GetJobsByStatus(ctx, correlationId, data1.Progress)
	for _, item := range page.Data {
		c.jobClient.UpdateInNotStarted(ctx, correlationId, item.Id, item.Owner)
	}

	updateWorker := c.workerPersist.UpdateWorker(ctx, correlationId, data1Worker.Stop, data1Worker.NoWork)

	return fmt.Sprintf("worker №%s stop", updateWorker.Id), err
}
