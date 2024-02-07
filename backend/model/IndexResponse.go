package model

// IndexResponse representa la estructura de la respuesta de la API para obtener la lista de índices.
type IndexResponse struct {
	List []Index `json:"list"`
	Page struct {
		PageNum  int `json:"page_num"`
		PageSize int `json:"page_size"`
		Total    int `json:"total"`
	} `json:"page"`
}
