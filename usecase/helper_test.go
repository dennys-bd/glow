package usecase_test

import (
	"time"

	"github.com/dennys-bd/glow/entity"
)

type DataClass struct {
	Class entity.Class `json:"data"`
}

type PostClass struct {
	Name      string    `json:"name,omitempty"`
	Capacity  uint      `json:"capacity,omitempty"`
	StartDate time.Time `json:"start_date,omitempty"`
	EndDate   time.Time `json:"end_date,omitempty"`
}

// Class is used to check the validation
type Class struct {
	Capacity  uint      `validate:"required"`
	Name      string    `validate:"required"`
	StartDate time.Time `validate:"required"`
	EndDate   time.Time `validate:"required,gtefield=StartDate"`
}

func Must(i interface{}, err error) interface{} {
	if err != nil {
		panic(err)
	}
	return i
}
