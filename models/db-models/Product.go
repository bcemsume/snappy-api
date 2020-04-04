package dbmodels

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Product s
type Product struct {
	*gorm.Model
	Restaurant  Restaurant `gorm:"association_autoupdate:false"`
	Description string
	Price       uint
	FinishDate  time.Time
}
