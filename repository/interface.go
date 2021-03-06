package repository

import "github.com/dennys-bd/glow/entity"

type ClassRepoInf interface {
	Create(c *entity.Class) error
	Find(id uint) (*entity.Class, error)
	List(params map[string]interface{}) ([]*entity.Class, error)
}

type BookingRepoInf interface {
	Create(c *entity.Booking) error
	Find(id uint) (*entity.Booking, error)
	List(params map[string]interface{}) ([]*entity.Booking, error)
	Delete(id uint) error
}
