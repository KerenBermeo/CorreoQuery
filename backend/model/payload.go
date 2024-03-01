package model

type Payload struct {
	Index   string  `json:"index"`
	Records []Email `json:"records"`
}
