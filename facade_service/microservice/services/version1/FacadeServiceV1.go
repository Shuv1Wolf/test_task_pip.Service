package services1

import (
	"net/http"

	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	rpcservices "github.com/pip-services3-go/pip-services3-rpc-go/services"
	operations1 "test-task-pip.service/facade_service/microservice/operations/version1"
)

type FacadeServiceV1 struct {
	*rpcservices.RestService
	workerOperations *operations1.WorkerOperationsV1
	keysOperations   *operations1.KeysOperationsV1
	jobOperations    *operations1.JobsOperationsV1
}

func NewFacadeServiceV1() *FacadeServiceV1 {
	c := &FacadeServiceV1{
		workerOperations: operations1.NewWorkerOperationsV1(),
		keysOperations:   operations1.NewKeysOperationsV1(),
		jobOperations:    operations1.NewJobsOperationsV1(),
	}
	c.RestService = rpcservices.InheritRestService(c)
	c.BaseRoute = "api/v1"
	return c
}

func (c *FacadeServiceV1) Configure(config *cconf.ConfigParams) {
	c.RestService.Configure(config)
	c.workerOperations.Configure(config)
}

func (c *FacadeServiceV1) SetReferences(references cref.IReferences) {
	c.RestService.SetReferences(references)
	c.workerOperations.SetReferences(references)
}

func (c *FacadeServiceV1) Register() {

	// Restore session middleware
	c.RegisterOpenApiSpec("")

	c.registerContentManagementRoutes()
}

func (c *FacadeServiceV1) registerContentManagementRoutes() {
	// Worker routes
	c.RegisterRoute("get", "/worker/start", nil,
		func(res http.ResponseWriter, req *http.Request) { c.workerOperations.Start(res, req) })
	c.RegisterRoute("get", "/worker/stop", nil,
		func(res http.ResponseWriter, req *http.Request) { c.workerOperations.Stop(res, req) })
	c.RegisterRoute("get", "/worker/get_status", nil,
		func(res http.ResponseWriter, req *http.Request) { c.workerOperations.GetStatus(res, req) })
	c.RegisterRoute("get", "/worker/get_work_alias", nil,
		func(res http.ResponseWriter, req *http.Request) { c.workerOperations.GetWorkAlias(res, req) })

	// Keys routes
	c.RegisterRoute("get", "/keys/get_keys", nil,
		func(res http.ResponseWriter, req *http.Request) { c.keysOperations.GetKeys(res, req) })
	c.RegisterRoute("get", "/keys/get_key_by_id", nil,
		func(res http.ResponseWriter, req *http.Request) { c.keysOperations.GetkeyById(res, req) })
	c.RegisterRoute("get", "/keys/get_key_by_owner", nil,
		func(res http.ResponseWriter, req *http.Request) { c.keysOperations.GetKeyByOwner(res, req) })
	c.RegisterRoute("post", "/keys/update", nil,
		func(res http.ResponseWriter, req *http.Request) { c.keysOperations.UpdateKey(res, req) })
	c.RegisterRoute("delete", "/keys/delete_by_id", nil,
		func(res http.ResponseWriter, req *http.Request) { c.keysOperations.DeleteKeyById(res, req) })

	// Jobs routes
	c.RegisterRoute("get", "/jobs/get_jobs", nil,
		func(res http.ResponseWriter, req *http.Request) { c.jobOperations.GetJobs(res, req) })
	c.RegisterRoute("get", "/jobs/get_not_started_jobs", nil,
		func(res http.ResponseWriter, req *http.Request) { c.jobOperations.GetNotStartedJobs(res, req) })
	c.RegisterRoute("get", "/jobs/get_completed_jobs", nil,
		func(res http.ResponseWriter, req *http.Request) { c.jobOperations.GetCompletedJobs(res, req) })
	c.RegisterRoute("get", "/jobs/get_progress_jobs", nil,
		func(res http.ResponseWriter, req *http.Request) { c.jobOperations.GetProgressJobs(res, req) })
}
