package dbmodels

import (
	"time"

	"github.com/jinzhu/gorm"
)

// ClaimEvent s
type ClaimEvent struct {
	*gorm.Model
	UserID   uint
	User     User
	IsUsed   bool
	ScanDate time.Time
}
