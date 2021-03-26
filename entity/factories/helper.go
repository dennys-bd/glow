package factories

import (
	"github.com/bluele/factory-go/factory"
	"github.com/dennys-bd/glow/entity"
	"gorm.io/gorm"
)

type key int

var timeToStringLayout = "2006-01-02"
var stringToTimeLayout = "Monday 2 Jan 2006"

const (
	DBContextKey key = iota
	// ...
)

var SkipValidation = factory.NewFactory(
	entity.Entity{},
).Attr("SkipValidation", func(args factory.Args) (interface{}, error) {
	return true, nil
})

var createHandler = func(args factory.Args) error {
	if ctx := args.Context().Value(DBContextKey); ctx != nil {
		db := ctx.(*gorm.DB)
		return db.Create(args.Instance()).Error
	}
	return nil
}
