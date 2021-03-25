package migrations

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var m202103242257 = gormigrate.Migration{
	ID: "202103242257",
	Migrate: func(tx *gorm.DB) error {

		type Class struct {
			ID uint
		}

		type Booking struct {
			ID        uint
			Name      string
			Date      time.Time
			ClassID   uint
			Class     Class `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
			CreatedAt time.Time
			UpdatedAt time.Time
			DeletedAt *time.Time `sql:"index"`
		}

		return tx.AutoMigrate(&Booking{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("bookings")
	},
}
