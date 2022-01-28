package structures

import (
	"encoding/json"
	"polar/models"
)

type Metadata struct {
	Name   string
	Size   int64
	Chunks []MetadataBloc
}

type MetadataBloc struct {
	Checksum string
	Objects  [][]byte // <ip>/<port>/<uid> : 93.95.228.237/1312/2fb2cfa7bac33a8b02a4b0ce8f85c46feb90f4ea20697c52fb84a855baf9202b
}

// Append object to db
func (m Metadata) Create() models.Object {
	toJson, err := json.Marshal(&m)
	if err != nil {
		panic(err)
	}

	return models.Object{
		Uid:  "0xACAB",
		Data: toJson,
	}
}

// Delete object from db
func (m Metadata) Parse(o models.Object) {
	err := json.Unmarshal(o.Data, &m)
	if err != nil {
		panic(err)
	}
}
