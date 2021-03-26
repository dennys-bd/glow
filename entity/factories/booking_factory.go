package factories

import (
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/bluele/factory-go/factory"
	"github.com/dennys-bd/glow/entity"
)

type BookingGroup struct {
	Bookings []*entity.Booking
}

var BookingFactory = factory.NewFactory(
	&entity.Booking{},
).Attr("Name", func(args factory.Args) (interface{}, error) {
	return randomdata.FullName(randomdata.RandomGender), nil
}).Attr("Date", func(args factory.Args) (interface{}, error) {
	return time.Parse(stringToTimeLayout, randomdata.FullDate())
}).SubFactory("Entity", SkipValidation).OnCreate(createHandler).SubFactory("Class", ClassFactory)

var BookingGroupFactory = factory.NewFactory(
	&BookingGroup{},
).SubSliceFactory("Bookings", BookingFactory, func() int {
	return randomdata.Number(1, 5)
})
