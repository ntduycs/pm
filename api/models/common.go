package models

type PageRequest struct {
	Page      int    `json:"page"`
	Size      int    `json:"size"`
	Sort      string `json:"sort"`
	Direction string `json:"direction"`
}

type PageResponse struct {
	Page  int `json:"page"`
	Size  int `json:"size"`
	Pages int `json:"pages"`
	Total int `json:"total"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
