package repository

import "github.com/dennys-bd/glow/entity"

type BookingRepository struct{}

func ProvideBookingRepo() BookingRepository {
	return BookingRepository{}
}

func (r BookingRepository) Create(b *entity.Booking) error {
	return db.Create(b).Error
}

func (r BookingRepository) Find(id uint) (*entity.Booking, error) {
	var booking entity.Booking
	err := db.First(&booking, id).Error
	return &booking, err
}

func (r BookingRepository) List(params map[string]interface{}) ([]*entity.Booking, error) {
	var bookings []*entity.Booking
	err := db.Find(&bookings, params).Error
	return bookings, err
}

func (r BookingRepository) Delete(id uint) error {
	return db.Delete(&entity.Booking{}, id).Error
}
