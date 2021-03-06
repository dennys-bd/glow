// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package wire

import (
	"github.com/dennys-bd/glow/repository"
	"github.com/dennys-bd/glow/usecase"
	"github.com/google/wire"
)

// Injectors from wire.go:

func InjectClassController() (usecase.ClassController, error) {
	classRepository := repository.ProvideClassRepo()
	classController, err := usecase.ProvideClassCtrl(classRepository)
	if err != nil {
		return usecase.ClassController{}, err
	}
	return classController, nil
}

func InjectBookingController() (usecase.BookingController, error) {
	bookingRepository := repository.ProvideBookingRepo()
	classRepository := repository.ProvideClassRepo()
	bookingController, err := usecase.ProvideBookingCtrl(bookingRepository, classRepository)
	if err != nil {
		return usecase.BookingController{}, err
	}
	return bookingController, nil
}

// wire.go:

var classRepoSet = wire.NewSet(repository.ProvideClassRepo, wire.Bind(new(repository.ClassRepoInf), new(repository.ClassRepository)))

var bookingRepoSet = wire.NewSet(repository.ProvideBookingRepo, wire.Bind(new(repository.BookingRepoInf), new(repository.BookingRepository)))
