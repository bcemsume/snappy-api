package dbmodels

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Claim struct {
	*gorm.Model
	CampaingID, UserID uint
	IsUsed             bool
	Claim              int
	UsedTime           time.Time
}
