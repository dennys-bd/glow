package entity

import (
	"time"
)

type Class struct {
	Entity

	Capacity  uint      `json:"capacity"`
	Name      string    `json:"name"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}
