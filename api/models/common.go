package models

type PageRequest struct {
	Page      int    `json:"page" validate:"required,min=1" default:"1"`
	Size      int    `json:"size" validate:"required,min=1,max=100" default:"20"`
	Sort      string `json:"sort" default:"name"`
	Direction string `json:"direction" default:"asc"`
}

type IDRequest struct {
	ID int `json:"id" validate:"required,min=1"`
}

type PageResponse struct {
	Page  int `json:"page"`
	Size  int `json:"size"`
	Pages int `json:"pages"`
	Total int `json:"total"`
	Count int `json:"count"`
}

type EmptyResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Message string       `json:"message"`
	Details []Validation `json:"details,omitempty"`
}

type Validation struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	Value   any    `json:"value,omitempty"`
}
