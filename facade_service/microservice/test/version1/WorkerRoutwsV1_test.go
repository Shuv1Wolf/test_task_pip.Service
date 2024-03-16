package version1_test

// import (
// 	"testing"

// 	"net/http"

// 	"github.com/stretchr/testify/assert"
// 	operations1 "test-task-pip.service/facade_service/microservice/operations/version1"
// )

// type workerOperationsV1Test struct {
// 	workerOperations *operations1.WorkerOperationsV1
// }

// func newWorkerOperationsV1Test() *workerOperationsV1Test {
// 	c := &workerOperationsV1Test{
// 		workerOperations: operations1.NewWorkerOperationsV1(),
// 	}
// 	return c
// }

// func (c *workerOperationsV1Test) testWorkerOperations(t *testing.T) {
// 	status := func(res http.ResponseWriter, req *http.Request) { c.workerOperations.GetStatus(res, req) }
// 	assert.Equal(t, status, "stop")
// }

// func TestWorkerOperationsV1(t *testing.T) {
// 	c := newWorkerOperationsV1Test()

// 	t.Run("test", c.testWorkerOperations)
// }
