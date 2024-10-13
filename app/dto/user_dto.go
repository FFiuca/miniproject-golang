package dto

import (
	"fmt"
	"project1/app/models"
	"time"
)

type UserRequestCreate struct {
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=8"`
	PasswordConfirm string `json:"password_confirm" validate:"required,eqfield=Password"`
}

type UserRequestUpdate struct {
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=8"`
	PasswordConfirm string `json:"password_confirm" validate:"required,eqfield=Password"`
}

type UserRequestSearch struct {
	Params map[string]any `json:"params,omitempty"`
}

type UserResponse struct {
	ID uint `json:"id"`
	// StatusID uint `json:"status_id"`
	Email     string         `json:"email"`
	Status    StatusResponse `json:"status"`
	CreatedAt time.Time      `json:"created_at"`
}

func UserToResponse(m *models.User) *UserResponse {
	a := &UserResponse{
		ID:        m.ID,
		Email:     m.Email,
		CreatedAt: m.CreatedAt,
	}
	fmt.Println("user to response")

	if m.Status.ID != 0 {
		a.Status = *StatusToResponse(&m.Status)
	}

	// a.Status =

	return a
}
