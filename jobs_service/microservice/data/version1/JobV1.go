package data1

type JobV1 struct {
	Id     string `json:"id" bson:"_id"`
	Owner  string `json:"owner" bson:"owner"`
	Status string `json:"status" bson:"status"`
}

func (k JobV1) Clone() JobV1 {
	return JobV1{
		Id:     k.Id,
		Owner:  k.Owner,
		Status: k.Status,
	}
}
