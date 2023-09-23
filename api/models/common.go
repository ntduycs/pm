package models

type PageRequest struct {
	Page      int32  `json:"page"`
	Size      int32  `json:"size"`
	Sort      string `json:"sort"`
	Direction string `json:"direction"`
}

type PageResponse struct {
	Page  int32 `json:"page"`
	Size  int32 `json:"size"`
	Pages int32 `json:"pages"`
	Total int32 `json:"total"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
