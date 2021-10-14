package store

type SetKeyValueRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Done struct {
	Done string `json:"done"`
}

type KeyValueModel struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
