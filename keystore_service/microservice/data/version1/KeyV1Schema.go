package data1

import (
	cconv "github.com/pip-services3-gox/pip-services3-commons-gox/convert"
	cvalid "github.com/pip-services3-gox/pip-services3-commons-gox/validate"
)

type KeyV1Schema struct {
	cvalid.ObjectSchema
}

func NewKeyV1schema() *KeyV1Schema {
	c := KeyV1Schema{}
	c.ObjectSchema = *cvalid.NewObjectSchema()

	c.WithOptionalProperty("id", cconv.String)
	c.WithRequiredProperty("key", cconv.String)
	c.WithRequiredProperty("owner", cconv.String)
	return &c
}
