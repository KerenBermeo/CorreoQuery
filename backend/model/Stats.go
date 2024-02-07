package model

// Stats representa la estructura de las estadísticas de un índice.
type Stats struct {
	DocTimeMin  int `json:"doc_time_min"`
	DocTimeMax  int `json:"doc_time_max"`
	DocNum      int `json:"doc_num"`
	StorageSize int `json:"storage_size"`
	WalSize     int `json:"wal_size"`
}
