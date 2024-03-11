package data1

type KeyV1 struct {
	Id    string `json:"id" bson:"_id"`
	Key   string `json:"key" bson:"key"`
	Owner string `json:"owner" bson:"owner"`
}

func (k KeyV1) Clone() KeyV1 {
	return KeyV1{
		Id:    k.Id,
		Key:   k.Key,
		Owner: k.Owner,
	}
}
