package dtos

import (
	"fmt"

	"github.com/google/uuid"
)

type CreateTransactionDTO struct {
	Title       string  `json:"title"`
	Description *string `json:"description"`
	Value       float64 `json:"value"`
	Type        string  `json:"type"`
}

func (dto *CreateTransactionDTO) Validate() error {
	if dto.Title == "" && dto.Value <= 0 && dto.Type == "" {
		return fmt.Errorf("request body is empty or malformed")
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

type UpdateTransactionRequestDTO struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	Value       float64   `json:"value"`
	Type        string    `json:"type"`
}

func (dto *UpdateTransactionRequestDTO) Validate() error {
	if dto.Title == "" && dto.Value <= 0 && dto.Type == "" && dto.Description == nil {
		return fmt.Errorf("request body is empty or malformed")
	}

	return nil
}

type TransactionResponseDTO struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	Value       float64   `json:"value"`
	Type        string    `json:"type"`
}

type UpdateTransactionResponseDTO struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	Value       float64   `json:"value"`
	Type        string    `json:"type"`
}
