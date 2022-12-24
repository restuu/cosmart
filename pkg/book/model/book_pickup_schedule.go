package model

import "time"

// BookSchedule ...
type BookSchedule struct {
	BookKey  string    `json:"book_key" validate:"required"`
	PickupAt time.Time `json:"pickup_at" validate:"required"`
}
