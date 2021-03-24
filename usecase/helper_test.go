package usecase_test

import (
	"time"

	"github.com/dennys-bd/glow/entity"
)

type DataClass struct {
	Class entity.Class `json:"data"`
}

type PostClass struct {
	Name      string    `json:"name"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

func Must(i interface{}, err error) interface{} {
	if err != nil {
		panic(err)
	}
	return i
}
