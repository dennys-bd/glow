package usecase_test

import (
	"time"

	"github.com/dennys-bd/glow/entity"
	"gopkg.in/khaiql/dbcleaner.v2"
)

var Cleaner = dbcleaner.New()

type DataClass struct {
	Class entity.Class `json:"data"`
}

type DataBooking struct {
	Booking entity.Booking `json:"data"`
}

type PostClass struct {
	Name      string    `json:"name,omitempty"`
	Capacity  uint      `json:"capacity,omitempty"`
	StartDate time.Time `json:"start_date,omitempty"`
	EndDate   time.Time `json:"end_date,omitempty"`
}

type PostBooking struct {
	Name    string    `json:"name,omitempty"`
	Date    time.Time `json:"date,omitempty"`
	ClassID uint      `json:"class_id,omitempty"`
}

// Class is used to check the validation
type Class struct {
	Capacity  uint      `validate:"required"`
	Name      string    `validate:"required"`
	StartDate time.Time `validate:"required"`
	EndDate   time.Time `validate:"required,gtefield=StartDate"`
}

type Booking struct {
	Name    string    `validate:"required"`
	Date    time.Time `validate:"required,ltecsfield=Class.EndDate,gtecsfield=Class.StartDate"`
	ClassID uint      `validate:"required"`
	Class   *Class    `validate:"-"`
}

func Must(i interface{}, err error) interface{} {
	if err != nil {
		panic(err)
	}
	return i
}
