package dbmodels

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Campaign s
type Campaign struct {
	*gorm.Model
	ProductID  uint
	Product    Product
	Claim      int
	FinishDate time.Time
}
