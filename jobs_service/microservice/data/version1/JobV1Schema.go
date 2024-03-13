package data1

import (
	cconv "github.com/pip-services3-gox/pip-services3-commons-gox/convert"
	cvalid "github.com/pip-services3-gox/pip-services3-commons-gox/validate"
)

type JobV1Schema struct {
	cvalid.ObjectSchema
}

func NewKeyV1schema() *JobV1Schema {
	c := JobV1Schema{}
	c.ObjectSchema = *cvalid.NewObjectSchema()

	c.WithOptionalProperty("id", cconv.String)
	c.WithRequiredProperty("owner", cconv.String)
	c.WithRequiredProperty("status", cconv.String)
	return &c
}
