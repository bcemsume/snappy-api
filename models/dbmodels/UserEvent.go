package dbmodels

import (
	"time"

	"github.com/jinzhu/gorm"
)

// UserEvent s
type UserEvent struct {
	*gorm.Model
	User      User
	IsUsed    bool
	StampDate time.Time
}
