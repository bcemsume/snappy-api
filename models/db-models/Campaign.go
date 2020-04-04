package dbmodels

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Campaign s
type Campaign struct {
	*gorm.Model
	Restaurant Restaurant `gorm:"association_autoupdate:false"`
	Product    Product    `gorm:"association_autoupdate:false"`
	Claim      int
	FinishDate time.Time
}
