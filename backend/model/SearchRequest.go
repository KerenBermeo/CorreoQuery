package model

type SearchRequest struct {
	SearchType string   `json:"search_type"`
	Query      Query    `json:"query"`
	SortFields []string `json:"sort_fields"`
	From       int      `json:"from"`
	MaxResults int      `json:"maxresults"`
	Source     []string `json:"_source"`
}

type Query struct {
	Term  string `json:"term"`
	Field string `json:"field"`
}
