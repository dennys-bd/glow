package entity

import (
	"time"
)

type Class struct {
	Entity

	Capacity  uint      `json:"capacity" binding:"required"`
	Name      string    `json:"name" binding:"required"`
	StartDate time.Time `json:"start_date" binding:"required"`
	EndDate   time.Time `json:"end_date" binding:"required,gtefield=StartDate"`

	Bookings []Booking `json:"bookings,omitempty"`
}
