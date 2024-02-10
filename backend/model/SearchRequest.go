package model

type SearchRequest struct {
	SearchType string   `json:"search_type"`
	Query      struct{} `json:"query"`
	SortFields []string `json:"sort_fields"`
	From       int      `json:"from"`
	MaxResults int      `json:"max_results"`
	Source     []string `json:"_source"`
}
