package model

type Email struct {
	Date    string   `json:"Date"`
	From    string   `json:"From"`
	To      []string `json:"To"`
	Subject string   `json:"Subject"`
	Content string   `json:"Content"`
}
