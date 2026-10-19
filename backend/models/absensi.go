package models

import "time"

type Absensi struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Type      string    `json:"type"`
	Timestamp time.Time `json:"timestamp"`
}