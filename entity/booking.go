package entity

import (
	"time"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type Booking struct {
	Entity

	Name string    `json:"name" binding:"required"`
	Date time.Time `json:"date" binding:"required" validate:"ltecsfield=Class.EndDate,gtecsfield=Class.StartDate"`

	ClassID uint   `json:"class_id" binding:"required"`
	Class   *Class `json:"-" binding:"-"`
}

func (b *Booking) BeforeCreate(tx *gorm.DB) error {
	if b.SkipValidation {
		return nil
	}
	return validator.New().Struct(b)
}
