package factories

import (
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/bluele/factory-go/factory"
	"github.com/dennys-bd/glow/entity"
)

type ClassGroup struct {
	Classes []*entity.Class
}

var ClassFactory = factory.NewFactory(
	&entity.Class{},
).Attr("Capacity", func(args factory.Args) (interface{}, error) {
	return uint(randomdata.Number(20)), nil
}).Attr("Name", func(args factory.Args) (interface{}, error) {
	return randomdata.SillyName(), nil
}).Attr("StartDate", func(args factory.Args) (interface{}, error) {
	now := time.Now()
	future := now.AddDate(0, 0, 5)
	date := randomdata.FullDateInRange(now.Format(timeToStringLayout), future.Format(timeToStringLayout))
	return time.Parse(stringToTimeLayout, date)
}).Attr("EndDate", func(args factory.Args) (interface{}, error) {
	first := time.Now().AddDate(0, 0, 6)
	future := first.AddDate(0, 0, 10)
	date := randomdata.FullDateInRange(first.Format(timeToStringLayout), future.Format(timeToStringLayout))
	return time.Parse(stringToTimeLayout, date)
}).OnCreate(createHandler)

var ClassGroupFactory = factory.NewFactory(
	&ClassGroup{},
).SubSliceFactory("Classes", ClassFactory, func() int {
	return randomdata.Number(1, 5)
})
