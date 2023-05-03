package dtos

import (
	"fmt"

	"github.com/google/uuid"
)

type CreateAccountDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (dto *CreateAccountDTO) Validate() error {

	if dto.Name == "" && dto.Email == "" && dto.Password == "" {
		return fmt.Errorf("request body is empty or malformed")
	}

	if dto.Name == "" {
		return fmt.Errorf("name is empty or malformed")
	}

	if dto.Email == "" {
		return fmt.Errorf("email is empty or malformed")
	}

	if dto.Password == "" {
		return fmt.Errorf("password is empty or malformed")
	}

	return nil
}

type UpdateAccountDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (dto *UpdateAccountDTO) Validate() error {

	if dto.Name == "" && dto.Email == "" && dto.Password == "" {
		return fmt.Errorf("request body is empty or malformed")
	}

	return nil
}

type CreateAccountResponseDTO struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

type UpdateAccountResponseDTO struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

type ShowAccountResponseDTO struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}
