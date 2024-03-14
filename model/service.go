package model

import "time"

type Service struct {
	ServiceID    int       `json:"service_id"`
	ServiceName  string    `json:"service_name"`
	ServiceDesc  string    `json:"service_desc"`
	ServicePrice int       `json:"service_price"`
	UpdatedAt    time.Time `json:"update_at"`
	CreatedAt    time.Time `json:"created_at"`
}
