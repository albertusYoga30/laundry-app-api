package model

import "time"

type Duration struct {
	DurationID   int       `json:"duration_id"`
	DurationName string    `json:"duration_name"`
	DurationDays int       `json:"duration_days"`
	UpdatedAt    time.Time `json:"update_at"`
	CreatedAt    time.Time `json:"created_at"`
}
