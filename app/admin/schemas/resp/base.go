package resp

type Select struct {
	Label string      `json:"label" structs:"label"`
	Value interface{} `json:"value" structs:"value"`
}
