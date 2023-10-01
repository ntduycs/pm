package constants

import "time"

const (
	StatusActive   = "active"
	StatusInactive = "inactive"
	StatusPending  = "pending"
)

var (
	CommonStatuses = []string{
		StatusActive,
		StatusInactive,
		StatusPending,
	}

	TimeZone, _ = time.LoadLocation("Asia/Saigon")
)
