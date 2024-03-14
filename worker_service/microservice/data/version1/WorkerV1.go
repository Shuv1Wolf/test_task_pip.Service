package data1

type Worker struct {
	Id        string
	Status    string
	WorkAlias string
	// name to identify the job the worker is currently working on (job.id + job.owner)
}
