package dtos

import "fmt"

type CreateTransactionDTO struct {
	Title       string  `json:"title"`
	Description *string `json:"description"`
	Value       float64 `json:"value"`
	Type        string  `json:"type"`
	Owner       string  `json:"owner"`
}

func (dto *CreateTransactionDTO) Validate() error {
	if dto.Title == "" && dto.Value <= 0 && dto.Type == "" && dto.Owner == "" {
		return fmt.Errorf("request body is empty or malformed")
	}

	if dto.Owner == "" {
		return fmt.Errorf("owner is empty or malformed")
	}

	if dto.Title == "" {
		return fmt.Errorf("title is empty or malformed")
	}

	if dto.Value <= 0 {
		return fmt.Errorf("value is empty or malformed")
	}

	if dto.Type == "" {
		return fmt.Errorf("type is empty or malformed")
	}

	return nil
}
