package model

import "time"

type Image struct {
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	EditedAt  time.Time `json:"edited_at"`
}
