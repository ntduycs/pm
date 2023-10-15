package models

import "project-management/constants"

type EaWeeklyMember struct {
	Member  *Member           `json:"member"`
	Efforts []*EaWeeklyEffort `json:"efforts"`
}

type EaWeeklyEffort struct {
	Category constants.LabelCategory `json:"category"`
	Time     int                     `json:"time"`
}

type GetEaWeeklyReportRequest struct {
	ListRequest
}

type GetEaWeeklyReportResponse struct {
	ListResponse
	Items []EaWeeklyMember `json:"items"`
}

type UploadEaWeeklyReportResponse struct {
	ListResponse
	Items []EaWeeklyMember `json:"items"`
}
