package dbmodels

import (
	"time"

	"github.com/jinzhu/gorm"
)

// ClaimEvent s
type ClaimEvent struct {
	*gorm.Model
	User     User     `gorm:"association_autoupdate:false"`
	Campaign Campaign `gorm:"association_autoupdate:false"`
	isUsed   bool
	UsedTime time.Time
}
