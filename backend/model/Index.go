package model

// Index representa la estructura de un índice.
type Index struct {
	Name        string            `json:"name"`
	StorageType string            `json:"storage_type"`
	ShardNum    int               `json:"shard_num"`
	Settings    map[string]string `json:"settings"`
	Mappings    struct {
		Properties map[string]Property `json:"properties"`
	} `json:"mappings"`
	Stats   Stats  `json:"stats"`
	Version string `json:"version"`
}
