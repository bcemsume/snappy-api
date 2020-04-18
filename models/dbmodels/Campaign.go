package dbmodels

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Campaign s
type Campaign struct {
	*gorm.Model
	ProductID  uint
	Claim      int
	FinishDate time.Time
	Users      []*User
}
