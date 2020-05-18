package model

type MovieResponse struct {
	Page         int     `json:"page"`
	TotalResults int     `json:"total_results"`
	TotalPages   int     `json:"total_pages"`
	Results      []Movie `json:"results"`
}
