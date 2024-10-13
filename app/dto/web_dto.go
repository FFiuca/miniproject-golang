package dto

type WebResponse[T any] struct {
	Status int `json:"status"`
	Data   T   `json:"data,omitempty"`
}

type WebResponsePagination struct {
}
