package models

type Transaction struct {
	BaseModel
	Title       string  `json:"title"`
	Description *string `json:"description"`
	Value       float64 `json:"value"`
	Type        string  `json:"type"`
	Owner       string  `json:"owner"`
}
