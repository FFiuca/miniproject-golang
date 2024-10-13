package dto

import "project1/app/models"

type StatusResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type StatusRequestCreate struct {
	Name string `json:"name" validate:"required"`
}

type StatusRequestUpdate struct {
	ID   uint   `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

// type StatusRequestDelete struct {
// 	ID uint `json:"id"`
// }

// type StatusRequestDetail struct {
// 	ID uint `json:"id"`
// }

func StatusToResponse(data *models.Status) *StatusResponse {
	return &StatusResponse{
		ID:   data.ID,
		Name: data.Name,
	}
}
