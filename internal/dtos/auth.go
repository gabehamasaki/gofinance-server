package dtos

import "fmt"

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (dto *AuthRequest) Validate() error {
	if dto.Email == "" && dto.Password == "" {
		return fmt.Errorf("request body is empty or malformed")
	}

	if dto.Email == "" {
		return fmt.Errorf("email is empty or malformed")
	}

	if dto.Password == "" {
		return fmt.Errorf("password is empty or malformed")
	}

	return nil
}
