package entity

import (
	"time"
)

type Class struct {
	Entity

	Capacity  uint      `json:"capacity" binding:"required"`
	Name      string    `json:"name" binding:"required"`
	StartDate time.Time `json:"start_date" time_format:"2006-01-02" binding:"required"`
	EndDate   time.Time `json:"end_date" time_format:"2006-01-02" binding:"required,gtecsfield=StartDate"`
}
