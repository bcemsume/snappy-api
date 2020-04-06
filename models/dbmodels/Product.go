package dbmodels

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Product s
type Product struct {
	*gorm.Model
	Description  string
	Price        float64
	FinishDate   time.Time
	RestaurantID uint `gorm:"not null"`
	Restaurant   Restaurant
	Campaigns    []Campaign
}
